package cmd

import (
	"github.com/spf13/cobra"
)

var portsCmd = &cobra.Command{
	Use:   "ports",
	Short: "Show externally exposed ports (0.0.0.0 only)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return portList()
	},
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show private and public IP addresses",
	RunE: func(cmd *cobra.Command, args []string) error {
		return portIP()
	},
}

func init() {
	rootCmd.AddCommand(portsCmd)
	rootCmd.AddCommand(ipCmd)
}