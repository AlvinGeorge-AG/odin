//go:build windows

package cmd

import "fmt"

func runUpdateInstall() error {
	fmt.Println("Updating on Windows isn't wired to an installer script yet.")
	fmt.Println("Best option right now:")
	fmt.Println("- Download the latest Windows binary from GitHub Releases")
	fmt.Println("- Replace your existing odin.exe with the new one")
	fmt.Println("If you want, I can add an `install.ps1` flow so `odin update` can self-update on Windows.")
	return nil
}
