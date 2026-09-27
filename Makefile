# Day-Off monorepo entry points. Target names are FROZEN once merged:
# CI and branch protection call them by name.
SHELL := /bin/bash
.DEFAULT_GOAL := help

FE := frontend
BE := backend
PL := pipeline

.PHONY: help setup hooks dev down pipeline lint test eval secrets promote sync clean

help: ## List targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-9s\033[0m %s\n",$$1,$$2}'

setup: hooks ## One-time local setup (deps + hooks + .env)
	@test -f .env || { cp .env.example .env; echo "created .env, fill in keys"; }
	npm install
	@if [ -f $(FE)/package.json ]; then (cd $(FE) && npm install); fi
	@if [ -f $(BE)/go.mod ]; then (cd $(BE) && go mod download); fi
	@if [ -f $(PL)/pyproject.toml ]; then (cd $(PL) && uv sync); fi
	@echo "setup done. run: make dev"

hooks: ## Install git hooks from .githooks/
	@git config core.hooksPath .githooks
	@chmod +x .githooks/* scripts/*.sh 2>/dev/null || true
	@echo "hooks installed"

dev: ## Run the whole stack: local Supabase + backend + frontend
	npx supabase start
	docker compose up

down: ## Stop the stack
	docker compose down
	npx supabase stop

pipeline: ## Run the nightly data pipeline once (batch job)
	docker compose --profile batch run pipeline

lint: ## Lint every part that exists
	@if [ -f $(FE)/package.json ]; then echo "> frontend"; (cd $(FE) && npm run lint && npx next typegen && npx tsc --noEmit); else echo "skip frontend"; fi
	@if [ -n "$$(find $(BE) -name '*.go' -print -quit 2>/dev/null)" ]; then echo "> backend"; (cd $(BE) && out=$$(gofmt -l .) && { [ -z "$$out" ] || { echo "gofmt needed:"; echo "$$out"; exit 1; }; } && go vet ./...); else echo "skip backend (no .go files yet)"; fi
	@if [ -f $(PL)/pyproject.toml ]; then echo "> pipeline"; (cd $(PL) && uv run ruff check . && uv run ruff format --check .); else echo "skip pipeline"; fi

test: ## Test every part that exists
	@if [ -f $(FE)/package.json ]; then echo "> frontend"; (cd $(FE) && npm run test --if-present); else echo "skip frontend"; fi
	@if [ -n "$$(find $(BE) -name '*.go' -print -quit 2>/dev/null)" ]; then echo "> backend"; (cd $(BE) && go test ./...); else echo "skip backend (no .go files yet)"; fi
	@if [ -f $(PL)/pyproject.toml ]; then echo "> pipeline"; (cd $(PL) && { uv run pytest; rc=$$?; [ $$rc -eq 0 ] || [ $$rc -eq 5 ]; }); else echo "skip pipeline"; fi

eval: ## nDCG@5 gate (stub until eval harness exists)
	@echo "eval-gate: stub, passes until eval/ harness lands"

secrets: ## Scan all branches for leaked keys (needs docker)
	docker run --rm -v "$$PWD:/repo" ghcr.io/gitleaks/gitleaks:latest git /repo --log-opts="--all" --no-banner

promote: ## Open release PR dev -> main (merge commit)
	gh pr create --base main --head dev --title "chore(release): promote dev to main" --template promote.md

sync: ## After a hotfix: open PR main -> dev via temp branch
	@git fetch origin && b="chore/LAP-0-sync-main-$$(date +%Y%m%d)" && \
	git push origin "origin/main:refs/heads/$$b" && \
	gh pr create --base dev --head "$$b" --title "chore(release): sync main into dev" --body "Back-merge after hotfix."

clean: ## Remove build artifacts
	rm -rf node_modules $(FE)/.next $(FE)/node_modules $(BE)/bin $(PL)/.venv
