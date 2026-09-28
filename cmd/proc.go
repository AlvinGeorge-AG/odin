package cmd

import (
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill <pid>",
	Short: "Force kill a process by PID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return procKill(args[0])
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}