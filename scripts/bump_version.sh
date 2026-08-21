#!/usr/bin/env bash
# Set every manifest in .version-files.json to the same version.
# Usage: scripts/bump_version.sh 0.2.0

set -euo pipefail

new="${1:-}"
[[ $# -eq 1 && "$new" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] ||
    { echo "usage: ${0##*/} <X.Y.Z>" >&2; exit 2; }

command -v jq >/dev/null || { echo "jq is required" >&2; exit 1; }

cd "$(git rev-parse --show-toplevel)"
mapfile -t files < <(jq -r '.files[]' .version-files.json)

for file in "${files[@]}"; do
    [ -f "$file" ] || { echo "missing: $file" >&2; exit 1; }
done

for file in "${files[@]}"; do
    jq --arg v "$new" '.version = $v' "$file" >"$file.tmp"
    mv "$file.tmp" "$file"
    echo "$file -> $new"
done
