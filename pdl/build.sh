#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
bin_dir="${root_dir}/bin"
build_dir="${root_dir}/pdlc/build"

mkdir -p "${bin_dir}"

echo ">>> building pdl"
GOOS="${GOOS:-}" GOARCH="${GOARCH:-}" go build -C "${root_dir}" -o "${bin_dir}/pdl" ./cmd/pdl

echo ">>> building pdlbuild"
GOOS="${GOOS:-}" GOARCH="${GOARCH:-}" go build -C "${root_dir}" -o "${bin_dir}/pdlbuild" ./cmd/pdlbuild

echo ">>> building pdlgen"
GOOS="${GOOS:-}" GOARCH="${GOARCH:-}" go build -C "${root_dir}" -o "${bin_dir}/pdlgen" ./cmd/pdlgen

echo ">>> building db2pdl"
GOOS="${GOOS:-}" GOARCH="${GOARCH:-}" go build -C "${root_dir}/pdl-orm" -o "${bin_dir}/db2pdl" ./cmd/db2pdl

echo ">>> building pdlc"
cmake_args=()
if [[ -n "${CMAKE_GENERATOR:-}" ]]; then
  cmake_args+=(-G "${CMAKE_GENERATOR}")
fi
"${root_dir}/pdlc/build.sh" "${cmake_args[@]}"
echo ">>> done"
