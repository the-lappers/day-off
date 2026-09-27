# Playbook: add a database migration

Every schema change is a versioned file. Three databases run this project (local, hosted
dev, hosted prod) and a migration file is the only thing that reaches all three.

The tool is the Supabase CLI. `npm install` at the repo root installs it (`supabase` in
`package.json`); run it through `npx`. Migrations live in `supabase/migrations/` and apply
in filename order.

## Steps

1. Create the migration:
   ```
   npx supabase migration new <short_snake_case_name>
   ```
   The CLI prefixes a UTC timestamp, e.g. `20260921155203_baseline.sql`. Keep that name.
2. Write the SQL in the generated file under `supabase/migrations/`.
3. Apply locally and check it:
   ```
   npx supabase start        # skip if make dev is already running
   npx supabase db reset
   npx supabase migration list --local
   ```
   `db reset` replays every migration from scratch, then runs `supabase/seed.sql`, which
   is how CI and a new teammate will experience it. `migration list` shows what applied.
4. Open the PR. List the migration in the PR body.
5. Merging to `dev` applies it to the hosted dev database. A promotion to `main` applies
   it to production.

## The baseline

The first migration, `20260921155203_baseline.sql`, only enables the extensions the schema
builds on (`vector` for pgvector). Tables come in their own migrations after it.

## Undoing a change

There are no down migrations. Before merging, delete your migration file and run
`npx supabase db reset`. After merging, write a new migration that reverses it.

## Rules

- **Never edit tables in the Supabase dashboard.** Local or hosted, no exceptions. A hand
  edit is invisible to everyone else and disappears on the next reset.
- **Never edit a migration that has already merged.** Write a new one.
- Migrations run before the app deploys, so a migration must work against the *old* code
  as well: add columns before writing to them, drop them a release later.
- Prod migrations happen at promotion time. Take a manual `pg_dump` first, since the free
  plan has no automatic backups.
- Data fixes are migrations too. Seed or test data goes in `supabase/seed.sql`, never in a
  migration, and uses fake text only.
