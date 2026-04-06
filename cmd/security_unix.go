//go:build !windows

package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func securityOpenPorts() error {
	out, err := exec.Command("sh", "-c", "ss -tuln | awk 'NR==1 || $5 ~ /^0\\.0\\.0\\.0/ || $5 ~ /^\\[::\\]/'").Output()
	if err != nil {
		return fmt.Errorf("Failed to Run odin open ports : %w", err)
	}
	printHeader("Open Ports")
	fmt.Println(string(out))
	return nil
}

func securityFirewallStatus() error {
	if os.Getuid() != 0 {
		fmt.Println("❌ This command requires sudo. Run: sudo odin firewall status")
		os.Exit(1)
	}
	out, err := exec.Command("ufw", "status").Output()
	if err != nil {
		return fmt.Errorf("Failed to Run odin firewall status : %w", err)
	}
	printHeader("Firewall Status")
	fmt.Println(string(out))
	return nil
}
