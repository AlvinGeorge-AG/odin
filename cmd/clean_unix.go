//go:build !windows

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var aptCmd = &cobra.Command{
	Use:   "apt",
	Short: "Remove unused packages and clean apt cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getuid() != 0 {
			fmt.Println("❌ This command requires sudo. Run: sudo odin clean apt")
			os.Exit(1)
		}
		out, err := exec.Command("sh", "-c", "apt autoremove && apt clean").Output()
		if err != nil {
			return fmt.Errorf("Failed to Run Odin clean apt! : %w", err)
		}
		printHeader("Clean")
		fmt.Println(string(out))
		return nil
	},
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Clears cached files under ~/.cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := exec.Command("sh", "-c", "rm -rf ~/.cache/thumbnails/* && rm -rf ~/.cache/*").Output()
		if err != nil {
			return fmt.Errorf("Failed to Run Odin clean cache! : %w", err)
		}
		printHeader("Cache Clean")
		fmt.Println(string(out))
		return nil
	},
}

func init() {
	cleanCmd.AddCommand(aptCmd)
	cleanCmd.AddCommand(cacheCmd)
}
