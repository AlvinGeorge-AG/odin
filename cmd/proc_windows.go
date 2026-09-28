//go:build windows

package cmd

import (
	"fmt"
	"strconv"
)

func procKill(pidStr string) error {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf("invalid PID: %s (must be a number)", pidStr)
	}
	return fmt.Errorf("process kill not implemented on Windows yet (PID: %d)", pid)
}