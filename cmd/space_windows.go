//go:build windows

package cmd

import (
	"fmt"
)

func spaceUsage(path string) error {
	return fmt.Errorf("disk usage command not implemented on Windows yet (path: %s)", path)
}