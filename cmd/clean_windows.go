//go:build windows

package cmd

import (
	"fmt"
	"os"
)

func cleanApt() error {
	printHeader("Clean")
	fmt.Println("This subcommand is Linux/Unix-only.")
	fmt.Println("On Windows, consider: winget upgrade --all")
	fmt.Println("Or: choco upgrade all (if using Chocolatey)")
	_ = os.ErrInvalid
	return nil
}

func cleanCache() error {
	out, err := runPowerShell(`$paths = @("$env:LOCALAPPDATA\Temp", "$env:LOCALAPPDATA\Microsoft\Windows\INetCache"); foreach ($p in $paths) { if (Test-Path $p) { Get-ChildItem -LiteralPath $p -Force -ErrorAction SilentlyContinue | Remove-Item -Recurse -Force -ErrorAction SilentlyContinue } }; "Cleared: " + ($paths -join ", ")`)
	if err != nil {
		return fmt.Errorf("failed to run odin clean-cache: %w\n%s", err, string(out))
	}
	printHeader("Cache Clean")
	fmt.Println(string(out))
	return nil
}