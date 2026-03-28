#!/bin/bash

set -e

# ─────────────────────────────────────────────
#  Odin Installer
#  https://github.com/AlvinGeorge-AG/odin
# ─────────────────────────────────────────────

REPO="AlvinGeorge-AG/odin"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="odin"

# ── Colors ────────────────────────────────────
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
RESET='\033[0m'

# ── Helpers ───────────────────────────────────
info()    { echo -e "${BLUE}[~]${RESET} $1"; }
success() { echo -e "${GREEN}[✓]${RESET} $1"; }
warning() { echo -e "${YELLOW}[!]${RESET} $1"; }
error()   { echo -e "${RED}[✗]${RESET} $1"; exit 1; }

# ── Banner ────────────────────────────────────
echo ""
echo -e "${BOLD}⚡ Odin Installer${RESET}"
echo -e "   Developer CLI Toolkit for Linux"
echo -e "   https://github.com/${REPO}"
echo ""

# ── Check OS ──────────────────────────────────
if [[ "$OSTYPE" != "linux-gnu"* ]]; then
    error "Odin only supports Linux. Detected: $OSTYPE"
fi

# ── Detect Architecture ───────────────────────
ARCH=$(uname -m)
case $ARCH in
    x86_64)  ARCH_LABEL="amd64" ;;
    aarch64) ARCH_LABEL="arm64" ;;
    armv7l)  ARCH_LABEL="arm64" ;;
    *)       error "Unsupported architecture: $ARCH" ;;
esac

info "Detected architecture: $ARCH ($ARCH_LABEL)"

# ── Check sudo ────────────────────────────────
if ! command -v sudo &>/dev/null; then
    error "sudo is required but not found."
fi

# ── Check Dependencies ────────────────────────
echo ""
info "Checking dependencies..."

MISSING=()

command -v lsof    &>/dev/null || MISSING+=("lsof")
command -v curl    &>/dev/null || MISSING+=("curl")
command -v ufw     &>/dev/null || MISSING+=("ufw")
command -v sensors &>/dev/null || MISSING+=("lm-sensors")

if [ ${#MISSING[@]} -ne 0 ]; then
    warning "Missing dependencies: ${MISSING[*]}"
    info "Installing missing packages..."
    sudo apt-get update -qq
    sudo apt-get install -y "${MISSING[@]}" &>/dev/null
    success "Dependencies installed."
else
    success "All dependencies are present."
fi

# ── Get Latest Release ────────────────────────
echo ""
info "Fetching latest release..."

LATEST=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep '"tag_name"' \
    | cut -d '"' -f 4)

if [ -z "$LATEST" ]; then
    error "Could not fetch latest release. Check your internet connection."
fi

success "Latest version: $LATEST"

# ── Download Binary ───────────────────────────
BINARY_URL="https://github.com/${REPO}/releases/download/${LATEST}/odin-linux-${ARCH_LABEL}"
TMP_PATH="/tmp/odin"

info "Downloading Odin ${LATEST} for linux-${ARCH_LABEL}..."

curl -sL "$BINARY_URL" -o "$TMP_PATH"

if [ ! -f "$TMP_PATH" ]; then
    error "Download failed. Binary not found at $BINARY_URL"
fi

# ── Install Binary ────────────────────────────
sudo mv "$TMP_PATH" "$INSTALL_DIR/$BINARY_NAME"
sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"

# ── Verify Installation ───────────────────────
if command -v odin &>/dev/null; then
    echo ""
    echo -e "${GREEN}${BOLD}✅ Odin installed successfully!${RESET}"
    echo ""
    echo -e "   Run ${BOLD}odin --help${RESET} to get started."
    echo ""
    echo -e "   Example Run ${BOLD}odin ip"
    echo ""
else
    error "Installation failed. Binary not found in PATH."
fi