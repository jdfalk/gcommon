#!/usr/bin/env python3
# file: scripts/create_github_issue.py
# version: 1.2.0
# guid: 12345678-90ab-cdef-1234-567890abcdef

"""
GitHub Issue Creation Script for gcommon Implementation

This script creates a comprehensive GitHub issue with all TODOs, implementation guides,
and coding standards for assignment to GitHub Copilot or other agents.

The script automatically detects GitHub tokens from:
1. GITHUB_TOKEN environment variable
2. GH_TOKEN environment variable (GitHub CLI)
3. GitHub CLI auth status (gh auth token)
4. Git config (github.token)
5. GitHub CLI config file (~/.config/gh/config.yml)

Usage:
    python3 scripts/create_github_issue.py [--dry-run] [--token TOKEN]

Requirements:
    pip install PyGithub
"""

import argparse
import os
import sys
from pathlib import Path
from typing import Optional

try:
    from github import Github
    from github.GithubException import GithubException
except ImportError:
    print("❌ PyGithub is required. Install with: pip install PyGithub")
    sys.exit(1)


def get_github_token() -> Optional[str]:
    """
    Automatically detect GitHub token from various sources.

    Checks in order:
    1. GITHUB_TOKEN environment variable
    2. GH_TOKEN environment variable (GitHub CLI)
    3. GitHub CLI auth status
    4. ~/.gitconfig for github.token
    5. ~/.config/gh/config.yml (GitHub CLI config)

    Returns:
        GitHub token if found, None otherwise
    """
    # Check environment variables
    for env_var in ["GITHUB_TOKEN", "GH_TOKEN"]:
        token = os.getenv(env_var)
        if token:
            print(f"✅ Found GitHub token from {env_var} environment variable")
            return token

    # Try GitHub CLI
    try:
        import subprocess

        result = subprocess.run(
            ["gh", "auth", "token"], capture_output=True, text=True, timeout=5
        )
        if result.returncode == 0 and result.stdout.strip():
            print("✅ Found GitHub token from GitHub CLI (gh auth token)")
            return result.stdout.strip()
    except (subprocess.SubprocessError, FileNotFoundError, ImportError):
        pass

    # Check git config
    try:
        import subprocess

        result = subprocess.run(
            ["git", "config", "--global", "github.token"],
            capture_output=True,
            text=True,
            timeout=5,
        )
        if result.returncode == 0 and result.stdout.strip():
            print("✅ Found GitHub token from git config (github.token)")
            return result.stdout.strip()
    except (subprocess.SubprocessError, FileNotFoundError, ImportError):
        pass

    # Check GitHub CLI config file
    gh_config_path = Path.home() / ".config" / "gh" / "config.yml"
    if gh_config_path.exists():
        try:
            with open(gh_config_path, "r") as f:
                content = f.read()
                if "oauth_token:" in content:
                    # Extract token from YAML-like format (basic parsing)
                    for line in content.split("\n"):
                        if "oauth_token:" in line:
                            token = line.split("oauth_token:")[1].strip().strip("\"'")
                            if token:
                                print(
                                    "✅ Found GitHub token from ~/.config/gh/config.yml"
                                )
                                return token
        except Exception:
            pass

    print("⚠️  No GitHub token found automatically")
    return None


class GitHubIssueCreator:
    """Creates comprehensive GitHub issues with TODO lists and implementation guides."""

    def __init__(self, token: str, repo_name: str, dry_run: bool = False):
        """Initialize the GitHub issue creator."""
        self.dry_run = dry_run
        self.repo_name = repo_name

        if not dry_run:
            self.github = Github(token)
            self.repo = self.github.get_repo(repo_name)

        self.repo_root = Path(__file__).parent.parent

    def load_file_content(self, file_path: str) -> str:
        """Load content from a file, with error handling."""
        full_path = self.repo_root / file_path
        try:
            with open(full_path, "r", encoding="utf-8") as f:
                content = f.read()
            return content
        except FileNotFoundError:
            print(f"⚠️  Warning: File not found: {full_path}")
            return f"<!-- File not found: {file_path} -->"
        except Exception as e:
            print(f"⚠️  Warning: Error reading {full_path}: {e}")
            return f"<!-- Error reading file: {file_path} -->"

    def generate_issue_title(self) -> str:
        """Generate the issue title."""
        return "feat(gcommon): Complete implementation of all 54 TODO items - Full service architecture"

    def generate_issue_body(self) -> str:
        """Generate the main issue body with essential sections (stays under GitHub's 65,536 char limit)."""
        sections = [
            self._generate_summary_section(),
            self._generate_priority_breakdown(),
            self._generate_todo_list_section(),
            self._generate_quick_implementation_guide(),
        ]

        return "\n\n".join(sections)

    def generate_detailed_comment(self) -> str:
        """Generate detailed comment with comprehensive sections."""
        sections = [
            "## 📋 Detailed Implementation Information",
            "",
            "This comment contains the comprehensive implementation details, coding standards, and technical requirements.",
            "",
            self._generate_implementation_guide_section(),
            self._generate_coding_standards_section(),
            self._generate_architecture_references(),
            self._generate_acceptance_criteria(),
            self._generate_technical_requirements(),
        ]

        return "\n".join(sections)

    def _generate_summary_section(self) -> str:
        """Generate the issue summary section."""
        return """## 🎯 Summary

This comprehensive implementation task covers all 54 TODO items for the gcommon services architecture. The goal is to create a complete, production-ready microservices platform following the hybrid architecture pattern established by the health service.

### 🔥 **CRITICAL PRIORITY**: Authentication Service v2 Extension
The most urgent task is expanding the AuthService from 6 methods to 18 methods using authpb/v2 protobuf definitions. This is blocking subtitle-manager integration.

### 📊 Scope Overview
- **Priority 1 (7 items)**: Authentication system with API keys, OAuth2, session management
- **Priority 2 (6 items)**: Core services (Health, Config, Web, Queue, Database, Metrics)
- **Priority 3 (2 items)**: Organization and Media services
- **Infrastructure (7 items)**: HTTP REST API, gRPC server, database, config, logging, security, service discovery
- **Testing (1 item)**: Comprehensive test infrastructure
- **Deployment (3 items)**: Docker, Kubernetes, CI/CD
- **Operations (6 items)**: Monitoring, backup, performance, security, load balancing
- **Documentation (5 items)**: API docs, architecture, development guides
- **Integration (5 items)**: Client SDKs, third-party services, webhooks, migration tools, CLI
- **Enterprise (9 items)**: Advanced caching, messaging, tracing, multi-tenancy, HA, compliance, analytics, SSO, security
- **Final Integration (3 items)**: Subtitle-manager integration, production readiness, end-to-end testing"""

    def _generate_priority_breakdown(self) -> str:
        """Generate priority breakdown section."""
        return """## 🚨 Priority Breakdown

### 🔴 **IMMEDIATE (Priority 1)**
1. **AuthService Backend Implementation** - Expand from 6 to 18 v2 methods
2. **Database Schema for Advanced Auth Features** - API keys, OAuth2, sessions, profiles
3. **API Key Authentication System** - X-API-Key header, scopes, rotation
4. **OAuth2 Provider Integration** - GitHub, Google integration flows
5. **Enhanced Session Management** - Session info, extension, listing, metadata
6. **User Profile Management System** - Profile CRUD, preferences, password changes

### 🟡 **HIGH (Priority 2)**
7. **Health Service Implementation** - System health checks, dependency status
8. **Config Service Implementation** - Centralized configuration, hot-reload
9. **Web Service Implementation** - HTTP/gRPC bridge, middleware, WebSocket
10. **Queue Service Implementation** - Message queuing, job scheduling, retry
11. **Database Service Implementation** - Connection pooling, migrations, transactions
12. **Metrics Service Implementation** - Collection, aggregation, alerting

### 🟢 **MEDIUM (Priority 3)**
13. **Organization Service Implementation** - Multi-tenancy, teams, roles
14. **Media Service Implementation** - File handling, conversion, streaming"""

    def _generate_quick_implementation_guide(self) -> str:
        """Generate quick implementation guide for main issue."""
        return """## 🚀 Quick Start Implementation Guide

### Step 1: AuthService Expansion (Priority 1)
1. **Location**: `/pkg/authpb/v2/auth.proto` (protobuf definitions already exist)
2. **Implementation**: `services/auth/service.go` - expand from 6 to 18 methods
3. **Key Methods**: Login, ValidateToken, AuthorizeAccess, GenerateToken, RefreshToken, RevokeToken, AuthenticateApiKey, CreateApiKey, RevokeApiKey, ListApiKeys, InitiateOAuth, HandleOAuthCallback, ConfigureOAuth, GetSessionInfo, ExtendSession, ListSessions, GetUserProfile, UpdateUserProfile, ChangePassword

### Step 2: Database Schema Updates
1. **API Keys Table**: Store API keys with scopes and metadata
2. **OAuth Configs**: Store OAuth2 provider configurations
3. **User Sessions**: Enhanced session tracking with metadata
4. **User Profiles**: Extended user profile information

### Step 3: Service Implementation Pattern
Follow the hybrid architecture established by the health service:
- gRPC service implementation in `services/`
- HTTP REST API endpoints in `internal/handlers/`
- Database operations in `internal/storage/`
- Configuration in `internal/config/`

**📋 See comment below for detailed implementation guide, coding standards, and technical requirements.**"""

    def _generate_todo_list_section(self) -> str:
        """Generate the complete TODO list section."""
        todo_content = self.load_file_content(
            "docs/COMPREHENSIVE_IMPLEMENTATION_TODO.md"
        )
        return f"""## 📋 Complete TODO List (54 Items)

<details>
<summary>Click to expand full TODO list</summary>

{todo_content}

</details>"""

    def _generate_implementation_guide_section(self) -> str:
        """Generate implementation guide section."""
        return """## 📖 Implementation Guide

### Reference Architecture Pattern
Follow the hybrid architecture pattern established by the existing health service implementation:
1. Complete gRPC Service Implementation using protobuf APIs
2. HTTP REST API Compatibility for backward compatibility
3. CLI Integration with cobra flags and viper configuration
4. Comprehensive Testing with unit tests covering all functionality
5. Production-Ready Features with background monitoring and graceful shutdown
6. Hybrid Architecture with internal domain types and protobuf boundary conversion

### Critical AuthService v2 Methods Missing

The authpb/v2 service needs these 12 additional methods beyond the existing 6:

7. **AuthenticateApiKey** - Validate API key authentication
8. **CreateApiKey** - Generate new API keys with scopes
9. **RevokeApiKey** - Invalidate existing API keys
10. **ListApiKeys** - List user's active API keys
11. **InitiateOAuth** - Start OAuth2 authentication flow
12. **HandleOAuthCallback** - Process OAuth2 callback responses
13. **ConfigureOAuth** - Admin OAuth2 provider management
14. **GetSessionInfo** - Retrieve detailed session information
15. **ExtendSession** - Prolong session expiration
16. **ListSessions** - List user's active sessions
17. **GetUserProfile** - Retrieve user profile and preferences
18. **UpdateUserProfile** - Update user profile data"""

    def _generate_coding_standards_section(self) -> str:
        """Generate coding standards section."""
        go_instructions = self.load_file_content(
            ".github/instructions/go.instructions.md"
        )
        protobuf_instructions = self.load_file_content(
            ".github/instructions/protobuf.instructions.md"
        )
        commit_instructions = self.load_file_content(
            ".github/instructions/commit-messages.instructions.md"
        )

        return f"""## 📏 Coding Standards & Guidelines

All implementation must strictly follow these coding standards:

### Go Language Standards
<details>
<summary>Click to expand Go coding instructions</summary>

```markdown
{go_instructions}
```

</details>

### Protobuf Standards
<details>
<summary>Click to expand Protobuf coding instructions</summary>

```markdown
{protobuf_instructions}
```

</details>

### Commit Message Standards
<details>
<summary>Click to expand Commit message instructions</summary>

```markdown
{commit_instructions}
```

</details>

### Key Requirements
- **Go Version**: Must use Go 1.23.0 or higher (MANDATORY)
- **Edition 2023**: All protobuf files must use `edition = "2023";`
- **1-1-1 Pattern**: One message/enum/service per protobuf file
- **File Headers**: All files must include path, version, and GUID headers
- **Module Prefixes**: Use consistent module prefixes to avoid naming conflicts"""

    def _generate_architecture_references(self) -> str:
        """Generate architecture references section."""
        return """## 🏗️ Architecture References

### Existing Implementations to Follow
- **Health Service**: `services/health/` - Reference implementation for all services
- **AuthService v1**: `services/auth/` - Base implementation to extend
- **Protobuf v2 APIs**: `pkg/authpb/v2/` - Generated protobuf definitions

### Generated Protobuf Packages Available (v1 dropped repo-wide, v2 only)
- `github.com/falkcorp/gcommon/pkg/authpb/v2`
- `github.com/falkcorp/gcommon/pkg/commonpb/v2`
- `github.com/falkcorp/gcommon/pkg/configpb/v2`
- `github.com/falkcorp/gcommon/pkg/databasepb/v2`
- `github.com/falkcorp/gcommon/pkg/healthpb/v2`
- `github.com/falkcorp/gcommon/pkg/mediapb/v2`
- `github.com/falkcorp/gcommon/pkg/metricspb/v2`
- `github.com/falkcorp/gcommon/pkg/organizationpb/v2`
- `github.com/falkcorp/gcommon/pkg/queuepb/v2`
- `github.com/falkcorp/gcommon/pkg/webpb/v2`

### Build & Generation Workflow
```bash
# Generate protobuf code and tidy modules
make generate

# Build and test
make build
make test

# Manual path fixing if needed
make fix-paths
```"""

    def _generate_acceptance_criteria(self) -> str:
        """Generate acceptance criteria section."""
        pr_instructions = self.load_file_content(
            ".github/instructions/pull-request-descriptions.instructions.md"
        )

        return f"""## ✅ Acceptance Criteria

### Functional Requirements
- [ ] All 54 TODO items implemented and tested
- [ ] AuthService expanded from 6 to 18 methods with full v2 API support
- [ ] All services follow the hybrid architecture pattern
- [ ] Complete database schemas for authentication features
- [ ] API key authentication system with scoped permissions
- [ ] OAuth2 integration for GitHub and Google
- [ ] Enhanced session management with metadata tracking
- [ ] User profile management with preferences and settings
- [ ] All protobuf files follow Edition 2023 and 1-1-1 pattern
- [ ] Go code follows 1.23.0+ standards and conventions

### Technical Requirements
- [ ] All code passes `go test ./...`
- [ ] All code passes `go vet ./...`
- [ ] All protobuf files pass `buf lint`
- [ ] All protobuf files generate successfully with `buf generate`
- [ ] No breaking changes to existing APIs
- [ ] Backward compatibility maintained for subtitle-manager
- [ ] All file headers include required path, version, GUID
- [ ] All commit messages follow conventional commit format
- [ ] All PRs follow standardized description template

### Pull Request Description Standards

<details>
<summary>Pull Request Description Standards</summary>

```markdown
{pr_instructions}
```

</details>

### Integration Testing
- [ ] Subtitle-manager integration tests pass
- [ ] All gRPC services respond correctly
- [ ] HTTP REST API endpoints function properly
- [ ] Database migrations execute successfully
- [ ] OAuth2 flows complete end-to-end
- [ ] API key authentication works across services"""

    def _generate_technical_requirements(self) -> str:
        """Generate technical requirements section."""
        return """## 🔧 Technical Requirements

### Development Environment
- **Go**: Version 1.23.0 or higher (MANDATORY)
- **Protobuf**: Edition 2023 support
- **Tools**: buf CLI, protoc, protoc-gen-go, protoc-gen-go-grpc
- **Dependencies**: See go.mod files in pkg/ subdirectories

### File Structure Requirements
```
gcommon/
├── services/           # Service implementations
│   ├── auth/          # AuthService v1/v2 (PRIORITY 1)
│   ├── health/        # Reference implementation
│   ├── config/        # ConfigService (Priority 2)
│   ├── web/           # WebService (Priority 2)
│   ├── queue/         # QueueService (Priority 2)
│   ├── database/      # DatabaseService (Priority 2)
│   ├── metrics/       # MetricsService (Priority 2)
│   ├── organization/  # OrganizationService (Priority 3)
│   └── media/         # MediaService (Priority 3)
├── internal/          # Infrastructure components
│   ├── auth/         # Auth database schemas & logic
│   ├── api/          # HTTP REST API layer
│   ├── server/       # gRPC server infrastructure
│   ├── db/           # Database connection management
│   ├── config/       # Configuration management
│   └── ...           # Other internal packages
├── pkg/              # Generated protobuf packages (DO NOT EDIT)
├── scripts/          # Build and automation scripts
└── docs/            # Documentation
```

### Database Requirements
- **API Keys Table**: id, user_id, key_hash, name, scopes, created_at, expires_at, last_used_at
- **OAuth Sessions Table**: id, user_id, provider, oauth_token, refresh_token, expires_at
- **Sessions Table**: id, user_id, token_hash, expires_at, extended_count, metadata
- **User Profiles Table**: id, user_id, preferences, settings, profile_data
- **Migration Support**: Schema versioning and migration scripts

### Testing Requirements
- **Unit Tests**: 80%+ coverage for all services
- **Integration Tests**: Service-to-service communication
- **End-to-End Tests**: Complete user flows
- **Performance Tests**: Load testing for critical paths
- **Security Tests**: Authentication and authorization flows

### Deployment Requirements
- **Docker**: Multi-stage builds for all services
- **Kubernetes**: Deployments, services, ingress, configmaps
- **CI/CD**: GitHub Actions workflows with automated testing
- **Monitoring**: Prometheus metrics, Grafana dashboards
- **Logging**: Structured logging with correlation IDs"""

    def create_issue(self) -> Optional[str]:
        """Create the GitHub issue with main content and detailed comment."""
        title = self.generate_issue_title()
        body = self.generate_issue_body()
        detailed_comment = self.generate_detailed_comment()

        if self.dry_run:
            print("🔍 DRY RUN - Issue would be created with:")
            print(f"Title: {title}")
            print(f"Main body length: {len(body)} characters")
            print(f"Detailed comment length: {len(detailed_comment)} characters")
            print(
                f"Total content length: {len(body) + len(detailed_comment)} characters"
            )
            print("\n" + "=" * 80)
            print("MAIN ISSUE BODY PREVIEW:")
            print("=" * 80)
            print(body[:1000] + "\n\n... [TRUNCATED FOR PREVIEW] ...")
            print("\n" + "=" * 80)
            print("DETAILED COMMENT PREVIEW:")
            print("=" * 80)
            print(detailed_comment[:1000] + "\n\n... [TRUNCATED FOR PREVIEW] ...")
            return None

        try:
            print(f"🚀 Creating issue in {self.repo_name}...")

            # Create the main issue
            issue = self.repo.create_issue(
                title=title,
                body=body,
                labels=[
                    "enhancement",
                    "epic",
                    "help wanted",
                    "good first issue",
                    "priority:high",
                ],
            )

            print(f"✅ Successfully created main issue: {issue.html_url}")

            # Add the detailed comment
            print("📋 Adding detailed implementation comment...")
            comment = issue.create_comment(detailed_comment)

            print("✅ Successfully added detailed comment")
            print(f"🔗 Main Issue: {issue.html_url}")
            print(f"🔗 Detailed Comment: {comment.html_url}")

            return issue.html_url

        except GithubException as e:
            print(f"❌ Failed to create GitHub issue: {e}")
            print(f"📊 Issue body length: {len(body)} characters")
            print(f"📊 Comment length: {len(detailed_comment)} characters")
            return None
        except Exception as e:
            print(f"❌ Unexpected error: {e}")
            return None


def main():
    """Main function."""
    parser = argparse.ArgumentParser(
        description="Create comprehensive GitHub issue for gcommon implementation"
    )
    parser.add_argument(
        "--dry-run", action="store_true", help="Preview the issue without creating it"
    )
    parser.add_argument(
        "--token", help="GitHub personal access token (or use GITHUB_TOKEN env var)"
    )
    parser.add_argument(
        "--repo", default="falkcorp/gcommon", help="Repository in owner/repo format"
    )

    args = parser.parse_args()

    # Get GitHub token automatically
    if args.token:
        token = args.token
        print("✅ Using provided GitHub token")
    else:
        token = get_github_token()

    if not token and not args.dry_run:
        print("❌ GitHub token required but not found.")
        print("💡 Options:")
        print("   1. Set GITHUB_TOKEN environment variable")
        print("   2. Set GH_TOKEN environment variable")
        print("   3. Use GitHub CLI: gh auth login")
        print("   4. Set git config: git config --global github.token YOUR_TOKEN")
        print("   5. Use --token argument: --token YOUR_TOKEN")
        sys.exit(1)

    # Create issue creator
    creator = GitHubIssueCreator(token or "", args.repo, args.dry_run)

    # Create the issue
    issue_url = creator.create_issue()

    if issue_url:
        print("\n🎉 Issue created successfully!")
        print(f"📋 URL: {issue_url}")
        print("💡 Next step: Assign this issue to GitHub Copilot for implementation")
    elif args.dry_run:
        print("\n✅ Dry run completed. Issue looks good to create!")
    else:
        print("\n❌ Failed to create issue")
        sys.exit(1)


if __name__ == "__main__":
    main()
