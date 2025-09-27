// file: services/auth/service.go  
// version: 1.1.0
// guid: n4o5p6q7-r8s9-0123-7890-234567890123

package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jdfalk/gcommon/services/auth/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationService implements the auth service with hybrid architecture
// Supports all v1 and v2 authentication methods following the established pattern
type AuthenticationService struct {
	// Core configuration (internal domain types)
	mu            sync.RWMutex
	jwtSecret     []byte
	rsaPrivKey    *rsa.PrivateKey
	rsaPubKey     *rsa.PublicKey
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
	
	// In-memory stores (in production, use proper database)
	users         map[string]*types.User         // userID -> User
	usersByEmail  map[string]*types.User         // email -> User  
	usersByName   map[string]*types.User         // username -> User
	apiKeys       map[string]*types.APIKey       // keyHash -> APIKey
	sessions      map[string]*types.Session      // sessionID -> Session
	oauthStates   map[string]*types.OAuth2State  // state -> OAuth2State
	providers     map[string]*types.OAuth2Provider // provider -> OAuth2Provider
	
	// Background cleanup
	cleanupTicker *time.Ticker
	shutdownCh    chan struct{}
}

// NewAuthService creates a new authentication service with full v2 capabilities
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
		tokenExpiry:   time.Hour,
		refreshExpiry: time.Hour * 24 * 7, // 7 days
		
		// Initialize in-memory stores
		users:        make(map[string]*types.User),
		usersByEmail: make(map[string]*types.User),
		usersByName:  make(map[string]*types.User),
		apiKeys:      make(map[string]*types.APIKey),
		sessions:     make(map[string]*types.Session),
		oauthStates:  make(map[string]*types.OAuth2State),
		providers:    make(map[string]*types.OAuth2Provider),
		
		shutdownCh:    make(chan struct{}),
	}

	// Add demo users with proper password hashing
	if err := service.createDemoUsers(); err != nil {
		return nil, fmt.Errorf("failed to create demo users: %w", err)
	}
	
	// Setup demo OAuth2 providers
	service.setupDemoOAuth2Providers()

	// Start background cleanup goroutine
	service.startBackgroundTasks()

	return service, nil
}

// createDemoUsers creates demo users with proper bcrypt password hashing
func (s *AuthenticationService) createDemoUsers() error {
	now := time.Now()
	
	// Admin user
	adminPwHash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}
	
	adminUser := &types.User{
		ID:           "admin-001",
		Username:     "admin",
		Email:        "admin@example.com",
		PasswordHash: string(adminPwHash),
		Roles:        []string{"admin", "user"},
		Profile: &types.UserProfile{
			FirstName:   "Admin",
			LastName:    "User",
			DisplayName: "Administrator",
			Language:    "en",
			Timezone:    "UTC",
		},
		CreatedAt: now,
		UpdatedAt: now,
		IsActive:  true,
	}
	
	s.users[adminUser.ID] = adminUser
	s.usersByEmail[adminUser.Email] = adminUser
	s.usersByName[adminUser.Username] = adminUser
	
	// Regular user
	userPwHash, err := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash user password: %w", err)
	}
	
	regularUser := &types.User{
		ID:           "user-001",
		Username:     "user",
		Email:        "user@example.com",
		PasswordHash: string(userPwHash),
		Roles:        []string{"user"},
		Profile: &types.UserProfile{
			FirstName:   "Regular",
			LastName:    "User",
			DisplayName: "User",
			Language:    "en",
			Timezone:    "UTC",
		},
		CreatedAt: now,
		UpdatedAt: now,
		IsActive:  true,
	}
	
	s.users[regularUser.ID] = regularUser
	s.usersByEmail[regularUser.Email] = regularUser  
	s.usersByName[regularUser.Username] = regularUser

	return nil
}

// setupDemoOAuth2Providers creates demo OAuth2 provider configurations
func (s *AuthenticationService) setupDemoOAuth2Providers() {
	// GitHub provider
	s.providers["github"] = &types.OAuth2Provider{
		Name:        "github",
		ClientID:    "demo-github-client-id",
		ClientSecret: "demo-github-client-secret",
		AuthURL:     "https://github.com/login/oauth/authorize",
		TokenURL:    "https://github.com/login/oauth/access_token",
		UserInfoURL: "https://api.github.com/user",
		Scopes:      []string{"read:user", "user:email"},
		RedirectURL: "http://localhost:8080/auth/callback/github",
		IsActive:    true,
	}
	
	// Google provider
	s.providers["google"] = &types.OAuth2Provider{
		Name:        "google",
		ClientID:    "demo-google-client-id",
		ClientSecret: "demo-google-client-secret",
		AuthURL:     "https://accounts.google.com/oauth2/auth",
		TokenURL:    "https://oauth2.googleapis.com/token",
		UserInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo",
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8080/auth/callback/google",
		IsActive:    true,
	}
}

// startBackgroundTasks starts cleanup and monitoring tasks
func (s *AuthenticationService) startBackgroundTasks() {
	s.cleanupTicker = time.NewTicker(15 * time.Minute) // Cleanup every 15 minutes
	
	go func() {
		for {
			select {
			case <-s.cleanupTicker.C:
				s.cleanupExpiredTokensAndSessions()
			case <-s.shutdownCh:
				s.cleanupTicker.Stop()
				return
			}
		}
	}()
}

// cleanupExpiredTokensAndSessions removes expired sessions and oauth states
func (s *AuthenticationService) cleanupExpiredTokensAndSessions() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	now := time.Now()
	
	// Clean up expired sessions
	for sessionID, session := range s.sessions {
		if now.After(session.ExpiresAt) {
			delete(s.sessions, sessionID)
		}
	}
	
	// Clean up expired OAuth2 states
	for state, oauthState := range s.oauthStates {
		if now.After(oauthState.ExpiresAt) {
			delete(s.oauthStates, state)
		}
	}
	
	log.Printf("Auth service cleanup completed at %v", now)
}

// Graceful shutdown
func (s *AuthenticationService) Close() error {
	close(s.shutdownCh)
	if s.cleanupTicker != nil {
		s.cleanupTicker.Stop()
	}
	return nil
}

// ==========================================
// PROTOBUF-DEPENDENT METHODS (TODO: Uncomment when protobuf issues are resolved)
// ==========================================
/*
// Login authenticates a user and returns JWT tokens (enhanced v2 implementation)
// TODO: Uncomment when protobuf issues are resolved
func (s *AuthenticationService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// Find user by username or email
	var user *types.User
	if strings.Contains(req.GetUsername(), "@") {
		user = s.usersByEmail[req.GetUsername()]
	} else {
		user = s.usersByName[req.GetUsername()]
	}
	
	if user == nil || !user.IsActive {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Verify password with bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.GetPassword())); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Generate tokens
	accessToken, err := s.generateToken(user.ID, user.Roles, s.tokenExpiry, "access")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := s.generateToken(user.ID, user.Roles, s.refreshExpiry, "refresh")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}
	
	// Update last login time
	now := time.Now()
	user.LastLoginAt = &now
	user.UpdatedAt = now

	return &authpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.tokenExpiry.Seconds()),
		UserId:       user.ID,
		Roles:        user.Roles,
	}, nil
}
*/

// Login with internal types (works without protobuf)
func (s *AuthenticationService) LoginInternal(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// Find user by username or email
	var user *types.User
	if strings.Contains(req.Username, "@") {
		user = s.usersByEmail[req.Username]
	} else {
		user = s.usersByName[req.Username]
	}
	
	if user == nil || !user.IsActive {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Verify password with bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate tokens
	accessToken, err := s.generateToken(user.ID, user.Roles, s.tokenExpiry, "access")
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := s.generateToken(user.ID, user.Roles, s.refreshExpiry, "refresh")
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}
	
	// Update last login time
	now := time.Now()
	user.LastLoginAt = &now
	user.UpdatedAt = now

	return &types.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.tokenExpiry.Seconds()),
	}, nil
}

// All the other protobuf methods are commented out for now
/*
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	// Convert internal response to protobuf
	return &authpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.tokenExpiry.Seconds()),
	}, nil
}

// ValidateToken validates a JWT token and returns user claims
func (s *AuthenticationService) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	// Convert protobuf to internal type
	internalReq := &types.ValidateTokenRequest{
		Token: req.GetToken(),
	}

	// Parse and validate token
	token, err := jwt.Parse(internalReq.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.rsaPubKey, nil
	})

	if err != nil || !token.Valid {
		return &authpb.ValidateTokenResponse{
			Valid:  false,
			UserId: "",
			Roles:  []string{},
		}, nil
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &authpb.ValidateTokenResponse{
			Valid:  false,
			UserId: "",
			Roles:  []string{},
		}, nil
	}

	userId, _ := claims["user_id"].(string)
	rolesInterface, _ := claims["roles"].([]interface{})
	roles := make([]string, len(rolesInterface))
	for i, role := range rolesInterface {
		roles[i] = role.(string)
	}

	return &authpb.ValidateTokenResponse{
		Valid:  true,
		UserId: userId,
		Roles:  roles,
	}, nil
}

// AuthorizeAccess checks if a user has permission for a specific action
func (s *AuthenticationService) AuthorizeAccess(ctx context.Context, req *authpb.AuthorizeRequest) (*authpb.AuthorizeResponse, error) {
	// First validate the token
	validateReq := &authpb.ValidateTokenRequest{Token: req.GetToken()}
	validateResp, err := s.ValidateToken(ctx, validateReq)
	if err != nil {
		return nil, err
	}

	if !validateResp.Valid {
		return &authpb.AuthorizeResponse{
			Authorized: false,
			Reason:     "invalid token",
			Roles:      []string{},
		}, nil
	}

	// Check role-based authorization
	userRoles := validateResp.Roles
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

	return &authpb.AuthorizeResponse{
		Authorized: authorized,
		Reason:     reason,
		Roles:      userRoles,
	}, nil
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

	return &authpb.GenerateTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(expiry.Seconds()),
		TokenType:    req.GetTokenType(),
	}, nil
}

// RefreshToken exchanges a refresh token for new access tokens
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	// Validate refresh token
	validateReq := &authpb.ValidateTokenRequest{Token: req.GetRefreshToken()}
	validateResp, err := s.ValidateToken(ctx, validateReq)
	if err != nil {
		return nil, err
	}

	if !validateResp.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	// Generate new tokens
	accessToken, err := s.generateToken(validateResp.UserId, validateResp.Roles, s.tokenExpiry, "access")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := s.generateToken(validateResp.UserId, validateResp.Roles, s.refreshExpiry, "refresh")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &authpb.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.tokenExpiry.Seconds()),
	}, nil
}

// RevokeToken invalidates a token making it unusable
func (s *AuthenticationService) RevokeToken(ctx context.Context, req *authpb.RevokeTokenRequest) (*authpb.RevokeTokenResponse, error) {
	// In a real implementation, you'd add the token to a blacklist
	// or store revoked tokens in a database

	// For now, just validate the token exists
	validateReq := &authpb.ValidateTokenRequest{Token: req.GetToken()}
	validateResp, err := s.ValidateToken(ctx, validateReq)
	if err != nil {
		return nil, err
	}

	if !validateResp.Valid {
		return &authpb.RevokeTokenResponse{
			Success: false,
			Message: "token is already invalid or expired",
		}, nil
	}

	// In production: add to blacklist, update database, etc.
	return &authpb.RevokeTokenResponse{
		Success: true,
		Message: "token revoked successfully",
	}, nil
}
*/

// ===========================================
// NEW V2 METHODS - API KEY AUTHENTICATION  
// ===========================================

// AuthenticateAPIKey validates an API key and returns user information
func (s *AuthenticationService) AuthenticateAPIKey(ctx context.Context, req *types.APIKeyAuthRequest) (*types.APIKeyAuthResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// Hash the provided API key to match stored hash
	keyHash := s.hashAPIKey(req.APIKey)
	
	apiKey, exists := s.apiKeys[keyHash]
	if !exists || !apiKey.IsActive {
		return &types.APIKeyAuthResponse{
			Valid:   false,
			Message: "invalid or inactive API key",
		}, nil
	}
	
	// Check expiration
	if apiKey.ExpiresAt != nil && time.Now().After(*apiKey.ExpiresAt) {
		return &types.APIKeyAuthResponse{
			Valid:   false,
			Message: "API key has expired",
		}, nil
	}
	
	// Update last used time
	now := time.Now()
	apiKey.LastUsedAt = &now
	
	return &types.APIKeyAuthResponse{
		Valid:  true,
		UserID: apiKey.UserID,
		Scopes: apiKey.Scopes,
	}, nil
}

// CreateAPIKey generates a new API key for a user
func (s *AuthenticationService) CreateAPIKey(ctx context.Context, req *types.CreateAPIKeyRequest) (*types.CreateAPIKeyResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Verify user exists
	user, exists := s.users[req.UserID]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	
	// Generate API key
	apiKeyValue, err := s.generateAPIKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate API key: %w", err)
	}
	
	// Create API key record
	now := time.Now()
	apiKey := &types.APIKey{
		ID:        generateID("ak"),
		UserID:    user.ID,
		Name:      req.Name,
		KeyHash:   s.hashAPIKey(apiKeyValue),
		Scopes:    req.Scopes,
		CreatedAt: now,
		ExpiresAt: req.ExpiresAt,
		IsActive:  true,
		Metadata:  req.Metadata,
	}
	
	// Store the API key
	s.apiKeys[apiKey.KeyHash] = apiKey
	
	return &types.CreateAPIKeyResponse{
		APIKeyID:  apiKey.ID,
		APIKey:    apiKeyValue, // Only returned once!
		Name:      apiKey.Name,
		Scopes:    apiKey.Scopes,
		ExpiresAt: apiKey.ExpiresAt,
		CreatedAt: apiKey.CreatedAt,
	}, nil
}

// RevokeAPIKey invalidates an API key
func (s *AuthenticationService) RevokeAPIKey(ctx context.Context, req *types.RevokeAPIKeyRequest) (*types.RevokeAPIKeyResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Find and revoke the API key
	for keyHash, apiKey := range s.apiKeys {
		if apiKey.ID == req.APIKeyID && apiKey.UserID == req.UserID {
			delete(s.apiKeys, keyHash)
			return &types.RevokeAPIKeyResponse{
				Success: true,
				Message: "API key revoked successfully",
			}, nil
		}
	}
	
	return &types.RevokeAPIKeyResponse{
		Success: false,
		Message: "API key not found or access denied",
	}, nil
}

// ListAPIKeys returns all API keys for a user (without the actual key values)
func (s *AuthenticationService) ListAPIKeys(ctx context.Context, req *types.ListAPIKeysRequest) (*types.ListAPIKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var userAPIKeys []types.APIKey
	for _, apiKey := range s.apiKeys {
		if apiKey.UserID == req.UserID {
			// Don't return the key hash
			sanitizedKey := *apiKey
			sanitizedKey.KeyHash = ""
			userAPIKeys = append(userAPIKeys, sanitizedKey)
		}
	}
	
	// Apply pagination
	start := int(req.Offset)
	limit := int(req.Limit)
	if limit == 0 {
		limit = 50 // Default limit
	}
	
	totalCount := int64(len(userAPIKeys))
	end := start + limit
	if end > len(userAPIKeys) {
		end = len(userAPIKeys)
	}
	
	var paginatedKeys []types.APIKey
	if start < len(userAPIKeys) {
		paginatedKeys = userAPIKeys[start:end]
	}
	
	return &types.ListAPIKeysResponse{
		APIKeys:    paginatedKeys,
		TotalCount: totalCount,
		HasMore:    end < len(userAPIKeys),
	}, nil
}

// ===========================================
// NEW V2 METHODS - OAUTH2 AUTHENTICATION
// ===========================================

// InitiateOAuth starts an OAuth2 authentication flow
func (s *AuthenticationService) InitiateOAuth(ctx context.Context, req *types.InitiateOAuthRequest) (*types.InitiateOAuthResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	provider, exists := s.providers[req.Provider]
	if !exists || !provider.IsActive {
		return nil, fmt.Errorf("OAuth2 provider '%s' not found or inactive", req.Provider)
	}
	
	// Generate state for security
	state := generateSecureState()
	
	// Store OAuth2 state
	s.oauthStates[state] = &types.OAuth2State{
		State:       state,
		Provider:    req.Provider,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(10 * time.Minute), // 10 minute expiry
		RedirectURL: req.RedirectURL,
	}
	
	// Build authorization URL
	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&state=%s&response_type=code",
		provider.AuthURL,
		url.QueryEscape(provider.ClientID),
		url.QueryEscape(provider.RedirectURL),
		url.QueryEscape(strings.Join(provider.Scopes, " ")),
		url.QueryEscape(state))
	
	return &types.InitiateOAuthResponse{
		AuthURL: authURL,
		State:   state,
	}, nil
}

// HandleOAuthCallback processes OAuth2 callback and completes authentication
func (s *AuthenticationService) HandleOAuthCallback(ctx context.Context, req *types.HandleOAuthCallbackRequest) (*types.HandleOAuthCallbackResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Verify state
	oauthState, exists := s.oauthStates[req.State]
	if !exists || time.Now().After(oauthState.ExpiresAt) {
		return &types.HandleOAuthCallbackResponse{
			Success: false,
			Message: "invalid or expired OAuth2 state",
		}, nil
	}
	
	// Clean up state
	delete(s.oauthStates, req.State)
	
	if oauthState.Provider != req.Provider {
		return &types.HandleOAuthCallbackResponse{
			Success: false,
			Message: "provider mismatch",
		}, nil
	}
	
	provider := s.providers[req.Provider]
	
	// In a real implementation, you would:
	// 1. Exchange code for access token with OAuth2 provider
	// 2. Get user info from OAuth2 provider  
	// 3. Create or update user record
	// 4. Generate JWT tokens for the user
	
	// For demo purposes, create a demo OAuth2 user
	userID := fmt.Sprintf("oauth-%s-%s", req.Provider, generateID("user"))
	
	// Generate tokens for the OAuth2 user
	accessToken, err := s.generateToken(userID, []string{"user"}, s.tokenExpiry, "access")
	if err != nil {
		return &types.HandleOAuthCallbackResponse{
			Success: false,
			Message: "failed to generate access token",
		}, nil
	}
	
	refreshToken, err := s.generateToken(userID, []string{"user"}, s.refreshExpiry, "refresh")
	if err != nil {
		return &types.HandleOAuthCallbackResponse{
			Success: false,
			Message: "failed to generate refresh token",
		}, nil
	}
	
	return &types.HandleOAuthCallbackResponse{
		Success:      true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
		Message:      fmt.Sprintf("OAuth2 authentication successful with %s", req.Provider),
	}, nil
}

// ===========================================
// NEW V2 METHODS - SESSION MANAGEMENT
// ===========================================

// GetSessionInfo retrieves detailed session information
func (s *AuthenticationService) GetSessionInfo(ctx context.Context, req *types.GetSessionInfoRequest) (*types.GetSessionInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	session, exists := s.sessions[req.SessionToken]
	if !exists {
		return &types.GetSessionInfoResponse{
			Valid:   false,
			Message: "session not found",
		}, nil
	}
	
	// Check if session is expired
	if time.Now().After(session.ExpiresAt) {
		return &types.GetSessionInfoResponse{
			Valid:   false,
			Message: "session has expired",
		}, nil
	}
	
	// Update last seen
	session.LastSeenAt = time.Now()
	
	return &types.GetSessionInfoResponse{
		Session: session,
		Valid:   true,
	}, nil
}

// ExtendSession prolongs a session's expiration time
func (s *AuthenticationService) ExtendSession(ctx context.Context, req *types.ExtendSessionRequest) (*types.ExtendSessionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	session, exists := s.sessions[req.SessionToken]
	if !exists {
		return &types.ExtendSessionResponse{
			Success: false,
			Message: "session not found",
		}, nil
	}
	
	// Check if session is still valid
	if time.Now().After(session.ExpiresAt) {
		return &types.ExtendSessionResponse{
			Success: false,
			Message: "session has already expired",
		}, nil
	}
	
	// Extend session
	session.ExpiresAt = session.ExpiresAt.Add(req.ExtendBy)
	session.LastSeenAt = time.Now()
	
	return &types.ExtendSessionResponse{
		Success:   true,
		ExpiresAt: session.ExpiresAt,
		Message:   "session extended successfully",
	}, nil
}

// ListSessions returns all active sessions for a user
func (s *AuthenticationService) ListSessions(ctx context.Context, req *types.ListSessionsRequest) (*types.ListSessionsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var userSessions []types.Session
	now := time.Now()
	
	for _, session := range s.sessions {
		if session.UserID == req.UserID && now.Before(session.ExpiresAt) {
			// Don't return the actual session token
			sanitizedSession := *session
			sanitizedSession.Token = ""
			userSessions = append(userSessions, sanitizedSession)
		}
	}
	
	// Apply pagination
	start := int(req.Offset)
	limit := int(req.Limit)
	if limit == 0 {
		limit = 50 // Default limit
	}
	
	totalCount := int64(len(userSessions))
	end := start + limit
	if end > len(userSessions) {
		end = len(userSessions)
	}
	
	var paginatedSessions []types.Session
	if start < len(userSessions) {
		paginatedSessions = userSessions[start:end]
	}
	
	return &types.ListSessionsResponse{
		Sessions:   paginatedSessions,
		TotalCount: totalCount,
		HasMore:    end < len(userSessions),
	}, nil
}

// ===========================================
// NEW V2 METHODS - USER PROFILE MANAGEMENT
// ===========================================

// GetUserProfile retrieves user profile information
func (s *AuthenticationService) GetUserProfile(ctx context.Context, req *types.GetUserProfileRequest) (*types.GetUserProfileResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	user, exists := s.users[req.UserID]
	if !exists {
		return &types.GetUserProfileResponse{
			Message: "user not found",
		}, nil
	}
	
	// Don't return password hash
	sanitizedUser := *user
	sanitizedUser.PasswordHash = ""
	
	return &types.GetUserProfileResponse{
		User: &sanitizedUser,
	}, nil
}

// UpdateUserProfile updates user profile information
func (s *AuthenticationService) UpdateUserProfile(ctx context.Context, req *types.UpdateUserProfileRequest) (*types.UpdateUserProfileResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	user, exists := s.users[req.UserID]
	if !exists {
		return &types.UpdateUserProfileResponse{
			Success: false,
			Message: "user not found",
		}, nil
	}
	
	// Update profile fields
	if user.Profile == nil {
		user.Profile = &types.UserProfile{}
	}
	
	user.Profile.FirstName = req.Profile.FirstName
	user.Profile.LastName = req.Profile.LastName
	user.Profile.DisplayName = req.Profile.DisplayName
	user.Profile.Avatar = req.Profile.Avatar
	user.Profile.Timezone = req.Profile.Timezone
	user.Profile.Language = req.Profile.Language
	user.Profile.Preferences = req.Profile.Preferences
	user.UpdatedAt = time.Now()
	
	// Don't return password hash
	sanitizedUser := *user
	sanitizedUser.PasswordHash = ""
	
	return &types.UpdateUserProfileResponse{
		Success: true,
		Message: "profile updated successfully",
		User:    &sanitizedUser,
	}, nil
}

// ChangePassword updates a user's password
func (s *AuthenticationService) ChangePassword(ctx context.Context, req *types.ChangePasswordRequest) (*types.ChangePasswordResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	user, exists := s.users[req.UserID]
	if !exists {
		return &types.ChangePasswordResponse{
			Success: false,
			Message: "user not found",
		}, nil
	}
	
	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)); err != nil {
		return &types.ChangePasswordResponse{
			Success: false,
			Message: "current password is incorrect",
		}, nil
	}
	
	// Hash new password
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return &types.ChangePasswordResponse{
			Success: false,
			Message: "failed to hash new password",
		}, nil
	}
	
	// Update password
	user.PasswordHash = string(newPasswordHash)
	user.UpdatedAt = time.Now()
	
	return &types.ChangePasswordResponse{
		Success: true,
		Message: "password updated successfully",
	}, nil
}

// ===========================================
// UTILITY METHODS
// ===========================================

// generateAPIKey creates a cryptographically secure API key
func (s *AuthenticationService) generateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return "gck_" + hex.EncodeToString(bytes), nil // gck = gcommon key
}

// hashAPIKey creates a hash of the API key for storage
func (s *AuthenticationService) hashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}

// generateSecureState creates a secure state parameter for OAuth2
func generateSecureState() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// generateID creates a unique ID with a prefix
func generateID(prefix string) string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return prefix + "_" + hex.EncodeToString(bytes)[:16]
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
// TODO: Uncomment when protobuf issues are resolved
/*
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
*/

// StartHTTPOnly starts only HTTP server without gRPC dependencies
func (s *AuthenticationService) StartHTTPOnly(httpPort int) error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "auth"}`))
	})

	log.Printf("Auth HTTP server listening on port %d", httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
}

// HTTP handlers for basic REST API
// TODO: Uncomment when protobuf issues are resolved
/*
/*
func (s *AuthenticationService) httpLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req authpb.LoginRequest
	// In production, properly parse JSON request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	req.Username = username
	req.Password = password

	resp, err := s.Login(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// In production, use proper JSON marshaling
	fmt.Fprintf(w, `{"access_token":"%s","refresh_token":"%s","expires_in":%d}`,
		resp.AccessToken, resp.RefreshToken, resp.ExpiresIn)
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

	req := &authpb.ValidateTokenRequest{Token: token}
	resp, err := s.ValidateToken(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if resp.Valid {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"valid":true,"user_id":"%s","roles":%v}`,
			resp.UserId, resp.Roles)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"valid":false}`))
	}
}
*/