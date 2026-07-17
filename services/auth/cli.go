// file: services/auth/cli.go
// version: 1.0.0
// guid: g2h3i4j5-k6l7-m8n9-o0p1-q2r3s4t5u6v7

package auth

import (
	"crypto/rsa"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AddCliFlags adds auth service configuration flags to a cobra command
func AddCliFlags(cmd *cobra.Command, prefix string) {
	if prefix != "" {
		prefix = prefix + "."
	}

	// JWT Configuration
	cmd.Flags().String(prefix+"auth.jwt-secret", "", "JWT signing secret key")
	cmd.Flags().Duration(prefix+"auth.jwt-expiration", 24*time.Hour, "JWT token expiration duration")
	cmd.Flags().Duration(prefix+"auth.refresh-expiration", 7*24*time.Hour, "Refresh token expiration duration")

	// OAuth2 Configuration
	cmd.Flags().Bool(prefix+"auth.enable-oauth2", false, "Enable OAuth2 authentication")
	cmd.Flags().String(prefix+"auth.oauth2-client-id", "", "OAuth2 client ID")
	cmd.Flags().String(prefix+"auth.oauth2-client-secret", "", "OAuth2 client secret")
	cmd.Flags().String(prefix+"auth.oauth2-redirect-url", "", "OAuth2 redirect URL")

	// Security Configuration
	cmd.Flags().Bool(prefix+"auth.require-2fa", false, "Require two-factor authentication")
	cmd.Flags().Int(prefix+"auth.max-login-attempts", 5, "Maximum login attempts before lockout")
	cmd.Flags().Duration(prefix+"auth.lockout-duration", 15*time.Minute, "Account lockout duration")
	cmd.Flags().Int(prefix+"auth.password-min-length", 8, "Minimum password length")
	cmd.Flags().Bool(prefix+"auth.password-require-special", true, "Require special characters in passwords")

	// CORS Configuration
	cmd.Flags().StringSlice(prefix+"auth.allowed-origins", []string{"*"}, "Allowed CORS origins")

	// RSA Key Paths (for production)
	cmd.Flags().String(prefix+"auth.public-key-path", "", "Path to RSA public key file")
	cmd.Flags().String(prefix+"auth.private-key-path", "", "Path to RSA private key file")

	// Bind flags to viper
	viper.BindPFlag(prefix+"auth.jwt_secret", cmd.Flags().Lookup(prefix+"auth.jwt-secret"))
	viper.BindPFlag(prefix+"auth.jwt_expiration", cmd.Flags().Lookup(prefix+"auth.jwt-expiration"))
	viper.BindPFlag(prefix+"auth.refresh_expiration", cmd.Flags().Lookup(prefix+"auth.refresh-expiration"))
	viper.BindPFlag(prefix+"auth.enable_oauth2", cmd.Flags().Lookup(prefix+"auth.enable-oauth2"))
	viper.BindPFlag(prefix+"auth.oauth2_client_id", cmd.Flags().Lookup(prefix+"auth.oauth2-client-id"))
	viper.BindPFlag(prefix+"auth.oauth2_client_secret", cmd.Flags().Lookup(prefix+"auth.oauth2-client-secret"))
	viper.BindPFlag(prefix+"auth.oauth2_redirect_url", cmd.Flags().Lookup(prefix+"auth.oauth2-redirect-url"))
	viper.BindPFlag(prefix+"auth.require_2fa", cmd.Flags().Lookup(prefix+"auth.require-2fa"))
	viper.BindPFlag(prefix+"auth.max_login_attempts", cmd.Flags().Lookup(prefix+"auth.max-login-attempts"))
	viper.BindPFlag(prefix+"auth.lockout_duration", cmd.Flags().Lookup(prefix+"auth.lockout-duration"))
	viper.BindPFlag(prefix+"auth.password_min_length", cmd.Flags().Lookup(prefix+"auth.password-min-length"))
	viper.BindPFlag(prefix+"auth.password_require_special", cmd.Flags().Lookup(prefix+"auth.password-require-special"))
	viper.BindPFlag(prefix+"auth.allowed_origins", cmd.Flags().Lookup(prefix+"auth.allowed-origins"))
	viper.BindPFlag(prefix+"auth.public_key_path", cmd.Flags().Lookup(prefix+"auth.public-key-path"))
	viper.BindPFlag(prefix+"auth.private_key_path", cmd.Flags().Lookup(prefix+"auth.private-key-path"))
}

// OAuth2ConfigFromViper creates OAuth2-specific configuration
func OAuth2ConfigFromViper(prefix string) *OAuth2Config {
	if prefix != "" {
		prefix = prefix + "."
	}

	return &OAuth2Config{
		Enabled:      viper.GetBool(prefix + "auth.enable_oauth2"),
		ClientID:     viper.GetString(prefix + "auth.oauth2_client_id"),
		ClientSecret: viper.GetString(prefix + "auth.oauth2_client_secret"),
		RedirectURL:  viper.GetString(prefix + "auth.oauth2_redirect_url"),
		Scopes:       parseScopes(viper.GetString(prefix + "auth.oauth2_scopes")),
		Endpoints: OAuth2Endpoints{
			Auth:     viper.GetString(prefix + "auth.oauth2_auth_url"),
			Token:    viper.GetString(prefix + "auth.oauth2_token_url"),
			UserInfo: viper.GetString(prefix + "auth.oauth2_userinfo_url"),
		},
	}
}

// OAuth2Config represents OAuth2 configuration
type OAuth2Config struct {
	Enabled      bool
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
	Endpoints    OAuth2Endpoints
}

// OAuth2Endpoints represents OAuth2 provider endpoints
type OAuth2Endpoints struct {
	Auth     string
	Token    string
	UserInfo string
}

// SecurityConfigFromViper creates security-specific configuration
func SecurityConfigFromViper(prefix string) *SecurityConfig {
	if prefix != "" {
		prefix = prefix + "."
	}

	return &SecurityConfig{
		Require2FA:          viper.GetBool(prefix + "auth.require_2fa"),
		MaxLoginAttempts:    viper.GetInt(prefix + "auth.max_login_attempts"),
		LockoutDuration:     viper.GetDuration(prefix + "auth.lockout_duration"),
		PasswordMinLength:   viper.GetInt(prefix + "auth.password_min_length"),
		PasswordRequireSpec: viper.GetBool(prefix + "auth.password_require_special"),
		AllowedOrigins:      viper.GetStringSlice(prefix + "auth.allowed_origins"),
		EnableRateLimit:     viper.GetBool(prefix + "auth.enable_rate_limit"),
		RateLimitRequests:   viper.GetInt(prefix + "auth.rate_limit_requests"),
		RateLimitWindow:     viper.GetDuration(prefix + "auth.rate_limit_window"),
	}
}

// SecurityConfig represents security-specific configuration
type SecurityConfig struct {
	Require2FA          bool
	MaxLoginAttempts    int
	LockoutDuration     time.Duration
	PasswordMinLength   int
	PasswordRequireSpec bool
	AllowedOrigins      []string
	EnableRateLimit     bool
	RateLimitRequests   int
	RateLimitWindow     time.Duration
}

// Helper functions

func parseScopes(scopesStr string) []string {
	if scopesStr == "" {
		return []string{"openid", "profile", "email"}
	}
	return strings.Split(scopesStr, ",")
}

// loadPublicKey loads an RSA public key from file (placeholder)
func loadPublicKey(path string) *rsa.PublicKey {
	// In production, implement actual key loading
	// keyData, err := ioutil.ReadFile(path)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// block, _ := pem.Decode(keyData)
	// key, err := x509.ParsePKIXPublicKey(block.Bytes)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// return key.(*rsa.PublicKey)
	return nil
}

// loadPrivateKey loads an RSA private key from file (placeholder)
func loadPrivateKey(path string) *rsa.PrivateKey {
	// In production, implement actual key loading
	// keyData, err := ioutil.ReadFile(path)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// block, _ := pem.Decode(keyData)
	// key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// return key
	return nil
}
