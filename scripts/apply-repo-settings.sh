#!/usr/bin/env bash
# LAP-98 + LAP-99. Usage: scripts/apply-repo-settings.sh [owner/repo]
# Needs: gh auth login (admin). Safe to re-run.
set -euo pipefail
REPO="${1:-the-lappers/day-off}"

echo "> ensure dev branch exists (from main)"
if ! gh api "repos/$REPO/branches/dev" --silent 2>/dev/null; then
  sha=$(gh api "repos/$REPO/git/ref/heads/main" --jq .object.sha)
  gh api -X POST "repos/$REPO/git/refs" -f ref=refs/heads/dev -f sha="$sha" --silent
  echo "  created dev at $sha"
fi

echo "> default branch = dev, squash (dev) + merge commit (main), auto-delete"
gh api -X PATCH "repos/$REPO" --silent \
  -f default_branch=dev \
  -F allow_squash_merge=true \
  -F allow_merge_commit=true \
  -F allow_rebase_merge=false \
  -F allow_auto_merge=true \
  -F delete_branch_on_merge=true \
  -F allow_update_branch=true \
  -f squash_merge_commit_title=PR_TITLE \
  -f squash_merge_commit_message=PR_BODY \
  -f merge_commit_title=PR_TITLE \
  -f merge_commit_message=PR_BODY

echo "> secret scanning + push protection"
gh api -X PATCH "repos/$REPO" --silent \
  -f 'security_and_analysis[secret_scanning][status]=enabled' \
  -f 'security_and_analysis[secret_scanning_push_protection][status]=enabled' \
  || echo "  skipped: needs public repo. use 'make secrets' meanwhile"

apply_ruleset() { # $1 = name, $2 = file
  local id
  id=$(gh api "repos/$REPO/rulesets" --jq ".[] | select(.name==\"$1\") | .id" 2>/dev/null || true)
  if [ -n "$id" ]; then
    gh api -X PUT "repos/$REPO/rulesets/$id" --input "$2" --silent && echo "  $1 updated"
  else
    gh api -X POST "repos/$REPO/rulesets" --input "$2" --silent && echo "  $1 created" \
      || echo "  $1 FAILED: private repo on GitHub Free cannot enforce rulesets"
  fi
}
echo "> rulesets"
apply_ruleset protect-dev  .github/rulesets/dev.json
apply_ruleset protect-main .github/rulesets/main.json
