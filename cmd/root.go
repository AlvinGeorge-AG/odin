package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "odin",
	Short: "Developer-focused Linux CLI toolkit",
	Long: `Odin abstracts painful, easy-to-forget Linux workflows into simple subcommands.Commands are grouped by category (port, proc, sys, clean, perm).`,
	SilenceUsage: true,
}

// Execute runs the root command tree.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
