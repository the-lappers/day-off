# Playbook: add a database migration

Every schema change is a versioned file. Three databases run this project (local, hosted
dev, hosted prod) and a migration file is the only thing that reaches all three.

## Steps

1. Create the migration:
   ```
   npx supabase migration new <short_snake_case_name>
   ```
2. Write the SQL in the generated file under `supabase/migrations/`.
3. Apply locally and check it:
   ```
   npx supabase db reset
   make dev
   ```
   `db reset` replays every migration from scratch, which is how CI and a new teammate
   will experience it.
4. Open the PR. List the migration in the PR body.
5. Merging to `dev` applies it to the hosted dev database. A promotion to `main` applies
   it to production.

## Rules

- **Never edit tables in the Supabase dashboard.** Local or hosted, no exceptions. A hand
  edit is invisible to everyone else and disappears on the next reset.
- **Never edit a migration that has already merged.** Write a new one.
- Migrations run before the app deploys, so a migration must work against the *old* code
  as well: add columns before writing to them, drop them a release later.
- Prod migrations happen at promotion time. Take a manual `pg_dump` first, since the free
  plan has no automatic backups.
- Data fixes are migrations too, but seed or test data belongs in a seed script, not in a
  migration.
