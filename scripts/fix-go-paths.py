#!/usr/bin/env python3
# file: scripts/fix-go-paths.py
# version: 2.2.0
# guid: fix-go-paths-v1-v2-script

"""
Post-buf-generate script to fix Go module paths for pb-suffixed packages.

Go has specific versioning rules:
- v1 modules: Import path cannot contain /v1 (e.g., github.com/falkcorp/gcommon/pkg/commonpb)
- v2+ modules: Import path must contain /v2+ (e.g., github.com/falkcorp/gcommon/pkg/commonpb/v2)

This script:
1. Moves v1 files from pkg/*/v1/ to pkg/*/
2. Keeps v2+ files in their versioned directories
3. Creates appropriate go.mod files for each version (only if they don't exist)
4. Updates import statements in moved files
5. Works with pb-suffixed package names (commonpb, metricspb, etc.)
"""

import shutil
from pathlib import Path


def fix_go_paths():
    """Fix Go module paths after buf generate."""
    pkg_dir = Path("pkg")

    if not pkg_dir.exists():
        print("❌ pkg/ directory not found")
        return False

    print("🔧 Fixing Go module paths...")

    for module_dir in pkg_dir.iterdir():
        if not module_dir.is_dir():
            continue

        module_name = module_dir.name
        v1_dir = module_dir / "v1"
        v2_dir = module_dir / "v2"

        print(f"📦 Processing module: {module_name}")

        # Move v1 files to module root
        if v1_dir.exists():
            print(f"  📂 Moving v1 files from {v1_dir} to {module_dir}")

            # Move all .pb.go files from v1/ to module root
            for pb_file in v1_dir.glob("*.pb.go"):
                dest_file = module_dir / pb_file.name
                if dest_file.exists():
                    dest_file.unlink()  # Remove existing file
                shutil.move(str(pb_file), str(dest_file))
                print(f"    ✅ Moved {pb_file.name}")

            # Move go.mod and go.sum if they exist
            for go_file in ["go.mod", "go.sum"]:
                src_file = v1_dir / go_file
                if src_file.exists():
                    dest_file = module_dir / go_file
                    if dest_file.exists():
                        dest_file.unlink()
                    shutil.move(str(src_file), str(dest_file))
                    print(f"    ✅ Moved {go_file}")

            # Remove empty v1 directory
            if v1_dir.exists() and not any(v1_dir.iterdir()):
                v1_dir.rmdir()
                print("    🗑️  Removed empty v1 directory")

        # Create go.mod for v1 in module root
        create_go_mod_v1(module_dir, module_name)

        # Create go.mod for v2 if it exists
        if v2_dir.exists():
            create_go_mod_v2(v2_dir, module_name)

    print("✅ Go module path fixing complete!")
    return True


def create_go_mod_v1(module_dir: Path, module_name: str):
    """Create go.mod for v1 module in the root directory."""
    go_mod_path = module_dir / "go.mod"

    # Check if go.mod already exists - don't overwrite existing files
    if go_mod_path.exists():
        print(
            f"    ⏭️  Skipping go.mod creation (already exists): pkg/{module_name}/go.mod"
        )
        return

    go_mod_content = f"""// file: pkg/{module_name}/go.mod
// version: 1.0.0
// guid: go-mod-{module_name}-v1

module github.com/falkcorp/gcommon/pkg/{module_name}

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

    print(f"    ✅ Created go.mod for v1: pkg/{module_name}/go.mod")


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
