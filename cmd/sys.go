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

func init(){
	sysCmd.AddCommand(infoCmd)
	sysCmd.AddCommand(tempCmd)
	rootCmd.AddCommand(sysCmd)
}