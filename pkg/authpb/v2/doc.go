// file: pkg/authpb/v2/doc.go
// version: 1.0.0
// guid: doc-authpb-v2

// Package authpb provides v2 authentication protocol buffer definitions.
//
// This package contains enhanced authentication and authorization services with:
//
// Core Authentication:
//   - Login with username/password
//   - Token validation and refresh
//   - Authorization and access control
//
// API Key Management:
//   - API key creation and authentication
//   - Scoped permissions and access control
//   - Key lifecycle management (create, list, revoke)
//
// OAuth2 Integration:
//   - OAuth2 flow initiation and callback handling
//   - Support for multiple providers (GitHub, Google, etc.)
//   - Configuration management for OAuth providers
//
// Session Management:
//   - Session creation and tracking
//   - Session extension and expiration
//   - Multi-device session management
//
// User Profile Management:
//   - Comprehensive user profile retrieval
//   - Profile updates with validation
//   - Secure password change functionality
//
// The v2 API provides backward compatibility with v1 while offering
// significantly enhanced functionality and improved field structures.
//
// Migration from v1:
//
//	The v2 API includes all v1 methods with improved implementations
//	plus 12 additional methods for enhanced functionality.
package authv2
