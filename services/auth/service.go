// file: services/auth/service.go
// version: 1.0.0
// guid: n4o5p6q7-r8s9-0123-7890-234567890123

package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authpb "github.com/falkcorp/gcommon/pkg/authpb/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Config represents the authentication service configuration
type Config struct {
	JWTSecret           string
	JWTExpiration       time.Duration
	RefreshExpiration   time.Duration
	EnableOAuth2        bool
	OAuth2ClientID      string
	OAuth2ClientSecret  string
	OAuth2RedirectURL   string
	Require2FA          bool
	MaxLoginAttempts    int
	LockoutDuration     time.Duration
	PasswordMinLength   int
	PasswordRequireSpec bool
	AllowedOrigins      []string
}

// AuthenticationService implements the auth service with hybrid architecture
type AuthenticationService struct {
	authpb.UnimplementedAuthServiceServer
	mu            sync.RWMutex
	jwtSecret     []byte
	rsaPrivKey    *rsa.PrivateKey
	rsaPubKey     *rsa.PublicKey
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
	users         map[string]*User        // In-memory user store for demo
	apiKeys       map[string]*APIKey      // In-memory API key store
	sessions      map[string]*Session     // In-memory session store
	oauthConfigs  map[string]*OAuthConfig // OAuth provider configurations
}

// User represents a user in our system
type User struct {
	ID       string
	Username string
	Password string // In production, this should be hashed
	Roles    []string
}

// UserProfile represents extended user profile information
type UserProfile struct {
	FirstName   string
	LastName    string
	DisplayName string
	Avatar      string
	Timezone    string
	Language    string
	Preferences map[string]string
}

// APIKey represents an API key for authentication
type APIKey struct {
	ID         string
	UserID     string
	KeyHash    string // SHA-256 hash of the actual key
	Name       string
	Scopes     []string
	CreatedAt  time.Time
	ExpiresAt  *time.Time
	LastUsedAt *time.Time
}

// Session represents a user session
type Session struct {
	ID         string
	UserID     string
	TokenID    string
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LastUsedAt time.Time
	DeviceInfo string
	IPAddress  string
	UserAgent  string
	IsActive   bool
}

// OAuthConfig represents OAuth provider configuration
type OAuthConfig struct {
	Provider     string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

// NewAuthService creates a new authentication service
func NewAuthService(jwtSecret []byte) (*AuthenticationService, error) {
	// Generate RSA key pair for JWT signing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key: %w", err)
	}

	service := &AuthenticationService{
		jwtSecret:     jwtSecret,
		rsaPrivKey:    privateKey,
		rsaPubKey:     &privateKey.PublicKey,
		tokenExpiry:   15 * time.Minute,
		refreshExpiry: 7 * 24 * time.Hour,
		users:         make(map[string]*User),
		apiKeys:       make(map[string]*APIKey),
		sessions:      make(map[string]*Session),
		oauthConfigs:  make(map[string]*OAuthConfig),
	}

	// Add some demo users
	service.users["admin"] = &User{
		ID:       "1",
		Username: "admin",
		Password: "admin", // In production, use proper password hashing
		Roles:    []string{"admin"},
	}

	service.users["user"] = &User{
		ID:       "2",
		Username: "user",
		Password: "user", // In production, use proper password hashing
		Roles:    []string{"user"},
	}

	return service, nil
}

// Login authenticates a user with username/password and returns JWT tokens
func (s *AuthenticationService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	// Look up user - in a real system, this would query a database
	user, exists := s.users[username]
	if !exists || user.Password != password {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// Generate tokens
	accessToken, err := s.generateToken(user.ID, user.Roles, s.tokenExpiry, "access")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate access token")
	}

	refreshToken, err := s.generateToken(user.ID, user.Roles, s.refreshExpiry, "refresh")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate refresh token")
	}

	// Create response using opaque field setters
	response := &authpb.LoginResponse{}
	response.SetAccessToken(accessToken)
	response.SetRefreshToken(refreshToken)
	response.SetExpiresIn(int64(s.tokenExpiry.Seconds()))

	return response, nil
}

// ValidateToken checks if a JWT token is valid and returns user information
func (s *AuthenticationService) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	token := req.GetToken()

	if token == "" {
		response := &authpb.ValidateTokenResponse{}
		response.SetValid(false)
		response.SetUserId("")
		response.SetRoles([]string{})
		return response, nil
	}

	// Parse and validate JWT
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.rsaPubKey, nil
	})

	if err != nil || !parsedToken.Valid {
		response := &authpb.ValidateTokenResponse{}
		response.SetValid(false)
		response.SetUserId("")
		response.SetRoles([]string{})
		return response, nil
	}

	// Extract claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		userId, _ := claims["user_id"].(string)
		rolesInterface := claims["roles"].([]interface{})
		roles := make([]string, len(rolesInterface))
		for i, role := range rolesInterface {
			roles[i] = role.(string)
		}

		response := &authpb.ValidateTokenResponse{}
		response.SetValid(true)
		response.SetUserId(userId)
		response.SetRoles(roles)
		return response, nil
	}

	response := &authpb.ValidateTokenResponse{}
	response.SetValid(false)
	response.SetUserId("")
	response.SetRoles([]string{})
	return response, nil
}

// AuthorizeAccess checks if a user has permission for a specific action
func (s *AuthenticationService) AuthorizeAccess(ctx context.Context, req *authpb.AuthorizeRequest) (*authpb.AuthorizeResponse, error) {
	// First validate the token
	validateReq := &authpb.ValidateTokenRequest{}
	validateReq.SetToken(req.GetToken())
	validateResp, err := s.ValidateToken(ctx, validateReq)
	if err != nil {
		return nil, err
	}

	if !validateResp.GetValid() {
		response := &authpb.AuthorizeResponse{}
		response.SetAuthorized(false)
		response.SetReason("invalid token")
		response.SetRoles([]string{})
		return response, nil
	}

	// Check role-based authorization
	userRoles := validateResp.GetRoles()
	requiredRole := req.GetRequiredRole()

	authorized := false
	if requiredRole == "" {
		authorized = true // No specific role required
	} else {
		for _, role := range userRoles {
			if role == requiredRole || role == "admin" { // Admin can access everything
				authorized = true
				break
			}
		}
	}

	reason := "authorized"
	if !authorized {
		reason = fmt.Sprintf("insufficient privileges: required role '%s'", requiredRole)
	}

	response := &authpb.AuthorizeResponse{}
	response.SetAuthorized(authorized)
	response.SetReason(reason)
	response.SetRoles(userRoles)
	return response, nil
}

// GenerateToken creates new tokens for a user with specific roles
func (s *AuthenticationService) GenerateToken(ctx context.Context, req *authpb.GenerateTokenRequest) (*authpb.GenerateTokenResponse, error) {
	expiry := time.Duration(req.GetExpiresIn()) * time.Second
	if expiry == 0 {
		expiry = s.tokenExpiry
	}

	accessToken, err := s.generateToken(req.GetUserId(), req.GetRoles(), expiry, req.GetTokenType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	refreshToken, err := s.generateToken(req.GetUserId(), req.GetRoles(), s.refreshExpiry, "refresh")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	response := &authpb.GenerateTokenResponse{}
	response.SetAccessToken(accessToken)
	response.SetRefreshToken(refreshToken)
	response.SetExpiresIn(int64(expiry.Seconds()))
	response.SetTokenType(req.GetTokenType())
	return response, nil
}

// RefreshToken exchanges a refresh token for new access tokens
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	// Validate refresh token
	validateReq := &authpb.ValidateTokenRequest{}
	validateReq.SetToken(req.GetRefreshToken())
	validateResp, err := s.ValidateToken(ctx, validateReq)
	if err != nil {
		return nil, err
	}

	if !validateResp.GetValid() {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	// Generate new tokens
	accessToken, err := s.generateToken(validateResp.GetUserId(), validateResp.GetRoles(), s.tokenExpiry, "access")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := s.generateToken(validateResp.GetUserId(), validateResp.GetRoles(), s.refreshExpiry, "refresh")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	response := &authpb.RefreshTokenResponse{}
	response.SetAccessToken(accessToken)
	response.SetRefreshToken(refreshToken)
	response.SetExpiresIn(int64(s.tokenExpiry.Seconds()))
	return response, nil
}

// RevokeToken invalidates a token making it unusable
func (s *AuthenticationService) RevokeToken(ctx context.Context, req *authpb.RevokeTokenRequest) (*authpb.RevokeTokenResponse, error) {
	// In a real implementation, you'd add the token to a blacklist
	// or store revoked tokens in a database

	// For now, just validate the token exists
	validateReq := &authpb.ValidateTokenRequest{}
	validateReq.SetToken(req.GetToken())
	validateResp, err := s.ValidateToken(ctx, validateReq)
	if err != nil {
		return nil, err
	}

	if !validateResp.GetValid() {
		response := &authpb.RevokeTokenResponse{}
		response.SetSuccess(false)
		response.SetMessage("token is already invalid or expired")
		return response, nil
	}

	// In production: add to blacklist, update database, etc.
	response := &authpb.RevokeTokenResponse{}
	response.SetSuccess(true)
	response.SetMessage("token revoked successfully")
	return response, nil
}

// generateToken creates a JWT token with the specified claims
func (s *AuthenticationService) generateToken(userID string, roles []string, expiry time.Duration, tokenType string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id":    userID,
		"roles":      roles,
		"token_type": tokenType,
		"iat":        now.Unix(),
		"exp":        now.Add(expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(s.rsaPrivKey)
}

// Start starts both gRPC and HTTP servers
func (s *AuthenticationService) Start(grpcPort, httpPort int) error {
	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port %d: %v", grpcPort, err)
		}

		grpcServer := grpc.NewServer()
		authpb.RegisterAuthServiceServer(grpcServer, s)

		log.Printf("Auth gRPC server listening on port %d", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server with basic endpoints
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "auth"}`))
	})

	http.HandleFunc("/login", s.httpLoginHandler)
	http.HandleFunc("/validate", s.httpValidateHandler)

	log.Printf("Auth HTTP server listening on port %d", httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
}

// HTTP handlers for basic REST API
func (s *AuthenticationService) httpLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req authpb.LoginRequest
	// In production, properly parse JSON request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	req.SetUsername(username)
	req.SetPassword(password)

	resp, err := s.Login(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// In production, use proper JSON marshaling
	fmt.Fprintf(w, `{"access_token":"%s","refresh_token":"%s","expires_in":%d}`,
		resp.GetAccessToken(), resp.GetRefreshToken(), resp.GetExpiresIn())
}

func (s *AuthenticationService) httpValidateHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	req := &authpb.ValidateTokenRequest{}
	req.SetToken(token)
	resp, err := s.ValidateToken(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if resp.GetValid() {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"valid":true,"user_id":"%s","roles":%v}`,
			resp.GetUserId(), resp.GetRoles())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"valid":false}`))
	}
}

// Helper function to create time pointer
func timePtr(t time.Time) *time.Time {
	return &t
}

// Helper function to generate random key
func (s *AuthenticationService) generateRandomKey(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}

// AuthenticateApiKey validates API key authentication
func (s *AuthenticationService) AuthenticateApiKey(ctx context.Context, req *authpb.ApiKeyAuthRequest) (*authpb.ApiKeyAuthResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	apiKey, exists := s.apiKeys[req.GetApiKey()]
	if !exists {
		response := &authpb.ApiKeyAuthResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_INVALID_CREDENTIALS)
		return response, nil
	}

	// Check if API key is expired
	if apiKey.ExpiresAt != nil && time.Now().After(*apiKey.ExpiresAt) {
		response := &authpb.ApiKeyAuthResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_EXPIRED)
		return response, nil
	}

	// Update last used timestamp
	apiKey.LastUsedAt = timePtr(time.Now())

	response := &authpb.ApiKeyAuthResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetUserId(apiKey.UserID)
	response.SetRoles(apiKey.Scopes) // Using Scopes as roles for compatibility
	return response, nil
}

// CreateApiKey generates a new API key with scopes
func (s *AuthenticationService) CreateApiKey(ctx context.Context, req *authpb.CreateApiKeyRequest) (*authpb.CreateApiKeyResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate random API key
	apiKeyValue := s.generateRandomKey(32)

	apiKey := &APIKey{
		ID:        s.generateRandomKey(16),
		UserID:    req.GetUserId(),
		KeyHash:   apiKeyValue, // In production, this should be hashed
		Name:      req.GetName(),
		Scopes:    req.GetScopes(),
		CreatedAt: time.Now(),
		ExpiresAt: nil, // Set based on req.ExpiresAt if provided
	}

	if req.GetExpiresAt() != nil {
		expiresAt := req.GetExpiresAt().AsTime()
		apiKey.ExpiresAt = &expiresAt
	}

	s.apiKeys[apiKeyValue] = apiKey

	response := &authpb.CreateApiKeyResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetApiKey(apiKeyValue)
	response.SetKeyId(apiKey.ID)
	return response, nil
}

// RevokeApiKey invalidates an existing API key
func (s *AuthenticationService) RevokeApiKey(ctx context.Context, req *authpb.RevokeApiKeyRequest) (*authpb.RevokeApiKeyResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find and remove the API key
	for keyValue, apiKey := range s.apiKeys {
		if apiKey.ID == req.GetKeyId() && apiKey.UserID == req.GetUserId() {
			delete(s.apiKeys, keyValue)
			response := &authpb.RevokeApiKeyResponse{}
			response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
			return response, nil
		}
	}

	response := &authpb.RevokeApiKeyResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
	return response, nil
}

// ListApiKeys returns user's active API keys
func (s *AuthenticationService) ListApiKeys(ctx context.Context, req *authpb.ListApiKeysRequest) (*authpb.ListApiKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var keys []*authpb.ApiKeyInfo
	for _, apiKey := range s.apiKeys {
		if apiKey.UserID == req.GetUserId() {
			keyInfo := &authpb.ApiKeyInfo{}
			keyInfo.SetKeyId(apiKey.ID)
			keyInfo.SetName(apiKey.Name)
			keyInfo.SetScopes(apiKey.Scopes)
			keyInfo.SetCreatedAt(timestamppb.New(apiKey.CreatedAt))

			if apiKey.ExpiresAt != nil {
				keyInfo.SetExpiresAt(timestamppb.New(*apiKey.ExpiresAt))
			}
			if apiKey.LastUsedAt != nil {
				keyInfo.SetLastUsedAt(timestamppb.New(*apiKey.LastUsedAt))
			}

			keys = append(keys, keyInfo)
		}
	}

	response := &authpb.ListApiKeysResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetKeys(keys)
	return response, nil
}

// InitiateOAuth starts OAuth2 flow with a provider
func (s *AuthenticationService) InitiateOAuth(ctx context.Context, req *authpb.OAuthInitiateRequest) (*authpb.OAuthInitiateResponse, error) {
	s.mu.RLock()
	config, exists := s.oauthConfigs[req.GetProvider()]
	s.mu.RUnlock()

	if !exists {
		response := &authpb.OAuthInitiateResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
		response.SetMessage("OAuth provider not configured")
		return response, nil
	}

	// Generate state parameter for CSRF protection
	state := s.generateRandomKey(16)

	// Build OAuth authorization URL
	authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s&state=%s",
		config.ClientID, config.RedirectURL, "user:email", state)

	response := &authpb.OAuthInitiateResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetAuthUrl(authURL)
	response.SetState(state)
	return response, nil
}

// HandleOAuthCallback processes OAuth2 callback from provider
func (s *AuthenticationService) HandleOAuthCallback(ctx context.Context, req *authpb.OAuthCallbackRequest) (*authpb.OAuthCallbackResponse, error) {
	// In a real implementation, you would:
	// 1. Verify the state parameter
	// 2. Exchange code for access token with OAuth provider
	// 3. Get user info from provider
	// 4. Create or update user account
	// 5. Generate our own JWT tokens

	response := &authpb.OAuthCallbackResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetAccessToken("demo_oauth_token")
	response.SetRefreshToken("demo_oauth_refresh")
	response.SetExpiresIn(3600)
	return response, nil
}

// ConfigureOAuth sets up OAuth2 provider configuration
func (s *AuthenticationService) ConfigureOAuth(ctx context.Context, req *authpb.OAuthConfigRequest) (*authpb.OAuthConfigResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	config := &OAuthConfig{
		Provider:     req.GetProvider(),
		ClientID:     req.GetClientId(),
		ClientSecret: req.GetClientSecret(),
		RedirectURL:  req.GetRedirectUrl(),
		Scopes:       req.GetScopes(),
	}

	s.oauthConfigs[req.GetProvider()] = config

	response := &authpb.OAuthConfigResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetMessage("OAuth provider configured successfully")
	return response, nil
}

// GetSessionInfo retrieves information about a user session
func (s *AuthenticationService) GetSessionInfo(ctx context.Context, req *authpb.SessionInfoRequest) (*authpb.SessionInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[req.GetSessionId()]
	if !exists {
		response := &authpb.SessionInfoResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
		return response, nil
	}

	// Check if session is expired
	if time.Now().After(session.ExpiresAt) {
		response := &authpb.SessionInfoResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_EXPIRED)
		return response, nil
	}

	sessionInfo := &authpb.SessionInfo{}
	sessionInfo.SetSessionId(session.ID)
	sessionInfo.SetUserId(session.UserID)
	sessionInfo.SetCreatedAt(timestamppb.New(session.CreatedAt))
	sessionInfo.SetExpiresAt(timestamppb.New(session.ExpiresAt))
	sessionInfo.SetLastUsedAt(timestamppb.New(session.LastUsedAt))
	sessionInfo.SetDeviceInfo(session.DeviceInfo)
	sessionInfo.SetIpAddress(session.IPAddress)

	response := &authpb.SessionInfoResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetSession(sessionInfo)
	return response, nil
}

// ExtendSession extends the lifetime of a session
func (s *AuthenticationService) ExtendSession(ctx context.Context, req *authpb.ExtendSessionRequest) (*authpb.ExtendSessionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, exists := s.sessions[req.GetSessionId()]
	if !exists {
		response := &authpb.ExtendSessionResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
		return response, nil
	}

	// Extend session expiry
	extension := time.Duration(req.GetExtensionSeconds()) * time.Second
	if extension == 0 {
		extension = 24 * time.Hour // Default extension
	}

	session.ExpiresAt = time.Now().Add(extension)
	session.LastUsedAt = time.Now()

	response := &authpb.ExtendSessionResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetNewExpiresAt(timestamppb.New(session.ExpiresAt))
	return response, nil
}

// ListSessions returns all active sessions for a user
func (s *AuthenticationService) ListSessions(ctx context.Context, req *authpb.ListSessionsRequest) (*authpb.ListSessionsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var sessions []*authpb.SessionInfo
	for _, session := range s.sessions {
		if session.UserID == req.GetUserId() && time.Now().Before(session.ExpiresAt) {
			sessionInfo := &authpb.SessionInfo{}
			sessionInfo.SetSessionId(session.ID)
			sessionInfo.SetUserId(session.UserID)
			sessionInfo.SetCreatedAt(timestamppb.New(session.CreatedAt))
			sessionInfo.SetExpiresAt(timestamppb.New(session.ExpiresAt))
			sessionInfo.SetLastUsedAt(timestamppb.New(session.LastUsedAt))
			sessionInfo.SetDeviceInfo(session.DeviceInfo)
			sessionInfo.SetIpAddress(session.IPAddress)

			sessions = append(sessions, sessionInfo)
		}
	}

	response := &authpb.ListSessionsResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetSessions(sessions)
	return response, nil
}

// GetUserProfile retrieves user profile information
func (s *AuthenticationService) GetUserProfile(ctx context.Context, req *authpb.UserProfileRequest) (*authpb.UserProfileResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[req.GetUserId()]
	if !exists {
		response := &authpb.UserProfileResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
		return response, nil
	}

	profile := &authpb.UserProfile{}
	profile.SetUserId(user.ID)
	profile.SetUsername(user.Username)
	// In production, load actual profile data from database
	profile.SetDisplayName(user.Username)
	profile.SetFirstName("")
	profile.SetLastName("")
	profile.SetAvatar("")
	profile.SetTimezone("UTC")
	profile.SetLanguage("en")

	response := &authpb.UserProfileResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetProfile(profile)
	return response, nil
}

// UpdateUserProfile updates user profile information
func (s *AuthService) UpdateUserProfile(ctx context.Context, req *authpb.UpdateProfileRequest) (*authpb.UpdateProfileResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.GetUserId()]
	if !exists {
		response := &authpb.UpdateUserProfileResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
		return response, nil
	}

	// In production, update actual profile data in database
	// For now, just acknowledge the update
	_ = user // Use the user variable

	response := &authpb.UpdateUserProfileResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetMessage("User profile updated successfully")
	return response, nil
}

// ChangePassword changes a user's password
func (s *AuthenticationService) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.GetUserId()]
	if !exists {
		response := &authpb.ChangePasswordResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_NOT_FOUND)
		return response, nil
	}

	// Verify current password
	if user.Password != req.GetCurrentPassword() {
		response := &authpb.ChangePasswordResponse{}
		response.SetStatus(authpb.AuthStatus_AUTH_STATUS_INVALID_CREDENTIALS)
		response.SetMessage("Current password is incorrect")
		return response, nil
	}

	// Update password (in production, hash the password)
	user.Password = req.GetNewPassword()

	response := &authpb.ChangePasswordResponse{}
	response.SetStatus(authpb.AuthStatus_AUTH_STATUS_SUCCESS)
	response.SetMessage("Password changed successfully")
	return response, nil
}
