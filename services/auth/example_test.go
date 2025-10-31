// file: services/auth/example_test.go
// version: 1.0.0
// guid: auth-service-example-comprehensive-v2

package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jdfalk/gcommon/pkg/authpb"
	"github.com/jdfalk/gcommon/services/auth/types"
)

// ExampleAuthService demonstrates all authentication capabilities
func ExampleAuthService() {
	// Create a new authentication service
	jwtSecret := []byte("demo-jwt-secret-key-for-testing")
	service, err := NewAuthService(jwtSecret)
	if err != nil {
		log.Fatalf("Failed to create auth service: %v", err)
	}
	defer service.Close()

	ctx := context.Background()

	fmt.Println("=== GCommon AuthService v2 - Complete Authentication Demo ===")
	fmt.Println()

	// 1. Traditional Login
	fmt.Println("1. Traditional Username/Password Login:")
	loginReq := &authpb.LoginRequest{
		Username: "admin",
		Password: "admin123",
	}

	loginResp, err := service.Login(ctx, loginReq)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	fmt.Printf("   ✓ Login successful for user: %s\n", loginResp.UserId)
	fmt.Printf("   ✓ Access token: %s...\n", loginResp.AccessToken[:20])
	fmt.Printf("   ✓ User roles: %v\n", loginResp.Roles)
	fmt.Println()

	// 2. API Key Authentication
	fmt.Println("2. API Key Authentication:")
	createAPIKeyReq := &types.CreateAPIKeyRequest{
		UserID: "admin-001",
		Name:   "Production API Key",
		Scopes: []string{"read", "write", "admin"},
	}

	apiKeyResp, err := service.CreateAPIKey(ctx, createAPIKeyReq)
	if err != nil {
		log.Fatalf("API key creation failed: %v", err)
	}

	fmt.Printf("   ✓ Created API key: %s...\n", apiKeyResp.APIKey[:20])
	fmt.Printf("   ✓ API key ID: %s\n", apiKeyResp.APIKeyID)
	fmt.Printf("   ✓ Scopes: %v\n", apiKeyResp.Scopes)

	// Test API key authentication
	apiAuthResp, err := service.AuthenticateAPIKey(ctx, apiKeyResp.APIKey)
	if err != nil {
		log.Fatalf("API key authentication failed: %v", err)
	}

	fmt.Printf("   ✓ API key authentication: Valid=%v, UserID=%s\n", apiAuthResp.Valid, apiAuthResp.UserID)
	fmt.Println()

	// 3. OAuth2 Flow
	fmt.Println("3. OAuth2 Authentication Flow:")
	oauthReq := &types.InitiateOAuthRequest{
		Provider:    "github",
		RedirectURL: "http://localhost:8080/callback",
	}

	oauthResp, err := service.InitiateOAuth(ctx, oauthReq)
	if err != nil {
		log.Fatalf("OAuth2 initiation failed: %v", err)
	}

	fmt.Printf("   ✓ OAuth2 auth URL: %s\n", oauthResp.AuthURL[:60]+"...")
	fmt.Printf("   ✓ OAuth2 state: %s...\n", oauthResp.State[:16])

	// Simulate OAuth2 callback
	callbackReq := &types.HandleOAuthCallbackRequest{
		Provider: "github",
		Code:     "demo-authorization-code",
		State:    oauthResp.State,
	}

	callbackResp, err := service.HandleOAuthCallback(ctx, callbackReq)
	if err != nil {
		log.Fatalf("OAuth2 callback failed: %v", err)
	}

	fmt.Printf("   ✓ OAuth2 callback successful: %s\n", callbackResp.Message)
	fmt.Printf("   ✓ OAuth2 user ID: %s\n", callbackResp.UserID)
	fmt.Println()

	// 4. User Profile Management
	fmt.Println("4. User Profile Management:")
	getProfileReq := &types.GetUserProfileRequest{
		UserID: "admin-001",
	}

	profileResp, err := service.GetUserProfile(ctx, getProfileReq)
	if err != nil {
		log.Fatalf("Get profile failed: %v", err)
	}

	fmt.Printf("   ✓ User: %s (%s)\n", profileResp.User.Username, profileResp.User.Email)
	fmt.Printf("   ✓ Display name: %s\n", profileResp.User.Profile.DisplayName)
	fmt.Printf("   ✓ Account created: %s\n", profileResp.User.CreatedAt.Format("2006-01-02"))

	// Update profile
	updateProfileReq := &types.UpdateUserProfileRequest{
		UserID: "admin-001",
		Profile: types.UserProfile{
			FirstName:   "Super",
			LastName:    "Administrator",
			DisplayName: "Super Admin",
			Language:    "en",
			Timezone:    "America/New_York",
		},
	}

	updateResp, err := service.UpdateUserProfile(ctx, updateProfileReq)
	if err != nil {
		log.Fatalf("Update profile failed: %v", err)
	}

	fmt.Printf("   ✓ Profile updated: %s\n", updateResp.User.Profile.DisplayName)
	fmt.Println()

	// 5. Session Management (simulate sessions)
	fmt.Println("5. Session Management:")
	
	// Create a demo session
	sessionID := "demo-session-12345"
	session := &types.Session{
		ID:        sessionID,
		UserID:    "admin-001",
		Token:     "session-token-xyz",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		LastSeenAt: time.Now(),
		IPAddress: "192.168.1.100",
		UserAgent: "Mozilla/5.0 (Demo Browser)",
		IsActive:  true,
	}
	
	service.mu.Lock()
	service.sessions[sessionID] = session
	service.mu.Unlock()

	// Get session info
	sessionInfoReq := &types.GetSessionInfoRequest{
		SessionToken: sessionID,
	}

	sessionInfoResp, err := service.GetSessionInfo(ctx, sessionInfoReq)
	if err != nil {
		log.Fatalf("Get session info failed: %v", err)
	}

	fmt.Printf("   ✓ Session valid: %v\n", sessionInfoResp.Valid)
	fmt.Printf("   ✓ Session expires: %s\n", sessionInfoResp.Session.ExpiresAt.Format("2006-01-02 15:04"))
	fmt.Printf("   ✓ IP address: %s\n", sessionInfoResp.Session.IPAddress)

	// List user sessions
	listSessionsReq := &types.ListSessionsRequest{
		UserID: "admin-001",
	}

	listSessionsResp, err := service.ListSessions(ctx, listSessionsReq)
	if err != nil {
		log.Fatalf("List sessions failed: %v", err)
	}

	fmt.Printf("   ✓ Active sessions: %d\n", len(listSessionsResp.Sessions))
	fmt.Println()

	// 6. List API Keys
	fmt.Println("6. API Key Management:")
	listKeysReq := &types.ListAPIKeysRequest{
		UserID: "admin-001",
	}

	listKeysResp, err := service.ListAPIKeys(ctx, listKeysReq)
	if err != nil {
		log.Fatalf("List API keys failed: %v", err)
	}

	fmt.Printf("   ✓ Total API keys: %d\n", len(listKeysResp.APIKeys))
	for i, key := range listKeysResp.APIKeys {
		fmt.Printf("   %d. %s (ID: %s, Scopes: %v)\n", i+1, key.Name, key.ID, key.Scopes)
	}
	fmt.Println()

	// 7. Token Validation (existing functionality)
	fmt.Println("7. Token Validation:")
	validateReq := &authpb.ValidateTokenRequest{
		Token: loginResp.AccessToken,
	}

	validateResp, err := service.ValidateToken(ctx, validateReq)
	if err != nil {
		log.Fatalf("Token validation failed: %v", err)
	}

	fmt.Printf("   ✓ Token valid: %v\n", validateResp.Valid)
	fmt.Printf("   ✓ User ID: %s\n", validateResp.UserId)
	fmt.Printf("   ✓ Roles: %v\n", validateResp.Roles)
	fmt.Println()

	fmt.Println("=== AuthService v2 Demo Complete ===")
	fmt.Println()
	fmt.Println("✅ Implemented Features:")
	fmt.Println("   • Traditional username/password authentication")
	fmt.Println("   • API key authentication with scopes")
	fmt.Println("   • OAuth2 flow support (GitHub, Google)")
	fmt.Println("   • Enhanced session management")
	fmt.Println("   • User profile management")
	fmt.Println("   • Password change functionality")
	fmt.Println("   • JWT token generation and validation")
	fmt.Println("   • Background cleanup of expired tokens")
	fmt.Println("   • Production-ready security (bcrypt, secure key generation)")
	fmt.Println("   • Comprehensive error handling and logging")
	fmt.Println()
	fmt.Println("🚀 Ready for subtitle-manager integration!")
}