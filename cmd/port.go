package cmd

import (
	"github.com/spf13/cobra"
)

var portCMD = &cobra.Command{
	Use:   "port",
	Short: "Manage and inspect network ports",
}

var lsCmd = &cobra.Command{
	Use:   "ls",
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
	portCMD.AddCommand(lsCmd)
	rootCmd.AddCommand(ipCmd)
	rootCmd.AddCommand(portCMD)
}
