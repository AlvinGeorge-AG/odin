package cmd

import (
	"github.com/spf13/cobra"
)

var spaceCmd = &cobra.Command{
	Use:   "space [path]",
	Short: "Show disk usage for a path (sorted by size, largest first)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		return spaceUsage(path)
	},
}

func init() {
	rootCmd.AddCommand(spaceCmd)
}