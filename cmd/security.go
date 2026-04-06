package cmd

import (
	"github.com/spf13/cobra"
)

var permOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Show ports that are externally exposed",
}

var portCmd = &cobra.Command{
	Use:   "ports",
	Short: "To show only externally exposed ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		return securityOpenPorts()
	},
}

var fireWall = &cobra.Command{
	Use:   "firewall",
	Short: "Show current firewall status and rules , Show ports that are externally exposed",
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "To show only externally exposed ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		return securityFirewallStatus()
	},
}

func init() {
	permOpenCmd.AddCommand(portCmd)
	rootCmd.AddCommand(permOpenCmd)

	fireWall.AddCommand(statusCmd)
	rootCmd.AddCommand(fireWall)
}
