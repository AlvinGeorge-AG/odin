//go:build windows

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tempCmdWin = &cobra.Command{
	Use:   "temp",
	Short: "Clear common Windows temp directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := runPowerShell(`$paths = @($env:TEMP, $env:TMP, "$env:WINDIR\Temp"); foreach ($p in $paths) { if (Test-Path $p) { Get-ChildItem -LiteralPath $p -Force -ErrorAction SilentlyContinue | Remove-Item -Recurse -Force -ErrorAction SilentlyContinue } }; "Cleared: " + ($paths -join ", ")`)
		if err != nil {
			return fmt.Errorf("failed to run odin clean temp: %w\n%s", err, string(out))
		}
		printHeader("Clean")
		fmt.Println(string(out))
		return nil
	},
}

var cacheCmdWin = &cobra.Command{
	Use:   "cache",
	Short: "Clear common user cache locations (best-effort)",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := runPowerShell(`$paths = @("$env:LOCALAPPDATA\Temp", "$env:LOCALAPPDATA\Microsoft\Windows\INetCache"); foreach ($p in $paths) { if (Test-Path $p) { Get-ChildItem -LiteralPath $p -Force -ErrorAction SilentlyContinue | Remove-Item -Recurse -Force -ErrorAction SilentlyContinue } }; "Cleared: " + ($paths -join ", ")`)
		if err != nil {
			return fmt.Errorf("failed to run odin clean cache: %w\n%s", err, string(out))
		}
		printHeader("Cache Clean")
		fmt.Println(string(out))
		return nil
	},
}

var aptCmdWin = &cobra.Command{
	Use:   "apt",
	Short: "Not available on Windows (use winget/choco instead)",
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader("Clean")
		fmt.Println("This subcommand is Linux/Unix-only.")
		fmt.Println("On Windows, consider: winget upgrade --all")
		fmt.Println("Or: choco upgrade all (if using Chocolatey)")
		_ = os.ErrInvalid
		return nil
	},
}

func init() {
	// Keep subcommand names consistent across OS where possible.
	cleanCmd.AddCommand(aptCmdWin)
	cleanCmd.AddCommand(cacheCmdWin)
	cleanCmd.AddCommand(tempCmdWin)
}
