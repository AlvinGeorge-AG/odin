# Odin v4.0.0 — Windows Support (without breaking Linux)

## Summary
- Added **Windows builds** for existing commands while **preserving Linux/Unix behavior**.
- Implemented OS-specific logic using **Go build tags** so each binary contains only the correct implementation.

## What’s new (Windows)
- **`odin sys`**
  - Uses PowerShell/CIM (`Get-CimInstance`, `Get-Volume`, `Get-Counter`) for system, disk, CPU, RAM, boot info.
  - `odin sys temp` prints a helpful message (Windows has no universal built-in temp sensor API).
- **`odin open ports`**
  - Uses `Get-NetTCPConnection` / `Get-NetUDPEndpoint` to show externally exposed listeners.
- **`odin firewall status`**
  - Uses `Get-NetFirewallProfile`.
- **`odin port ls`**
  - Uses `Get-NetTCPConnection -State Listen`.
- **`odin port ip`**
  - Uses `Get-NetIPAddress` + public IP via `Invoke-RestMethod` (best-effort).
- **`odin clean`**
  - Adds `clean temp` and `clean cache` for common Windows locations.
  - `clean apt` is Linux-only (prints guidance for `winget`/Chocolatey).

## Linux/Unix behavior
- Existing commands still use the original tools (`lscpu`, `free`, `df`, `ufw`, `ss`, `ip`, `apt`, etc.).
- No Linux commands were removed—only reorganized behind helpers.

## Build instructions
- Linux:

```bash
go build .
```

- Windows (from Linux/macOS cross-compile):

```bash
GOOS=windows GOARCH=amd64 go build -o odin.exe .
```

## Implementation detail (how OS selection works)
- OS selection is **compile-time** via build tags:
  - `//go:build windows` → compiled only for Windows
  - `//go:build !windows` → compiled for everything except Windows

