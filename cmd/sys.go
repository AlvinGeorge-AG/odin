package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func printHeader(data string) {
	upperData := strings.ToUpper(data)
	fmt.Printf("─────────────────────────────\nODIN · %s INFO\n─────────────────────────────\n", upperData)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "display information about the CPU architecture",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysInfo()
	},
}

var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "CPU/GPU temperatures",
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysTemp()
	},
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Show real-time CPU usage per core",
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysCPU()
	},
}

var memCmd = &cobra.Command{
	Use:   "ram",
	Short: "Show real-time memory usage",
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysRAM()
	},
}

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Disk usage, no tmpfs noise",
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysDisk()
	},
}

var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "show startup services and boot time",
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysBoot()
	},
}

var freeCmd = &cobra.Command{
	Use:   "free",
	Short: "Show disk space usage (df -h)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return sysFree()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(tempCmd)
	rootCmd.AddCommand(cpuCmd)
	rootCmd.AddCommand(memCmd)
	rootCmd.AddCommand(diskCmd)
	rootCmd.AddCommand(bootCmd)
	rootCmd.AddCommand(freeCmd)
}