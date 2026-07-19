#!/usr/bin/env python3
# file: scripts/fix-go-paths.py
# version: 2.2.0
# guid: fix-go-paths-v1-v2-script

"""
Post-buf-generate script to fix Go module paths for pb-suffixed packages.

v1 is dropped repo-wide (deprecated, no real consumers - see
docs/AUDIT-2026-07-18-post-org-migration.md Phase A) and excluded from
buf.gen.yaml's generation input, so buf generate never produces a v1/
subdirectory anymore. This script only needs to ensure each v2 module has
a go.mod.

Go's versioning rule this exists for: v2+ import paths must contain /v2+
(e.g., github.com/falkcorp/gcommon/pkg/commonpb/v2).
"""

from pathlib import Path


def fix_go_paths():
    """Ensure every pkg/*/v2 module has a go.mod after buf generate."""
    pkg_dir = Path("pkg")

    if not pkg_dir.exists():
        print("❌ pkg/ directory not found")
        return False

    print("🔧 Fixing Go module paths...")

    for module_dir in pkg_dir.iterdir():
        if not module_dir.is_dir():
            continue

        module_name = module_dir.name
        v2_dir = module_dir / "v2"

        print(f"📦 Processing module: {module_name}")

        if v2_dir.exists():
            create_go_mod_v2(v2_dir, module_name)

    print("✅ Go module path fixing complete!")
    return True


def create_go_mod_v2(v2_dir: Path, module_name: str):
    """Create go.mod for v2 module in the v2 directory."""
    go_mod_path = v2_dir / "go.mod"

    # Check if go.mod already exists - don't overwrite existing files
    if go_mod_path.exists():
        print(
            f"    ⏭️  Skipping go.mod creation (already exists): pkg/{module_name}/v2/go.mod"
        )
        return

    go_mod_content = f"""// file: pkg/{module_name}/v2/go.mod
// version: 1.0.0
// guid: go-mod-{module_name}-v2

module github.com/falkcorp/gcommon/pkg/{module_name}/v2

go 1.24

require (
\tgoogle.golang.org/grpc v1.65.0
\tgoogle.golang.org/protobuf v1.34.2
)

require (
\tgolang.org/x/net v0.25.0 // indirect
\tgolang.org/x/sys v0.20.0 // indirect
\tgolang.org/x/text v0.15.0 // indirect
\tgoogle.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)
"""

    with open(go_mod_path, "w") as f:
        f.write(go_mod_content)

    print(f"    ✅ Created go.mod for v2: pkg/{module_name}/v2/go.mod")


if __name__ == "__main__":
    success = fix_go_paths()
    exit(0 if success else 1)
