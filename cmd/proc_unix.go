//go:build !windows

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func procKill(pidStr string) error {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf("invalid PID: %s (must be a number)", pidStr)
	}

	// Check if process exists
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process with PID %d not found", pid)
	}

	// Try to signal with SIGKILL (-9) first without sudo
	err = process.Kill()
	if err == nil {
		fmt.Printf("✅ Process %d killed successfully\n", pid)
		return nil
	}

	// If permission denied, try with sudo
	if os.IsPermission(err) {
		fmt.Printf("⚠️  Permission denied. Attempting with sudo...\n")
		cmd := exec.Command("sudo", "kill", "-9", strconv.Itoa(pid))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err == nil {
			fmt.Printf("✅ Process %d killed successfully with sudo\n", pid)
			return nil
		}
		return fmt.Errorf("failed to kill process %d even with sudo: %w", pid, err)
	}

	return fmt.Errorf("failed to kill process %d: %w", pid, err)
}