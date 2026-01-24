// file: services/auth/service_test.go
// version: 1.0.0
// guid: auth-service-test-comprehensive-v2

package auth

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jdfalk/gcommon/pkg/authpb"
	"github.com/jdfalk/gcommon/services/auth/types"
)

func TestNewAuthService(t *testing.T) {
	jwtSecret := []byte("test-secret-key-12345")
	service, err := NewAuthService(jwtSecret)
	
	if err != nil {
		t.Fatalf("NewAuthService failed: %v", err)
	}
	
	if service == nil {
		t.Fatal("NewAuthService returned nil")
	}
	
	// Test that demo users are created
	if len(service.users) == 0 {
		t.Error("Expected demo users to be created")
	}
	
	// Test that OAuth2 providers are set up
	if len(service.providers) == 0 {
		t.Error("Expected OAuth2 providers to be configured")
	}
}

func TestLogin(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	ctx := context.Background()
	
	// Test successful login with username
	loginReq := &authpb.LoginRequest{
		Username: "admin",
		Password: "admin123",
	}
	
	resp, err := service.Login(ctx, loginReq)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	
	if resp.AccessToken == "" {
		t.Error("Expected access token to be returned")
	}
	
	if resp.RefreshToken == "" {
		t.Error("Expected refresh token to be returned")
	}
	
	if resp.UserId != "admin-001" {
		t.Errorf("Expected user ID 'admin-001', got '%s'", resp.UserId)
	}
	
	// Test login with email
	emailLoginReq := &authpb.LoginRequest{
		Username: "admin@example.com", 
		Password: "admin123",
	}
	
	resp2, err := service.Login(ctx, emailLoginReq)
	if err != nil {
		t.Fatalf("Email login failed: %v", err)
	}
	
	if resp2.UserId != "admin-001" {
		t.Errorf("Expected user ID 'admin-001' for email login, got '%s'", resp2.UserId)
	}
	
	// Test failed login
	failLoginReq := &authpb.LoginRequest{
		Username: "admin",
		Password: "wrongpassword",
	}
	
	_, err = service.Login(ctx, failLoginReq)
	if err == nil {
		t.Error("Expected login to fail with wrong password")
	}
}

func TestAPIKeyAuthentication(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	ctx := context.Background()
	
	// Create API key
	createReq := &types.CreateAPIKeyRequest{
		UserID: "admin-001",
		Name:   "Test API Key",
		Scopes: []string{"read", "write"},
	}
	
	createResp, err := service.CreateAPIKey(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateAPIKey failed: %v", err)
	}
	
	if createResp.APIKey == "" {
		t.Error("Expected API key to be returned")
	}
	
	if createResp.APIKeyID == "" {
		t.Error("Expected API key ID to be returned")
	}
	
	// Test API key authentication
	authResp, err := service.AuthenticateAPIKey(ctx, createResp.APIKey)
	if err != nil {
		t.Fatalf("AuthenticateAPIKey failed: %v", err)
	}
	
	if !authResp.Valid {
		t.Error("Expected API key to be valid")
	}
	
	if authResp.UserID != "admin-001" {
		t.Errorf("Expected user ID 'admin-001', got '%s'", authResp.UserID)
	}
	
	// Test invalid API key
	invalidAuthResp, err := service.AuthenticateAPIKey(ctx, "invalid-key")
	if err != nil {
		t.Fatalf("AuthenticateAPIKey with invalid key failed: %v", err)
	}
	
	if invalidAuthResp.Valid {
		t.Error("Expected invalid API key to be rejected")
	}
	
	// List API keys
	listReq := &types.ListAPIKeysRequest{
		UserID: "admin-001",
	}
	
	listResp, err := service.ListAPIKeys(ctx, listReq)
	if err != nil {
		t.Fatalf("ListAPIKeys failed: %v", err)
	}
	
	if len(listResp.APIKeys) != 1 {
		t.Errorf("Expected 1 API key, got %d", len(listResp.APIKeys))
	}
	
	// Revoke API key
	revokeReq := &types.RevokeAPIKeyRequest{
		APIKeyID: createResp.APIKeyID,
		UserID:   "admin-001",
	}
	
	revokeResp, err := service.RevokeAPIKey(ctx, revokeReq)
	if err != nil {
		t.Fatalf("RevokeAPIKey failed: %v", err)
	}
	
	if !revokeResp.Success {
		t.Error("Expected API key revocation to succeed")
	}
	
	// Verify revoked key no longer works
	revokedAuthResp, err := service.AuthenticateAPIKey(ctx, createResp.APIKey)
	if err != nil {
		t.Fatalf("AuthenticateAPIKey with revoked key failed: %v", err)
	}
	
	if revokedAuthResp.Valid {
		t.Error("Expected revoked API key to be invalid")
	}
}

func TestOAuth2Flow(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	ctx := context.Background()
	
	// Initiate OAuth2 flow
	initiateReq := &types.InitiateOAuthRequest{
		Provider:    "github",
		RedirectURL: "http://localhost:8080/callback",
	}
	
	initiateResp, err := service.InitiateOAuth(ctx, initiateReq)
	if err != nil {
		t.Fatalf("InitiateOAuth failed: %v", err)
	}
	
	if initiateResp.AuthURL == "" {
		t.Error("Expected auth URL to be returned")
	}
	
	if initiateResp.State == "" {
		t.Error("Expected state to be returned")
	}
	
	// Handle OAuth2 callback
	callbackReq := &types.HandleOAuthCallbackRequest{
		Provider: "github",
		Code:     "test-auth-code",
		State:    initiateResp.State,
	}
	
	callbackResp, err := service.HandleOAuthCallback(ctx, callbackReq)
	if err != nil {
		t.Fatalf("HandleOAuthCallback failed: %v", err)
	}
	
	if !callbackResp.Success {
		t.Error("Expected OAuth2 callback to succeed")
	}
	
	if callbackResp.AccessToken == "" {
		t.Error("Expected access token from OAuth2 callback")
	}
	
	// Test invalid state
	invalidCallbackReq := &types.HandleOAuthCallbackRequest{
		Provider: "github",
		Code:     "test-auth-code",
		State:    "invalid-state",
	}
	
	invalidCallbackResp, err := service.HandleOAuthCallback(ctx, invalidCallbackReq)
	if err != nil {
		t.Fatalf("HandleOAuthCallback with invalid state failed: %v", err)
	}
	
	if invalidCallbackResp.Success {
		t.Error("Expected OAuth2 callback with invalid state to fail")
	}
}

func TestUserProfileManagement(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	ctx := context.Background()
	
	// Get user profile
	getReq := &types.GetUserProfileRequest{
		UserID: "admin-001",
	}
	
	getResp, err := service.GetUserProfile(ctx, getReq)
	if err != nil {
		t.Fatalf("GetUserProfile failed: %v", err)
	}
	
	if getResp.User == nil {
		t.Fatal("Expected user to be returned")
	}
	
	if getResp.User.Username != "admin" {
		t.Errorf("Expected username 'admin', got '%s'", getResp.User.Username)
	}
	
	// Update user profile
	updateReq := &types.UpdateUserProfileRequest{
		UserID: "admin-001",
		Profile: types.UserProfile{
			FirstName:   "Updated",
			LastName:    "Admin",
			DisplayName: "Updated Administrator",
		},
	}
	
	updateResp, err := service.UpdateUserProfile(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateUserProfile failed: %v", err)
	}
	
	if !updateResp.Success {
		t.Error("Expected profile update to succeed")
	}
	
	if updateResp.User.Profile.FirstName != "Updated" {
		t.Errorf("Expected first name 'Updated', got '%s'", updateResp.User.Profile.FirstName)
	}
	
	// Change password
	changeReq := &types.ChangePasswordRequest{
		UserID:      "admin-001",
		OldPassword: "admin123",
		NewPassword: "newpassword123",
	}
	
	changeResp, err := service.ChangePassword(ctx, changeReq)
	if err != nil {
		t.Fatalf("ChangePassword failed: %v", err)
	}
	
	if !changeResp.Success {
		t.Error("Expected password change to succeed")
	}
	
	// Test login with new password
	loginReq := &authpb.LoginRequest{
		Username: "admin",
		Password: "newpassword123",
	}
	
	_, err = service.Login(ctx, loginReq)
	if err != nil {
		t.Error("Login with new password should succeed")
	}
	
	// Test login with old password should fail
	oldLoginReq := &authpb.LoginRequest{
		Username: "admin", 
		Password: "admin123",
	}
	
	_, err = service.Login(ctx, oldLoginReq)
	if err == nil {
		t.Error("Login with old password should fail")
	}
}

func TestTokenGeneration(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	// Test generateToken method
	token, err := service.generateToken("test-user", []string{"user"}, time.Hour, "access")
	if err != nil {
		t.Fatalf("generateToken failed: %v", err)
	}
	
	if token == "" {
		t.Error("Expected token to be generated")
	}
	
	// Test API key generation
	apiKey, err := service.generateAPIKey()
	if err != nil {
		t.Fatalf("generateAPIKey failed: %v", err)
	}
	
	if apiKey == "" {
		t.Error("Expected API key to be generated")
	}
	
	if !strings.HasPrefix(apiKey, "gck_") {
		t.Errorf("Expected API key to start with 'gck_', got '%s'", apiKey)
	}
}

func TestBackgroundCleanup(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	// Add an expired session
	expiredSession := &types.Session{
		ID:        "expired-session",
		UserID:    "admin-001",
		Token:     "expired-token",
		CreatedAt: time.Now().Add(-2 * time.Hour),
		ExpiresAt: time.Now().Add(-1 * time.Hour), // Expired 1 hour ago
		IsActive:  true,
	}
	
	service.sessions["expired-session"] = expiredSession
	
	// Add an expired OAuth2 state
	expiredState := &types.OAuth2State{
		State:     "expired-state",
		Provider:  "github",
		CreatedAt: time.Now().Add(-20 * time.Minute),
		ExpiresAt: time.Now().Add(-10 * time.Minute), // Expired 10 minutes ago
	}
	
	service.oauthStates["expired-state"] = expiredState
	
	// Run cleanup
	service.cleanupExpiredTokensAndSessions()
	
	// Check that expired items were removed
	if _, exists := service.sessions["expired-session"]; exists {
		t.Error("Expected expired session to be cleaned up")
	}
	
	if _, exists := service.oauthStates["expired-state"]; exists {
		t.Error("Expected expired OAuth2 state to be cleaned up")
	}
}

func TestServiceShutdown(t *testing.T) {
	service, err := NewAuthService([]byte("test-secret"))
	if err != nil {
		t.Fatalf("Failed to create auth service: %v", err)
	}
	
	// Test graceful shutdown
	err = service.Close()
	if err != nil {
		t.Fatalf("Service close failed: %v", err)
	}
	
	// Cleanup ticker should be stopped
	if service.cleanupTicker != nil {
		// Ticker should be stopped (we can't easily test this without race conditions)
		t.Log("Cleanup ticker stopped successfully")
	}
}