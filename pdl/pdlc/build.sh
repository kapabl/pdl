#!/usr/bin/env bash
set -euo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
build_dir="${project_dir}/build"
vcpkg_dir="${project_dir}/vcpkg"
toolchain="${vcpkg_dir}/scripts/buildsystems/vcpkg.cmake"
bin_dir="${project_dir}/../bin"

if [[ ! -d "${vcpkg_dir}" ]]; then
  git clone https://github.com/microsoft/vcpkg.git "${vcpkg_dir}"
fi

if [[ ! -f "${vcpkg_dir}/vcpkg" ]]; then
  (cd "${vcpkg_dir}" && ./bootstrap-vcpkg.sh -disableMetrics)
fi

triplet="${VCPKG_TRIPLET:-x64-linux}"
install_root="${project_dir}/vcpkg_installed"
(cd "${project_dir}" && env VCPKG_ROOT="${vcpkg_dir}" "${vcpkg_dir}/vcpkg" install --triplet "${triplet}" --disable-metrics --x-install-root="${install_root}")

export VCPKG_ROOT="${vcpkg_dir}"
export VCPKG_TARGET_TRIPLET="${triplet}"
export VCPKG_INSTALLED_DIR="${install_root}"

configure_args=("$@")
should_configure=0
if [[ ! -f "${build_dir}/CMakeCache.txt" ]]; then
  should_configure=1
fi
if [[ "${PDL_NATIVE_REBUILD:-}" == "1" ]]; then
  rm -rf "${build_dir}"
  should_configure=1
fi
if [[ "${should_configure}" -eq 1 ]]; then
  cmake -S "${project_dir}" -B "${build_dir}" -DCMAKE_TOOLCHAIN_FILE="${toolchain}" "${configure_args[@]}"
fi

cmake --build "${build_dir}" --target pdlc

mkdir -p "${bin_dir}"
cp "${build_dir}/pdlc" "${bin_dir}/pdlc2"
