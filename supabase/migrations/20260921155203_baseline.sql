-- Baseline (LAP-21): the extensions the schema builds on. Tables arrive in their own
-- migrations after this one. Never edit a merged migration; write a new one
-- (docs/playbooks/add-a-migration.md).

-- pgvector: venue and query embeddings (docs/ARCHITECTURE.md).
create extension if not exists vector with schema extensions;
