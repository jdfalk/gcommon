#!/usr/bin/env python3
# file: scripts/release-manager.py
# version: 2.0.0
# guid: b9f7c8d3-2a4e-4c5b-8f1a-9e6d7b2c1a3f
# last-edited: 2026-07-19

"""
Automated Release Management Script

This script handles the complete release workflow:
1. Finds current released version and latest tags
2. Ensures repository is clean and on main branch
3. Calculates next version (patch/minor/major)
4. Creates and pushes tags
5. Generates GitHub releases with changelogs
6. Runs go mod tidy on all modules

Usage:
    python3 scripts/release-manager.py [patch|minor|major]
    python3 scripts/release-manager.py patch  # Default
"""

import argparse
import json
import logging
import re
import subprocess
import sys
from datetime import datetime
from pathlib import Path
from typing import Dict, List, Optional, Tuple

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger(__name__)


class ReleaseManager:
    """Manages the complete release workflow."""

    def __init__(self, repo_path: Path):
        self.repo_path = repo_path
        self.current_version = None
        self.next_version = None

    def run_command(self, cmd: List[str], check: bool = True, capture_output: bool = True) -> subprocess.CompletedProcess:
        """Run a command and return the result."""
        logger.debug(f"Running command: {' '.join(cmd)}")
        try:
            result = subprocess.run(
                cmd,
                cwd=self.repo_path,
                check=check,
                capture_output=capture_output,
                text=True
            )
            return result
        except subprocess.CalledProcessError as e:
            logger.error(f"Command failed: {' '.join(cmd)}")
            logger.error(f"Error: {e.stderr}")
            if check:
                sys.exit(1)
            return e

    def ensure_clean_main_branch(self) -> bool:
        """Ensure we're on main branch and repository is clean."""
        logger.info("🔍 Checking repository state...")

        # Check current branch
        result = self.run_command(["git", "branch", "--show-current"])
        current_branch = result.stdout.strip()

        if current_branch != "main":
            logger.error(f"❌ Not on main branch (currently on: {current_branch})")
            logger.info("💡 Please switch to main branch: git checkout main")
            return False

        # Check for uncommitted changes
        result = self.run_command(["git", "status", "--porcelain"])
        if result.stdout.strip():
            logger.error("❌ Repository has uncommitted changes")
            logger.info("💡 Please commit or stash changes before releasing")
            return False

        # Check if we're ahead of origin
        result = self.run_command(["git", "status", "--porcelain=v1", "--branch"])
        status_lines = result.stdout.strip().split('\n')
        if any('[ahead' in line for line in status_lines):
            logger.warning("⚠️  Local branch is ahead of origin")
            logger.info("💡 Consider pushing changes: git push origin main")

        logger.info("✅ Repository is clean and on main branch")
        return True

    def get_current_version(self) -> Optional[str]:
        """Get the current/latest version from git tags."""
        logger.info("🔍 Finding current version...")

        try:
            # Get all tags sorted by version
            result = self.run_command(["git", "tag", "-l", "--sort=-version:refname"])
            tags = [tag.strip() for tag in result.stdout.split('\n') if tag.strip()]

            # Filter to semantic version tags
            version_pattern = re.compile(r'^v?\d+\.\d+\.\d+$')
            version_tags = [tag for tag in tags if version_pattern.match(tag)]

            if not version_tags:
                logger.info("No semantic version tags found, starting with v0.1.0")
                return "v0.0.0"  # Will be incremented to v0.1.0

            latest_tag = version_tags[0]
            logger.info(f"📋 Current version: {latest_tag}")
            self.current_version = latest_tag
            return latest_tag

        except subprocess.CalledProcessError:
            logger.warning("Could not determine current version, starting with v0.1.0")
            return "v0.0.0"

    def calculate_next_version(self, bump_type: str) -> str:
        """Calculate the next version based on bump type."""
        if not self.current_version:
            self.current_version = self.get_current_version()

        # Parse current version
        version_match = re.match(r'^v?(\d+)\.(\d+)\.(\d+)$', self.current_version)
        if not version_match:
            logger.error(f"Invalid version format: {self.current_version}")
            sys.exit(1)

        major, minor, patch = map(int, version_match.groups())

        if bump_type == "major":
            major += 1
            minor = 0
            patch = 0
        elif bump_type == "minor":
            minor += 1
            patch = 0
        elif bump_type == "patch":
            patch += 1
        else:
            logger.error(f"Invalid bump type: {bump_type}. Use: patch, minor, or major")
            sys.exit(1)

        next_version = f"v{major}.{minor}.{patch}"
        logger.info(f"📈 Next version: {next_version} ({bump_type} bump)")
        self.next_version = next_version
        return next_version

    def run_go_mod_tidy(self) -> bool:
        """Run go mod tidy on the single repo-wide module.

        gcommon is one Go module rooted at github.com/falkcorp/gcommon/v2 -
        see docs/AUDIT-2026-07-18-post-org-migration.md Phase D. There used
        to be a go.mod per pkg/*pb/v2 family plus internal/services/
        services/auth, which required walking the tree; that's gone now.
        """
        logger.info("🔧 Running go mod tidy...")

        result = subprocess.run(
            ["go", "mod", "tidy"],
            cwd=self.repo_path,
            check=False,
            capture_output=True,
            text=True
        )

        if result.returncode == 0:
            logger.info("    ✅ go mod tidy successful")
            return True

        logger.error("    ❌ go mod tidy failed")
        if result.stderr:
            logger.error(f"    Error: {result.stderr.strip()}")
        return False

    def generate_changelog(self) -> str:
        """Generate changelog for the release."""
        logger.info("📝 Generating changelog...")

        try:
            # Get commits since last tag
            if self.current_version and self.current_version != "v0.0.0":
                result = self.run_command([
                    "git", "log", f"{self.current_version}..HEAD",
                    "--oneline", "--no-merges"
                ])
            else:
                result = self.run_command([
                    "git", "log", "--oneline", "--no-merges", "-10"
                ])

            commits = [line.strip() for line in result.stdout.split('\n') if line.strip()]

            if not commits:
                return "## Changes\n\n- Minor updates and improvements"

            # Categorize commits
            features = []
            fixes = []
            chores = []

            for commit in commits:
                if commit.startswith(('feat:', 'feature:')):
                    features.append(commit)
                elif commit.startswith(('fix:', 'bugfix:')):
                    fixes.append(commit)
                else:
                    chores.append(commit)

            changelog_parts = ["## Changes\n"]

            if features:
                changelog_parts.append("### ✨ Features")
                for feature in features:
                    changelog_parts.append(f"- {feature}")
                changelog_parts.append("")

            if fixes:
                changelog_parts.append("### 🐛 Bug Fixes")
                for fix in fixes:
                    changelog_parts.append(f"- {fix}")
                changelog_parts.append("")

            if chores:
                changelog_parts.append("### 🔧 Other Changes")
                for chore in chores[:5]:  # Limit to 5 most recent
                    changelog_parts.append(f"- {chore}")
                changelog_parts.append("")

            return "\n".join(changelog_parts)

        except subprocess.CalledProcessError:
            return "## Changes\n\n- Package updates and improvements"

    def _create_and_push_one_tag(self, tag_name: str, message: str) -> bool:
        """Create one annotated tag and push it to origin."""
        result = self.run_command(["git", "tag", "-a", tag_name, "-m", message], check=False)
        if result.returncode != 0:
            logger.error(f"Failed to create tag: {tag_name}")
            return False

        result = self.run_command(["git", "push", "origin", tag_name], check=False)
        if result.returncode != 0:
            logger.error(f"Failed to push tag: {tag_name}")
            return False

        logger.info(f"    ✅ {tag_name}")
        return True

    def create_and_push_tags(self) -> bool:
        """Create and push the single repo-wide release tag.

        gcommon is one Go module (github.com/falkcorp/gcommon/v2) rooted at
        the repo root - see docs/AUDIT-2026-07-18-post-org-migration.md
        Phase D. There is deliberately no per-submodule nested tagging
        anymore: a prior version of this script tagged root with a real
        vX.Y.Z alongside identically-versioned nested pkg/*pb/v2 tags, which
        collided under Go's module-path-major-version rule because root's
        go.mod had no /v2 suffix of its own (root@vX.Y.Z resolved as
        +incompatible and broke resolution for the nested tags too). Do not
        reintroduce nested/submodule tags without re-reading that finding.
        """
        if not self.next_version:
            logger.error("Next version not calculated")
            return False

        logger.info(f"🏷️  Creating tag: {self.next_version}")
        if not self._create_and_push_one_tag(self.next_version, f"Release {self.next_version}"):
            return False

        logger.info(f"✅ Successfully created and pushed tag: {self.next_version}")
        return True

    def create_github_release(self, changelog: str) -> bool:
        """Create GitHub release using gh CLI."""
        if not self.next_version:
            logger.error("Next version not calculated")
            return False

        logger.info(f"🚀 Creating GitHub release: {self.next_version}")

        # Check if gh CLI is available
        try:
            self.run_command(["gh", "--version"])
        except subprocess.CalledProcessError:
            logger.error("GitHub CLI (gh) not found. Please install: https://cli.github.com/")
            return False

        # Create release
        release_title = f"Release {self.next_version}"

        try:
            # Create release with changelog
            result = self.run_command([
                "gh", "release", "create", self.next_version,
                "--title", release_title,
                "--notes", changelog,
                "--latest"
            ], check=False)

            if result.returncode != 0:
                logger.error(f"Failed to create GitHub release: {result.stderr}")
                return False

            logger.info(f"✅ Successfully created GitHub release: {self.next_version}")
            return True

        except subprocess.CalledProcessError as e:
            logger.error(f"Failed to create GitHub release: {e}")
            return False

    def run_release_workflow(self, bump_type: str = "patch") -> bool:
        """Run the complete release workflow."""
        logger.info("🚀 Starting automated release workflow...")
        logger.info(f"📅 Release date: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

        # Step 1: Ensure clean main branch
        if not self.ensure_clean_main_branch():
            return False

        # Step 2: Get current version and calculate next
        self.get_current_version()
        self.calculate_next_version(bump_type)

        # Step 3: Run go mod tidy
        if not self.run_go_mod_tidy():
            logger.warning("go mod tidy had some issues, but continuing...")

        # Step 4: Generate changelog
        changelog = self.generate_changelog()
        logger.info("📝 Generated changelog:")
        print("\n" + changelog + "\n")

        # Step 5: Create and push tags
        if not self.create_and_push_tags():
            return False

        # Step 6: Create GitHub release
        if not self.create_github_release(changelog):
            logger.warning("GitHub release creation failed, but tags were created successfully")

        logger.info("🎉 Release workflow completed successfully!")
        logger.info(f"📦 Released version: {self.next_version}")

        return True


def main():
    """Main function."""
    parser = argparse.ArgumentParser(
        description="Automated release management for gcommon repository"
    )
    parser.add_argument(
        "bump_type",
        nargs="?",
        choices=["patch", "minor", "major"],
        default="patch",
        help="Type of version bump (default: patch)"
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Show what would be done without making changes"
    )
    parser.add_argument(
        "--verbose",
        action="store_true",
        help="Enable verbose logging"
    )

    args = parser.parse_args()

    if args.verbose:
        logging.getLogger().setLevel(logging.DEBUG)

    # Get repository path
    script_dir = Path(__file__).parent
    repo_path = script_dir.parent

    logger.info(f"🏠 Repository: {repo_path}")
    logger.info(f"📊 Bump type: {args.bump_type}")

    if args.dry_run:
        logger.info("🔍 DRY RUN MODE - No changes will be made")
        # TODO: Implement dry run mode
        logger.warning("Dry run mode not yet implemented")
        return 1

    # Create release manager and run workflow
    release_manager = ReleaseManager(repo_path)
    success = release_manager.run_release_workflow(args.bump_type)

    return 0 if success else 1


if __name__ == "__main__":
    sys.exit(main())
