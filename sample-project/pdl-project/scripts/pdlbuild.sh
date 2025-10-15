#!/usr/bin/env sh
set -euo pipefail

SCRIPT_DIR="$(dirname "$0")"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
DEFAULT_BIN_DIR="$(cd "$PROJECT_ROOT/../../bin" 2>/dev/null && pwd)"
CLI_PATH=""
if [ -n "${PDL_BIN_PATH:-}" ] && [ -x "${PDL_BIN_PATH%/}/pdl" ]; then
  CLI_PATH="${PDL_BIN_PATH%/}/pdl"
elif [ -n "$DEFAULT_BIN_DIR" ] && [ -x "$DEFAULT_BIN_DIR/pdl" ]; then
  CLI_PATH="$DEFAULT_BIN_DIR/pdl"
else
  CLI_PATH="$(command -v pdl || true)"
fi

if [ -z "$CLI_PATH" ] || [ ! -x "$CLI_PATH" ]; then
  echo "pdl binary not found; install the CLI or set PDL_BIN_PATH" >&2
  exit 1
fi

"$CLI_PATH" --build --config "$PROJECT_ROOT/pdl.config.json" "$@"

OUTPUT_GO_DIR="$PROJECT_ROOT/output/db2pdl/go"
if [ -d "$OUTPUT_GO_DIR" ]; then
  cat > "$OUTPUT_GO_DIR/go.mod" <<'MOD'
module github.com/kapablanka/pdl/sample/generated

go 1.21
MOD
fi
