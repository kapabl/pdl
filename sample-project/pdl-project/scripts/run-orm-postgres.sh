#!/usr/bin/env bash
set -euo pipefail

script_path="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
project_dir="$(cd "${script_path}/.." && pwd)"
config_file="${project_dir}/pdl.config.json"
environment_file="${project_dir}/.env.pg.local"

if [ ! -f "${environment_file}" ]; then
  echo "Environment file not found at ${environment_file}" >&2
  exit 1
fi

if [ ! -f "${config_file}" ]; then
  echo "Configuration file not found at ${config_file}" >&2
  exit 1
fi

set -a
source "${environment_file}"
set +a

export PDL_OUTPUT="${project_dir}/output"
export PDL_DB2PDL_OUTPUT="${project_dir}/output/db2pdl"
export PDL_GEN_OUTPUT_PHP="${project_dir}/output/php"
export PDL_GEN_OUTPUT_JS="${project_dir}/output/js"
export PDL_GEN_OUTPUT_BUNDLE="${project_dir}/output/bundle"
export PDL_GEN_OUTPUT_GO="${project_dir}/output/go"
DB2PDL_BIN=""
if [ -n "${PDL_BIN_PATH:-}" ]; then
  DB2PDL_BIN="${PDL_BIN_PATH%/}/db2pdl"
fi
if [ ! -x "${DB2PDL_BIN}" ]; then
  DB2PDL_BIN=$(command -v db2pdl || true)
fi
if [ -z "${DB2PDL_BIN}" ] || [ ! -x "${DB2PDL_BIN}" ]; then
  echo "db2pdl binary not found; install PDL tooling or set PDL_BIN_PATH" >&2
  exit 1
fi
"${DB2PDL_BIN}" --run --config "${config_file}" --exit

mkdir -p "${project_dir}/output/db2pdl/go"
cat > "${project_dir}/output/db2pdl/go/go.mod" <<'MOD'
module github.com/kapablanka/pdl/sample

go 1.21

require github.com/kapablanka/pdl/pdl/infra/go v0.0.0

replace github.com/kapablanka/pdl/pdl/infra/go => ../../../../pdl/infra/go
MOD
