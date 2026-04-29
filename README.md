# agent-sandbox

Throwaway repo for agent integration tests and experiments. Real PRs land here when
opt-in test/integration runs are exercised against live Managed Agents.

Nothing in this repo is production. Branches and PRs may be force-pushed,
deleted, or closed without notice.

## Layout

Each backend / orchestrator under test gets its own top-level folder so
fixtures don't collide:

- `orchestra-ma/` — fixtures the orchestra Managed Agents path drives
  agents against. Holds a minimal Go module + the `docs/test-fixtures/`
  ambiguous spec used by the block / unblock acceptance smoke.
