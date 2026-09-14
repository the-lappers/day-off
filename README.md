# Day-Off

A one-day trip planner for Bangkok neighbourhoods (e.g. Talat Noi, Song Wat). Describe a "vibe" in free text (Thai or English) and the app scores venues against it, either suggesting one stop at a time or building a full-day route that respects opening hours and walking time. Plans can be re-solved live if something changes (a place closes, you're running late, etc.).

## Stack

| Layer | Choice | Notes |
|---|---|---|
| Frontend | Next.js (React) | port `3333` |
| Backend API | Go | port `8000`, health check at `/healthz` |
| Data pipeline | Python | nightly batch job (Thai sentiment + visual tagging); writes to Postgres, never called directly by the Go API |
| Database | PostgreSQL + pgvector | via [Supabase](https://supabase.com) — hosted in production, run locally through the Supabase CLI in dev |

Schema changes live as versioned migrations in `supabase/migrations/` — never applied by hand against a running database.

## Prerequisites

You need four things installed, regardless of OS: **Git**, **Node.js**, **Docker**, and **make**.

### Windows

- **Git** — [git-scm.com](https://git-scm.com/download/win), or `winget install Git.Git`
- **Node.js** (LTS, 20+) — [nodejs.org](https://nodejs.org), or `winget install OpenJS.NodeJS.LTS`
- **Docker Desktop** — [docker.com](https://www.docker.com/products/docker-desktop/). Requires WSL2 — the installer will prompt you to enable it if it isn't already.
- **make** — not included on Windows by default. Install via [Scoop](https://scoop.sh) (`scoop install make`) or [Chocolatey](https://chocolatey.org) (`choco install make`).

### macOS

- **Git** — included with Xcode Command Line Tools: `xcode-select --install`
- **Node.js** (LTS, 20+) — `brew install node`, or [nodejs.org](https://nodejs.org)
- **Docker Desktop** — [docker.com](https://www.docker.com/products/docker-desktop/), or `brew install --cask docker`
- **make** — included with Xcode Command Line Tools (same `xcode-select --install` as Git)

### Linux

- **Git** — usually preinstalled; otherwise `sudo apt install git` (Debian/Ubuntu) or `sudo dnf install git` (Fedora)
- **Node.js** (LTS, 20+) — via [nodesource](https://github.com/nodesource/distributions) or your distro's package manager
- **Docker Engine + Compose plugin** — [docs.docker.com/engine/install](https://docs.docker.com/engine/install/) (no "Desktop" needed on Linux, but you specifically need the `docker compose` plugin, not the legacy standalone `docker-compose` v1 binary)
- **make** — `sudo apt install build-essential` (Debian/Ubuntu) or `sudo dnf groupinstall "Development Tools"` (Fedora)

Verify everything is on your PATH before continuing:

```
git --version
node --version
docker --version
docker compose version
make --version
```

## Setup

```
git clone <repo-url>
cd day-off
npm install          # installs the Supabase CLI as a local devDependency
cp .env.example .env # fill in any real values if the defaults don't apply to you
```

## Running the stack

```
make dev
```

This runs `npx supabase start` (spins up a local Postgres + pgvector + Studio in Docker) followed by `docker compose up` (builds and starts the backend and frontend). It is the single command referenced by NFR-O4 — no other manual setup steps should be required.

Once it's up:
- Frontend: [http://localhost:3333](http://localhost:3333)
- Backend health check: [http://localhost:8000/healthz](http://localhost:8000/healthz)
- Supabase Studio (local DB admin UI): [http://127.0.0.1:54323](http://127.0.0.1:54323)

The nightly data pipeline is not part of `make dev` (it's a batch job, not a long-running service). Run it on demand with:

```
docker compose --profile batch run pipeline
```



## Cross-platform notes

- `docker-compose.yml` uses `host.docker.internal` so containers can reach the Postgres that `supabase start` runs on your host machine. This resolves automatically on Docker Desktop (Windows/Mac); on native Linux Docker Engine it needs an explicit `extra_hosts: host.docker.internal:host-gateway` entry, which is already included in the compose file for the services that need it.
- This setup has been validated on Windows (`docker compose config` resolves cleanly, `make`/`docker`/`node` all present and working). It has **not** yet been verified end-to-end on macOS or Linux — if you're setting up on one of those, please report back whether `make dev` works so this note can be updated.
