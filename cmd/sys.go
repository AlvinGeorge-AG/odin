package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
)

var sysCmd = &cobra.Command{
	Use :"sys",
	Short:"Displaying System Informations",
}
	
func printHeader(data string){
	fmt.Printf("─────────────────────────────\nODIN · %s Info\n─────────────────────────────\n",data)
}

var infoCmd = &cobra.Command{
	Use:"info",
	Short:"display information about the CPU architecture",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command,args []string) error {
		out,err := exec.Command("lscpu").Output()
		out1,err1 := exec.Command("free","-h").Output()
		out2,err2 := exec.Command("df","-h").Output()
		out3,err3 := exec.Command("uname","-a").Output()
		if err!=nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w",err)
		}
		if err1!=nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w",err1)
		}
		if err2!=nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w",err2)
		}
		if err3!=nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w",err3)
		}
		printHeader("User");
		fmt.Println(string(out3))
		printHeader("System");
		fmt.Println(string(out))
		printHeader("Memory");
		fmt.Println(string(out1))
		printHeader("Disk");
		fmt.Println(string(out2))
		return nil
	},
}

var tempCmd = &cobra.Command{
	Use:"temp",
	Short:"CPU/GPU temperatures",
	RunE : func(cmd *cobra.Command,args []string) error {
		out ,err := exec.Command("sensors").Output()
		if err!=nil {
			fmt.Errorf("Failed to Run odin sys temp : %w",err)
		}
		printHeader("Temperature")
		fmt.Println(string(out))
		return nil
	},
}

var cpuCmd = &cobra.Command{
	Use:"cpu",
	Short:"Show real-time CPU usage per core",
	RunE : func(cmd *cobra.Command,args []string) error {
		out1 ,err1 := exec.Command("sh","-c","top -bn1 | grep %Cpu").Output()
		if err1!=nil {
			return fmt.Errorf("Failed to Run odin sys cpu : %w",err1)
		}
		printHeader("CPU")
		fmt.Println(string(out1))
		return nil
	},
}

var memCmd = &cobra.Command{
	Use:"ram",
	Short:"Show real-time memory usage",
	RunE : func(cmd *cobra.Command,args []string) error {
		out1 ,err1 := exec.Command("sh","-c","top -bn1 | grep Mem").Output()
		out2,err2 := exec.Command("sh","-c","ps -eo user,pid,pcpu,pmem,comm --sort=-%mem | head").Output()
		if err1!=nil  {
			return fmt.Errorf("Failed to Run odin sys ram : %w",err1)
		}
		if err2!=nil  {
			return fmt.Errorf("Failed to Run odin sys ram : %w",err2)
		}
		printHeader("RAM")
		fmt.Println(string(out1))
		fmt.Println(string(out2))
		return nil
	},
}


var diskCmd = &cobra.Command{
	Use:"disk",
	Short:"Disk usage, no tmpfs noise",
	RunE : func(cmd *cobra.Command,args []string) error {
		out ,err := exec.Command("df","-h").Output()
		if err!=nil {
			fmt.Errorf("Failed to Run odin sys disk: %w",err)
		}
		printHeader("Disk Usage")
		fmt.Println(string(out))
		return nil
	},
}


func init(){
	sysCmd.AddCommand(infoCmd)
	sysCmd.AddCommand(tempCmd)
	sysCmd.AddCommand(cpuCmd)
	sysCmd.AddCommand(memCmd)
	sysCmd.AddCommand(diskCmd)
	rootCmd.AddCommand(sysCmd)
}