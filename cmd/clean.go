package cmd

import (
	"github.com/spf13/cobra"
)

var cleanAptCmd = &cobra.Command{
	Use:   "clean-apt",
	Short: "Remove unused packages and clean apt cache (requires sudo)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cleanApt()
	},
}

var cleanCacheCmd = &cobra.Command{
	Use:   "clean-cache",
	Short: "Clear cached files under ~/.cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cleanCache()
	},
}

func init() {
	rootCmd.AddCommand(cleanAptCmd)
	rootCmd.AddCommand(cleanCacheCmd)
}