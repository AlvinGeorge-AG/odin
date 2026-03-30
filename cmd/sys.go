package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var sysCmd = &cobra.Command{
	Use:   "sys",
	Short: "Monitor system health, CPU, RAM and disk",
}

func printHeader(data string) {
	upperData := strings.ToUpper(data)
	fmt.Printf("─────────────────────────────\nODIN · %s INFO\n─────────────────────────────\n", upperData)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "display information about the CPU architecture",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader("📊 SYSTEM")
		out, err := exec.Command("lscpu").Output()
		out1, err1 := exec.Command("free", "-h").Output()
		out2, err2 := exec.Command("df", "-h").Output()
		out3, err3 := exec.Command("uname", "-a").Output()
		if err != nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w", err)
		}
		if err1 != nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w", err1)
		}
		if err2 != nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w", err2)
		}
		if err3 != nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w", err3)
		}
		printHeader("User")
		fmt.Println(string(out3))
		printHeader("System")
		fmt.Println(string(out))
		printHeader("Memory")
		fmt.Println(string(out1))
		printHeader("Disk")
		fmt.Println(string(out2))
		return nil
	},
}

var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "CPU/GPU temperatures",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := exec.Command("sensors").Output()
		if err != nil {
			return fmt.Errorf("Failed to Run odin sys temp : %w", err)
		}
		printHeader("📊 Temperature")
		fmt.Println(string(out))
		return nil
	},
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Show real-time CPU usage per core",
	RunE: func(cmd *cobra.Command, args []string) error {
		out1, err1 := exec.Command("sh", "-c", "top -bn1 | grep %Cpu").Output()
		if err1 != nil {
			return fmt.Errorf("Failed to Run odin sys cpu : %w", err1)
		}
		printHeader("📊 CPU")
		fmt.Println(string(out1))
		return nil
	},
}

var memCmd = &cobra.Command{
	Use:   "ram",
	Short: "Show real-time memory usage",
	RunE: func(cmd *cobra.Command, args []string) error {
		out1, err1 := exec.Command("sh", "-c", "top -bn1 | grep Mem").Output()
		out2, err2 := exec.Command("sh", "-c", "ps -eo user,pid,pcpu,pmem,comm --sort=-%mem | head").Output()
		if err1 != nil {
			return fmt.Errorf("Failed to Run odin sys ram : %w", err1)
		}
		if err2 != nil {
			return fmt.Errorf("Failed to Run odin sys ram : %w", err2)
		}
		printHeader("📊 RAM")
		fmt.Println(string(out1))
		fmt.Println(string(out2))
		return nil
	},
}

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Disk usage, no tmpfs noise",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := exec.Command("df", "-h").Output()
		if err != nil {
			return fmt.Errorf("Failed to Run odin sys disk: %w", err)
		}
		printHeader("📊 Disk Usage")
		fmt.Println(string(out))
		return nil
	},
}

var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "show startup services and boot time",
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader("Boot")
		out1, err1 := exec.Command("who", "-b").Output()
		if err1 != nil {
			return fmt.Errorf("Failed to Run odin sys boot (who -b): %w", err1)
		}
		out2, err2 := exec.Command("uptime", "-p").Output()
		if err2 != nil {
			return fmt.Errorf("Failed to Run odin sys boot (uptime -p): %w", err2)
		}
		out3, err3 := exec.Command("systemd-analyze").Output()
		if err3 != nil {
			return fmt.Errorf("Failed to Run odin sys boot (systemd-analyze): %w", err3)
		}
		out4, err4 := exec.Command("sh", "-c", "systemd-analyze blame | head -10").Output()
		if err4 != nil {
			return fmt.Errorf("Failed to Run odin sys boot (systemd-analyze blame): %w", err4)
		}
		parts := strings.SplitN(strings.TrimSpace(string(out1)), "  ", 2)
		if len(parts) == 2 {
			fmt.Println("Boot time:", strings.TrimSpace(parts[1]))
		} else {
			fmt.Println(string(out1))
		}
		printHeader("Uptime")
		fmt.Println(strings.TrimSpace(string(out2)))
		printHeader("Systemd analyze")
		fmt.Println(strings.TrimSpace(string(out3)))
		printHeader("Top boot services")
		fmt.Println(strings.TrimSpace(string(out4)))
		return nil
	},
}

func init() {
	sysCmd.AddCommand(infoCmd)
	sysCmd.AddCommand(bootCmd)
	sysCmd.AddCommand(tempCmd)
	sysCmd.AddCommand(cpuCmd)
	sysCmd.AddCommand(memCmd)
	sysCmd.AddCommand(diskCmd)
	rootCmd.AddCommand(sysCmd)
}
