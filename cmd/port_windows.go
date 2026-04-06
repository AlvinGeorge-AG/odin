//go:build windows

package cmd

import (
	"fmt"
	"strings"
)

func portList() error {
	out, err := runPowerShell(`Get-NetTCPConnection -State Listen | Select-Object LocalAddress,LocalPort,OwningProcess | Sort-Object LocalPort | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("Failed to Run odin port ls : %w\n%s", err, string(out))
	}
	printHeader("Open Port'S")
	fmt.Println(string(out))
	return nil
}

func portIP() error {
	outPriv, err := runPowerShell(`Get-NetIPAddress -AddressFamily IPv4 | Where-Object { $_.IPAddress -notlike '169.254*' -and $_.IPAddress -ne '127.0.0.1' } | Select-Object InterfaceAlias,IPAddress | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("Failed to Run odin ip : %w\n%s", err, string(outPriv))
	}
	outPub, err2 := runPowerShell(`try { (Invoke-RestMethod -UseBasicParsing -Uri "https://api.ipify.org").ToString() } catch { "" }`)
	if err2 != nil {
		return fmt.Errorf("Failed to Run odin ip : %w\n%s", err2, string(outPub))
	}

	printHeader("🌐 IP Address")
	fmt.Printf("\nPrivate Interfaces:\n")
	fmt.Println(strings.TrimSpace(string(outPriv)))

	fmt.Printf("\nPublic Interfaces:\n")
	fmt.Println(strings.TrimSpace(string(outPub)))
	return nil
}
