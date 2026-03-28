package cmd


import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
)

var permOpenCmd = &cobra.Command{
	Use:"open",
	Short:"Show ports that are externally exposed",
}

var portCmd = &cobra.Command{
	Use:"ports",
	Short:"To show only externally exposed ports",
	RunE : func(cmd *cobra.Command,args []string) error {
		out,err := exec.Command("sh","-c","ss -tuln | awk 'NR==1 || $5 ~ /^0\\.0\\.0\\.0/ || $5 ~ /^\\[::\\]/'").Output()
		if err!=nil {
			fmt.Errorf("Failed to Run odin open ports : %w",err)
		}
		printHeader("Open Ports")
		fmt.Println(string(out))
		return nil
	},
}



func init(){
	permOpenCmd.AddCommand(portCmd)
	rootCmd.AddCommand(permOpenCmd)
}