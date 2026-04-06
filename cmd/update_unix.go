//go:build !windows

package cmd

import (
	"fmt"
	"os/exec"
)

func runUpdateInstall() error {
	fmt.Println("Updating...")
	fmt.Println("⚠️  You may be prompted for your sudo password...")

	_, err := exec.Command("sh", "-c", "curl -sSL https://raw.githubusercontent.com/AlvinGeorge-AG/odin/main/install.sh | bash").Output()
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Printf("✅ Updated successfully!\n")
	return nil
}
