package cmd

import (
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean system caches, logs and temp files",
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
