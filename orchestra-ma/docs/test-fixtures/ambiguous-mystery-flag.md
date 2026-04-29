# Mystery flag

## Goal

Add a `--debug` boolean flag to the `orchestra-ma` binary that enables
debug-level logging.

## Spec

- Flag name: `--debug`
- Type: `bool`
- Default: `false`
- Surface: top-level CLI flag on the `orchestra-ma` binary (parsed via
  the standard library `flag` package).
- Behavior:
  - When unset (default): logger emits at `slog.LevelInfo` and above.
  - When set: logger emits at `slog.LevelDebug` and above, and a
    `"debug logging enabled"` debug record is written at startup.
- Logger: `log/slog` text handler writing to `os.Stderr`.

## Verification

- `go vet ./...` passes.
- `go build ./...` passes.
- `orchestra-ma` (no flag) prints the existing info log and no debug
  record.
- `orchestra-ma --debug` additionally prints the `"debug logging
  enabled"` record at `level=DEBUG`.
