// file: services/auth/types/types.go
// version: 1.1.0
// guid: k6l7m8n9-o0p1-q2r3-s4t5-u6v7w8x9y0z1

package types

import (
	"context"
	"time"
)

// LoginRequest represents a user login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents a login response with tokens
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// ValidateTokenRequest represents a token validation request
type ValidateTokenRequest struct {
	Token string `json:"token"`
}

// ValidateTokenResponse represents a token validation response
type ValidateTokenResponse struct {
	Valid  bool     `json:"valid"`
	UserId string   `json:"user_id"`
	Roles  []string `json:"roles"`
}

// AuthorizeRequest represents an authorization request
type AuthorizeRequest struct {
	Token        string   `json:"token"`
	RequiredRole string   `json:"required_role"`
	Resource     string   `json:"resource"`
	Action       string   `json:"action"`
	Permissions  []string `json:"permissions"`
}

// AuthorizeResponse represents an authorization response
type AuthorizeResponse struct {
	Authorized bool     `json:"authorized"`
	Reason     string   `json:"reason"`
	Roles      []string `json:"roles"`
}

// GenerateTokenRequest represents a token generation request
type GenerateTokenRequest struct {
	UserId      string   `json:"user_id"`
	Roles       []string `json:"roles"`
	ExpiresIn   int64    `json:"expires_in"`
	TokenType   string   `json:"token_type"`
	Permissions []string `json:"permissions"`
}

// GenerateTokenResponse represents a token generation response
type GenerateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// RefreshTokenRequest represents a token refresh request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenResponse represents a token refresh response
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// RevokeTokenRequest represents a token revocation request
type RevokeTokenRequest struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

// RevokeTokenResponse represents a token revocation response
type RevokeTokenResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// AuthService defines the authentication service interface
type AuthService interface {
	ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error)
	AuthorizeAccess(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error)
	GenerateToken(ctx context.Context, req *GenerateTokenRequest) (*GenerateTokenResponse, error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error)
	RevokeToken(ctx context.Context, req *RevokeTokenRequest) (*RevokeTokenResponse, error)
}

// Role represents user roles
type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

// User represents a user in the system (internal domain type)
type User struct {
	ID           string            `json:"id"`
	Username     string            `json:"username"`
	Email        string            `json:"email"`
	PasswordHash string            `json:"-"` // Never serialize password
	Roles        []string          `json:"roles"`
	Profile      *UserProfile      `json:"profile,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	LastLoginAt  *time.Time        `json:"last_login_at,omitempty"`
	IsActive     bool              `json:"is_active"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// UserProfile represents additional user profile information
type UserProfile struct {
	FirstName   string            `json:"first_name,omitempty"`
	LastName    string            `json:"last_name,omitempty"`
	DisplayName string            `json:"display_name,omitempty"`
	Avatar      string            `json:"avatar,omitempty"`
	Timezone    string            `json:"timezone,omitempty"`
	Language    string            `json:"language,omitempty"`
	Preferences map[string]string `json:"preferences,omitempty"`
}

// APIKey represents an API key for authentication (internal domain type)
type APIKey struct {
	ID          string            `json:"id"`
	UserID      string            `json:"user_id"`
	Name        string            `json:"name"`
	KeyHash     string            `json:"-"` // Never serialize key hash
	Scopes      []string          `json:"scopes"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
	LastUsedAt  *time.Time        `json:"last_used_at,omitempty"`
	IsActive    bool              `json:"is_active"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// Session represents a user session (internal domain type)
type Session struct {
	ID           string            `json:"id"`
	UserID       string            `json:"user_id"`
	Token        string            `json:"-"` // Never serialize session token
	CreatedAt    time.Time         `json:"created_at"`
	ExpiresAt    time.Time         `json:"expires_at"`
	LastSeenAt   time.Time         `json:"last_seen_at"`
	IPAddress    string            `json:"ip_address,omitempty"`
	UserAgent    string            `json:"user_agent,omitempty"`
	IsActive     bool              `json:"is_active"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// OAuth2Provider represents OAuth2 provider configuration (internal domain type)
type OAuth2Provider struct {
	Name         string            `json:"name"`
	ClientID     string            `json:"client_id"`
	ClientSecret string            `json:"-"` // Never serialize client secret
	AuthURL      string            `json:"auth_url"`
	TokenURL     string            `json:"token_url"`
	UserInfoURL  string            `json:"user_info_url"`
	Scopes       []string          `json:"scopes"`
	RedirectURL  string            `json:"redirect_url"`
	IsActive     bool              `json:"is_active"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// OAuth2State represents OAuth2 flow state (internal domain type) 
type OAuth2State struct {
	State       string    `json:"state"`
	Provider    string    `json:"provider"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	RedirectURL string    `json:"redirect_url,omitempty"`
}

// Additional request/response types for expanded functionality

// API Key Authentication Types
type APIKeyAuthRequest struct {
	APIKey string `json:"api_key"`
}

type APIKeyAuthResponse struct {
	Valid   bool     `json:"valid"`
	UserID  string   `json:"user_id,omitempty"`
	Scopes  []string `json:"scopes,omitempty"`
	Message string   `json:"message,omitempty"`
}

type CreateAPIKeyRequest struct {
	UserID    string            `json:"user_id"`
	Name      string            `json:"name"`
	Scopes    []string          `json:"scopes"`
	ExpiresAt *time.Time        `json:"expires_at,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type CreateAPIKeyResponse struct {
	APIKeyID string    `json:"api_key_id"`
	APIKey   string    `json:"api_key"` // Only returned once
	Name     string    `json:"name"`
	Scopes   []string  `json:"scopes"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type RevokeAPIKeyRequest struct {
	APIKeyID string `json:"api_key_id"`
	UserID   string `json:"user_id"`
}

type RevokeAPIKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ListAPIKeysRequest struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit,omitempty"`
	Offset int32  `json:"offset,omitempty"`
}

type ListAPIKeysResponse struct {
	APIKeys    []APIKey `json:"api_keys"`
	TotalCount int64    `json:"total_count"`
	HasMore    bool     `json:"has_more"`
}

// OAuth2 Authentication Types
type InitiateOAuthRequest struct {
	Provider    string `json:"provider"`
	RedirectURL string `json:"redirect_url,omitempty"`
}

type InitiateOAuthResponse struct {
	AuthURL string `json:"auth_url"`
	State   string `json:"state"`
}

type HandleOAuthCallbackRequest struct {
	Provider string `json:"provider"`
	Code     string `json:"code"`
	State    string `json:"state"`
}

type HandleOAuthCallbackResponse struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	Message      string `json:"message,omitempty"`
}

// Session Management Types  
type GetSessionInfoRequest struct {
	SessionToken string `json:"session_token"`
}

type GetSessionInfoResponse struct {
	Session *Session `json:"session,omitempty"`
	Valid   bool     `json:"valid"`
	Message string   `json:"message,omitempty"`
}

type ExtendSessionRequest struct {
	SessionToken string        `json:"session_token"`
	ExtendBy     time.Duration `json:"extend_by"`
}

type ExtendSessionResponse struct {
	Success   bool      `json:"success"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	Message   string    `json:"message,omitempty"`
}

type ListSessionsRequest struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit,omitempty"`
	Offset int32  `json:"offset,omitempty"`
}

type ListSessionsResponse struct {
	Sessions   []Session `json:"sessions"`
	TotalCount int64     `json:"total_count"`
	HasMore    bool      `json:"has_more"`
}

// User Profile Management Types
type GetUserProfileRequest struct {
	UserID string `json:"user_id"`
}

type GetUserProfileResponse struct {
	User    *User  `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
}

type UpdateUserProfileRequest struct {
	UserID  string      `json:"user_id"`
	Profile UserProfile `json:"profile"`
}

type UpdateUserProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	User    *User  `json:"user,omitempty"`
}

type ChangePasswordRequest struct {
	UserID      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Extended AuthService interface with all authentication methods
type ExtendedAuthService interface {
	// Existing v1 methods
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error)
	AuthorizeAccess(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error)
	GenerateToken(ctx context.Context, req *GenerateTokenRequest) (*GenerateTokenResponse, error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error)
	RevokeToken(ctx context.Context, req *RevokeTokenRequest) (*RevokeTokenResponse, error)

	// New v2 methods - API Key Authentication
	AuthenticateAPIKey(ctx context.Context, req *APIKeyAuthRequest) (*APIKeyAuthResponse, error)
	CreateAPIKey(ctx context.Context, req *CreateAPIKeyRequest) (*CreateAPIKeyResponse, error)
	RevokeAPIKey(ctx context.Context, req *RevokeAPIKeyRequest) (*RevokeAPIKeyResponse, error)
	ListAPIKeys(ctx context.Context, req *ListAPIKeysRequest) (*ListAPIKeysResponse, error)

	// New v2 methods - OAuth2 Authentication
	InitiateOAuth(ctx context.Context, req *InitiateOAuthRequest) (*InitiateOAuthResponse, error)
	HandleOAuthCallback(ctx context.Context, req *HandleOAuthCallbackRequest) (*HandleOAuthCallbackResponse, error)

	// New v2 methods - Session Management  
	GetSessionInfo(ctx context.Context, req *GetSessionInfoRequest) (*GetSessionInfoResponse, error)
	ExtendSession(ctx context.Context, req *ExtendSessionRequest) (*ExtendSessionResponse, error)
	ListSessions(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error)

	// New v2 methods - User Profile Management
	GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*GetUserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *UpdateUserProfileRequest) (*UpdateUserProfileResponse, error)
	ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*ChangePasswordResponse, error)
}
