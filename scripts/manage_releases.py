#!/usr/bin/env python3
# file: scripts/manage_releases.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890

"""
GitHub Release Management Script

This script manages GitHub releases for the gcommon repository by:
1. Deleting all existing releases
2. Creating new releases for all current tags
3. Matching BSR label structure with proper versioning

Usage:
    python scripts/manage_releases.py --delete-all --create-all
    python scripts/manage_releases.py --delete-all
    python scripts/manage_releases.py --create-all
    python scripts/manage_releases.py --list
"""

import os
import sys
import json
import subprocess
import argparse
from typing import List, Dict, Optional
from datetime import datetime

class GitHubReleaseManager:
    def __init__(self, repo_owner: str = "falkcorp", repo_name: str = "gcommon"):
        self.repo_owner = repo_owner
        self.repo_name = repo_name
        self.repo_full = f"{repo_owner}/{repo_name}"

    def run_gh_command(self, args: List[str]) -> subprocess.CompletedProcess:
        """Run GitHub CLI command and return result"""
        cmd = ["gh"] + args
        print(f"Running: {' '.join(cmd)}")
        try:
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                check=False
            )
            if result.returncode != 0:
                print(f"Error: {result.stderr}")
                return result
            return result
        except FileNotFoundError:
            print("Error: GitHub CLI (gh) not found. Please install it first.")
            sys.exit(1)

    def get_all_releases(self) -> List[Dict]:
        """Get all releases from the repository"""
        print("Fetching all releases...")
        result = self.run_gh_command([
            "release", "list",
            "--repo", self.repo_full,
            "--json", "tagName,name,id,createdAt",
            "--limit", "100"
        ])

        if result.returncode != 0:
            print("Failed to fetch releases")
            return []

        try:
            releases = json.loads(result.stdout)
            print(f"Found {len(releases)} existing releases")
            return releases
        except json.JSONDecodeError:
            print("Failed to parse releases JSON")
            return []

    def delete_all_releases(self) -> bool:
        """Delete all existing releases"""
        releases = self.get_all_releases()

        if not releases:
            print("No releases found to delete")
            return True

        print(f"Deleting {len(releases)} releases...")
        success_count = 0

        for release in releases:
            tag_name = release.get("tagName", "unknown")
            print(f"Deleting release: {tag_name}")

            result = self.run_gh_command([
                "release", "delete", tag_name,
                "--repo", self.repo_full,
                "--yes"  # Skip confirmation
            ])

            if result.returncode == 0:
                success_count += 1
                print(f"✅ Deleted: {tag_name}")
            else:
                print(f"❌ Failed to delete: {tag_name}")

        print(f"Successfully deleted {success_count}/{len(releases)} releases")
        return success_count == len(releases)

    def get_all_tags(self) -> List[str]:
        """Get all tags from the repository"""
        print("Fetching all tags...")
        result = subprocess.run(
            ["git", "tag", "-l"],
            capture_output=True,
            text=True,
            cwd="/Users/jdfalk/repos/github.com/falkcorp/gcommon"
        )

        if result.returncode != 0:
            print("Failed to fetch tags")
            return []

        tags = [tag.strip() for tag in result.stdout.split('\n') if tag.strip()]
        # Sort tags to prioritize patch versions, then minor, then major
        def sort_key(tag):
            if tag.count('.') == 2:  # patch version (e.g., v1.9.5)
                return (0, tag)
            elif tag.count('.') == 1:  # minor version (e.g., v1.9)
                return (1, tag)
            else:  # major version (e.g., v1)
                return (2, tag)

        tags.sort(key=sort_key)
        print(f"Found {len(tags)} tags: {tags}")
        return tags

    def generate_release_notes(self, tag: str) -> str:
        """Generate release notes for a tag"""
        current_date = datetime.now().strftime("%Y-%m-%d")

        if tag.count('.') == 2:  # Patch version
            if 'v1.' in tag:
                return f"""# gcommon {tag}

Protocol buffer definitions and Go SDK for common types and services.

## Fixed in this release
- ✅ Corrected `go_package` declarations (removed trailing slashes)
- ✅ Fixed Go import path issues for all modules
- ✅ Updated BSR registry with corrected protobuf definitions

## What's included
- Common types and utilities
- Configuration management
- Database abstractions
- Health check definitions
- Media processing types
- Metrics collection
- Organization management
- Queue management
- Web service definitions

## Installation
```go
go get github.com/falkcorp/gcommon/pkg/commonpb/v2@{tag}
```

## BSR Registry
```bash
buf dep update buf.build/falkcorp/gcommon:{tag}
```

---
*Released: {current_date}*
*BSR: buf.build/falkcorp/gcommon*"""

            else:  # v2 patch
                return f"""# gcommon {tag} (v2 Series)

Next-generation protocol buffer definitions with enhanced features.

## Fixed in this release
- ✅ Corrected `go_package` declarations (removed trailing slashes)
- ✅ Fixed Go import path issues for all v2 modules
- ✅ Updated BSR registry with corrected protobuf definitions

## V2 Features
- Enhanced type safety
- Improved validation rules
- Backward compatibility with v1
- Performance optimizations

## Installation
```go
go get github.com/falkcorp/gcommon/pkg/commonpb/v2@{tag}
```

## BSR Registry
```bash
buf dep update buf.build/falkcorp/gcommon:{tag}
```

---
*Released: {current_date}*
*BSR: buf.build/falkcorp/gcommon*"""

        elif tag.count('.') == 1:  # Minor version
            series = "v1" if 'v1.' in tag else "v2"
            return f"""# gcommon {tag} ({series.upper()} Series)

Latest {series} series release with all recent fixes and improvements.

## What's Fixed
- ✅ All import path issues resolved
- ✅ BSR registry synchronized
- ✅ Go module compatibility restored

## Quick Start
```go
go get github.com/falkcorp/gcommon/pkg/commonpb/v2@{tag}
```

---
*Released: {current_date}*"""

        else:  # Major version
            series = "v1" if tag == "v1" else "v2"
            return f"""# gcommon {tag} (Latest {series.upper()})

Latest {series} release with all fixes and features.

## Installation
```go
go get github.com/falkcorp/gcommon/pkg/commonpb/v2@{tag}
```

---
*Released: {current_date}*"""

    def create_release_for_tag(self, tag: str) -> bool:
        """Create a release for a specific tag"""
        print(f"Creating release for tag: {tag}")

        release_notes = self.generate_release_notes(tag)
        release_title = f"gcommon {tag}"

        # Determine if this should be a prerelease
        is_prerelease = False  # All our tags are stable releases

        result = self.run_gh_command([
            "release", "create", tag,
            "--repo", self.repo_full,
            "--title", release_title,
            "--notes", release_notes,
            "--latest" if tag in ["v1", "v2"] else ""
        ])

        if result.returncode == 0:
            print(f"✅ Created release: {tag}")
            return True
        else:
            print(f"❌ Failed to create release: {tag}")
            return False

    def create_all_releases(self) -> bool:
        """Create releases for all tags"""
        tags = self.get_all_tags()

        if not tags:
            print("No tags found to create releases for")
            return True

        print(f"Creating releases for {len(tags)} tags...")
        success_count = 0

        for tag in tags:
            if self.create_release_for_tag(tag):
                success_count += 1

        print(f"Successfully created {success_count}/{len(tags)} releases")
        return success_count == len(tags)

    def list_releases(self):
        """List all current releases"""
        releases = self.get_all_releases()

        if not releases:
            print("No releases found")
            return

        print(f"\nCurrent releases ({len(releases)}):")
        print("-" * 60)
        for release in releases:
            tag = release.get("tagName", "unknown")
            name = release.get("name", "unknown")
            created = release.get("createdAt", "unknown")
            print(f"  {tag:10} | {name:25} | {created}")

def main():
    parser = argparse.ArgumentParser(description="Manage GitHub releases for gcommon")
    parser.add_argument("--delete-all", action="store_true", help="Delete all existing releases")
    parser.add_argument("--create-all", action="store_true", help="Create releases for all tags")
    parser.add_argument("--list", action="store_true", help="List current releases")
    parser.add_argument("--repo", default="falkcorp/gcommon", help="Repository (owner/name)")

    args = parser.parse_args()

    if not any([args.delete_all, args.create_all, args.list]):
        parser.print_help()
        return

    repo_parts = args.repo.split("/")
    if len(repo_parts) != 2:
        print("Error: Repository must be in format 'owner/name'")
        return

    manager = GitHubReleaseManager(repo_parts[0], repo_parts[1])

    try:
        if args.list:
            manager.list_releases()

        if args.delete_all:
            print("🗑️  Deleting all existing releases...")
            if not manager.delete_all_releases():
                print("Failed to delete all releases")
                return
            print("✅ All releases deleted successfully")

        if args.create_all:
            print("🚀 Creating releases for all tags...")
            if not manager.create_all_releases():
                print("Failed to create all releases")
                return
            print("✅ All releases created successfully")

        print("\n🎉 Release management completed successfully!")

    except KeyboardInterrupt:
        print("\n❌ Operation cancelled by user")
    except Exception as e:
        print(f"❌ Unexpected error: {e}")

if __name__ == "__main__":
    main()
