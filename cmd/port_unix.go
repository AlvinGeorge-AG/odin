//go:build !windows

package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func portList() error {
	out, err := exec.Command("sh", "-c", "lsof -i -P -n").Output()
	if err != nil {
		return fmt.Errorf("Failed to Run odin port ls : %w", err)
	}
	printHeader("Open Port'S")
	fmt.Println(string(out))
	return nil
}

func portIP() error {
	out, err := exec.Command("sh", "-c", "ip -4 addr show scope global | grep inet").Output()
	out1, err1 := exec.Command("sh", "-c", "curl -s https://api.ipify.org").Output()
	if err != nil {
		return fmt.Errorf("Failed to Run odin ip : %w", err)
	}
	if err1 != nil {
		return fmt.Errorf("Failed to Run odin ip : %w", err1)
	}

	lines := strings.Split(string(out), "\n")
	printHeader("🌐 IP Address")
	fmt.Printf("\nPrivate Interfaces:\n")
	n := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		n++
		fields := strings.Fields(line)
		ip := strings.Split(fields[1], "/")[0]
		iface := fields[len(fields)-1]

		fmt.Printf("%d  %-10s → %s\n", n, iface, ip)
	}
	fmt.Printf("\nPublic Interfaces:\n")
	fmt.Printf("%s\n", string(out1))

	return nil
}
