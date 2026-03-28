package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
)

var syscmd = &cobra.Command{
	Use :"sys",
	Short:"Displaying System Informations",
}
	


var infocmd = &cobra.Command{
	Use:"info",
	Short:"display information about the CPU architecture",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command,args []string) error {
		out,err := exec.Command("lscpu").Output()
		if err!=nil {
			return fmt.Errorf("Failed to Run lscpu! : %w",err)
		}
		fmt.Println(string(out))
		return nil
	},
}

func init(){
	syscmd.AddCommand(infocmd)
	rootcmd.AddCommand(syscmd)
}