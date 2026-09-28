package cmd

import (
	"github.com/spf13/cobra"
)

var openPortsCmd = &cobra.Command{
	Use:   "open-ports",
	Short: "Show listening sockets on all interfaces (externally exposed)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return securityOpenPorts()
	},
}

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Show current firewall status and rules (requires sudo)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return securityFirewallStatus()
	},
}

func init() {
	rootCmd.AddCommand(openPortsCmd)
	rootCmd.AddCommand(firewallCmd)
}