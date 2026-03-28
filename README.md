# ⚡ Odin
 
> **Developer CLI Toolkit for Linux**
 
A developer-focused Linux CLI toolkit that abstracts painful, hard-to-remember commands into simple, intuitive ones. Instead of hunting through man pages, just use Odin.
 
```bash
# Instead of this:
lsof -ti:3000 | xargs kill -9
 
# Just do this:
odin port kill 3000
```
 
---
 
## Install
 
```bash
curl -sSL https://raw.githubusercontent.com/AlvinGeorge-AG/odin/main/install.sh | bash
```
 
This will:
- Check and install any missing dependencies automatically
- Download the latest Odin binary for your architecture

## Usage

Run `odin` or `odin --help` for the full command tree.

### System (`odin sys`)

| Command | Description |
|--------|-------------|
| `odin sys info` | CPU (`lscpu`), memory (`free -h`), disk (`df -h`), kernel (`uname -a`). |
| `odin sys temp` | Temperatures via `sensors`. |
| `odin sys cpu` | Snapshot CPU line from `top`. |
| `odin sys ram` | Memory summary and top processes by `%mem`. |
| `odin sys disk` | Filesystem usage (`df -h`). |

### Network (`odin port`, `odin ip`)

| Command | Description |
|--------|-------------|
| `odin port ls` | Open sockets / ports (`lsof -i -P -n`). |
| `odin ip` | Private IPv4 addresses (global scope) and public IP via ipify. |

### Cleanup (`odin clean`)

| Command | Description |
|--------|-------------|
| `odin clean apt` | `apt autoremove` and `apt clean`. **Requires sudo.** |
| `odin clean cache` | Clears thumbnail and general files under `~/.cache`. |

### Security / exposure (`odin open`, `odin firewall`)

| Command | Description |
|--------|-------------|
| `odin open ports` | Listening sockets on all interfaces (`ss`, filtered). |
| `odin firewall status` | `ufw status`. **Requires sudo.** |

## Root privileges

Run with `sudo` when the tool tells you to:

- `sudo odin clean apt`
- `sudo odin firewall status`

---
 
## Architecture
 
```
odin/
├── main.go                 # Entry point
├── go.mod
├── install.sh              # One-line install script
└── cmd/
    ├── root.go             # Base odin command
    ├── port.go             # Port + IP commands
    ├── sys.go              # System health commands
    ├── proc.go             # Process management
    ├── clean.go            # Cleanup commands
    └── security.go         # Firewall + exposed ports
```
 
---

---
 
## Build from Source
 
Requirements: Go 1.21+
 
```bash
git clone https://github.com/AlvinGeorge-AG/odin.git
cd odin
go mod tidy
go build -o odin .
sudo mv odin /usr/local/bin/
```
 
---

---
 
## Contributing
 
Pull requests are welcome. If you find a painful Linux command worth abstracting, open an issue.
 
```bash
git clone https://github.com/AlvinGeorge-AG/odin.git
cd odin
# Add your command in cmd/ following the existing pattern
# Submit a PR
```
 
---

---
 
## License
 
MIT License — free to use, modify and distribute.
 
---