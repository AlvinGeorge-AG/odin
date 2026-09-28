//go:build !windows

package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func spaceUsage(path string) error {
	// Use du with human-readable output, max depth 1, then sort by size (human-readable)
	out, err := exec.Command("sh", "-c", fmt.Sprintf("du -h --max-depth=1 %s 2>/dev/null | sort -hr", path)).Output()
	if err != nil {
		// Try without sort -hr if not supported (some systems)
		out, err = exec.Command("sh", "-c", fmt.Sprintf("du -h --max-depth=1 %s 2>/dev/null", path)).Output()
		if err != nil {
			return fmt.Errorf("failed to get disk usage for %s: %w", path, err)
		}
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Printf("No files or directories found in %s\n", path)
		return nil
	}

	printHeader(fmt.Sprintf("💾 Disk Usage: %s", path))
	fmt.Println()

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			size := parts[0]
			name := strings.Join(parts[1:], " ")
			fmt.Printf("%-10s %s\n", size, name)
		} else if len(parts) == 1 {
			fmt.Printf("%-10s %s\n", parts[0], ".")
		}
	}

	return nil
}