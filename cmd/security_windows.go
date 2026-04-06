//go:build windows

package cmd

import (
	"fmt"
)

func securityOpenPorts() error {
	out, err := runPowerShell(`$tcp = Get-NetTCPConnection -State Listen | Where-Object { $_.LocalAddress -in @('0.0.0.0','::') } | Select-Object LocalAddress,LocalPort,OwningProcess; $udp = Get-NetUDPEndpoint | Where-Object { $_.LocalAddress -in @('0.0.0.0','::') } | Select-Object LocalAddress,LocalPort,OwningProcess; "TCP (LISTEN)"; $tcp | Sort-Object LocalPort | Format-Table -AutoSize | Out-String; "UDP"; $udp | Sort-Object LocalPort | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("Failed to Run odin open ports : %w\n%s", err, string(out))
	}
	printHeader("Open Ports")
	fmt.Println(string(out))
	return nil
}

func securityFirewallStatus() error {
	out, err := runPowerShell(`Get-NetFirewallProfile | Select-Object Name, Enabled, DefaultInboundAction, DefaultOutboundAction | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("Failed to Run odin firewall status : %w\n%s", err, string(out))
	}
	printHeader("Firewall Status")
	fmt.Println(string(out))
	return nil
}
