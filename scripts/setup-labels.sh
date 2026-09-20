#!/usr/bin/env bash
# Usage: scripts/setup-labels.sh [owner/repo]   (needs gh + yq-free python3)
set -euo pipefail
REPO="${1:-the-lappers/day-off}"

python3 - "$REPO" <<'PY' | while IFS=$'\t' read -r name color desc; do
import re, sys
for line in open(".github/labels.yml"):
    m = re.search(r'name: "([^"]+)",\s*color: "([^"]+)",\s*description: "([^"]*)"', line)
    if m: print("\t".join(m.groups()))
PY
  gh label create "$name" --repo "$REPO" --color "$color" --description "$desc" --force
done

# drop GitHub defaults we replaced
for old in bug enhancement documentation "good first issue" "help wanted" invalid question wontfix duplicate; do
  gh label delete "$old" --repo "$REPO" --yes 2>/dev/null || true
done
echo "labels synced"
