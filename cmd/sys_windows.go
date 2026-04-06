//go:build windows

package cmd

import (
	"fmt"
	"os/exec"
)

func runPowerShell(script string) ([]byte, error) {
	return exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script).CombinedOutput()
}

func sysInfo() error {
	printHeader("User")
	outUser, err := runPowerShell(`"$env:USERNAME@$env:COMPUTERNAME"; (Get-CimInstance Win32_OperatingSystem | Select-Object Caption, Version, OSArchitecture, CSName | Format-List | Out-String)`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys info (user/os): %w\n%s", err, string(outUser))
	}
	fmt.Println(string(outUser))

	printHeader("System")
	outSys, err := runPowerShell(`Get-CimInstance Win32_Processor | Select-Object Name, Manufacturer, NumberOfCores, NumberOfLogicalProcessors, MaxClockSpeed | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys info (cpu): %w\n%s", err, string(outSys))
	}
	fmt.Println(string(outSys))

	printHeader("Memory")
	outMem, err := runPowerShell(`$os = Get-CimInstance Win32_OperatingSystem; [pscustomobject]@{TotalGB=[math]::Round($os.TotalVisibleMemorySize/1MB,2); FreeGB=[math]::Round($os.FreePhysicalMemory/1MB,2)} | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys info (mem): %w\n%s", err, string(outMem))
	}
	fmt.Println(string(outMem))

	printHeader("Disk")
	outDisk, err := runPowerShell(`Get-Volume | Select-Object DriveLetter, FileSystemLabel, FileSystemType, SizeRemaining, Size | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys info (disk): %w\n%s", err, string(outDisk))
	}
	fmt.Println(string(outDisk))
	return nil
}

func sysTemp() error {
	printHeader("📊 Temperature")
	fmt.Println("Temperature sensors are not universally available on Windows via built-in tools.")
	fmt.Println("If you have vendor tools (e.g., HWiNFO), consider integrating them later.")
	return nil
}

func sysCPU() error {
	out, err := runPowerShell(`Get-Counter '\Processor(_Total)\% Processor Time' | Select-Object -ExpandProperty CounterSamples | Select-Object InstanceName, CookedValue | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys cpu: %w\n%s", err, string(out))
	}
	printHeader("📊 CPU")
	fmt.Println(string(out))
	return nil
}

func sysRAM() error {
	out, err := runPowerShell(`Get-CimInstance Win32_OperatingSystem | Select-Object @{n='TotalGB';e={[math]::Round($_.TotalVisibleMemorySize/1MB,2)}}, @{n='FreeGB';e={[math]::Round($_.FreePhysicalMemory/1MB,2)}} | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys ram: %w\n%s", err, string(out))
	}
	out2, err2 := runPowerShell(`Get-Process | Sort-Object WorkingSet -Descending | Select-Object -First 10 Name,Id,CPU,@{n='MemMB';e={[math]::Round($_.WorkingSet/1MB,1)}} | Format-Table -AutoSize | Out-String`)
	if err2 != nil {
		return fmt.Errorf("failed to run odin sys ram (top procs): %w\n%s", err2, string(out2))
	}
	printHeader("📊 RAM")
	fmt.Println(string(out))
	fmt.Println(string(out2))
	return nil
}

func sysDisk() error {
	out, err := runPowerShell(`Get-Volume | Select-Object DriveLetter, FileSystemLabel, FileSystemType, SizeRemaining, Size | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys disk: %w\n%s", err, string(out))
	}
	printHeader("📊 Disk Usage")
	fmt.Println(string(out))
	return nil
}

func sysBoot() error {
	outBoot, err := runPowerShell(`$os = Get-CimInstance Win32_OperatingSystem; "LastBootUpTime: $($os.LastBootUpTime)"; (Get-Date) - $os.LastBootUpTime | Select-Object Days,Hours,Minutes,Seconds | Format-List | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys boot: %w\n%s", err, string(outBoot))
	}
	printHeader("Boot")
	fmt.Println(string(outBoot))

	outServices, err := runPowerShell(`Get-Service | Where-Object {$_.Status -eq 'Running'} | Select-Object -First 15 Name,DisplayName,Status | Format-Table -AutoSize | Out-String`)
	if err != nil {
		return fmt.Errorf("failed to run odin sys boot (services): %w\n%s", err, string(outServices))
	}
	printHeader("Running services (sample)")
	fmt.Println(string(outServices))
	return nil
}
