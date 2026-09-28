# ⚡ Odin
  
> **Developer CLI Toolkit for Linux**
  
A developer-focused Linux CLI toolkit that abstracts painful, hard-to-remember commands into simple, intuitive ones. Instead of hunting through man pages, just use Odin.
  
```bash
# Instead of this:
lsof -ti:3000 | xargs kill -9
  
# Just do this:
odin kill 3000
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
  
Run `odin` or `odin --help` for the full command tree. All commands are now **2 words or less** for easy memorization.
 
### Process Management
 
| Command | Description |
|--------|-------------|
| `odin kill <pid>` | Force kill a process by PID (prompts for sudo if needed) |
| `odin ps` | Show top processes by CPU and memory (coming soon) |
 
### System Info
 
| Command | Description |
|--------|-------------|
| `odin info` | CPU (`lscpu`), memory (`free -h`), disk (`df -h`), kernel (`uname -a`) |
| `odin temp` | Temperatures via `sensors` |
| `odin cpu` | Snapshot CPU line from `top` |
| `odin ram` | Memory summary and top processes by `%mem` |
| `odin disk` | Filesystem usage (`df -h`) |
| `odin boot` | Boot time, uptime, and top boot services |
| `odin free` | Disk free space (`df -h`) |
 
### Disk Usage
 
| Command | Description |
|--------|-------------|
| `odin space [path]` | Disk usage for path (default: current dir), sorted by size |
 
### Network
 
| Command | Description |
|--------|-------------|
| `odin ports` | Open sockets / ports (`lsof -i -P -n`) |
| `odin ip` | Private IPv4 addresses (global scope) and public IP via ipify |
| `odin open-ports` | Listening sockets on all interfaces (`ss`, filtered) |
 
### Security
 
| Command | Description |
|--------|-------------|
| `odin firewall` | `ufw status` **Requires sudo.** |
 
### Cleanup
 
| Command | Description |
|--------|-------------|
| `odin clean-apt` | `apt autoremove` and `apt clean`. **Requires sudo.** |
| `odin clean-cache` | Clears thumbnail and general files under `~/.cache`. |
 
### Environment
 
| Command | Description |
|--------|-------------|
| `odin env` | Show all environment variables |
| `odin env-find <term>` | Find environment variable by name |
 
## Root privileges
 
Run with `sudo` when the tool tells you to:
 
- `sudo odin clean-apt`
- `sudo odin firewall`
 
---
  
## Architecture
  
```
odin/
├── main.go                 # Entry point
├── go.mod
├── install.sh              # One-line install script
└── cmd/
    ├── root.go             # Base odin command
    ├── port.go             # ports, ip
    ├── sys.go              # info, temp, cpu, ram, disk, boot, free
    ├── proc.go             # kill
    ├── space.go            # space
    ├── clean.go            # clean-apt, clean-cache
    ├── security.go         # open-ports, firewall
    └── env.go              # env, env-find
```
  
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
  
## Contributing
  
Pull requests are welcome. If you find a painful Linux command worth abstracting, open an issue.
  
```bash
git clone https://github.com/AlvinGeorge-AG/odin.git
cd odin
# Add your command in cmd/ following the existing pattern
# Submit a PR
```
  
---
  
## License
  
MIT License — free to use, modify and distribute.
  
---