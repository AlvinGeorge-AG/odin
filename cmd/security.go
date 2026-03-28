package cmd


import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
	"os"
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


var fireWall = &cobra.Command{
	Use:"firewall",
	Short:"Show ports that are externally exposed",
}


var statusCmd = &cobra.Command{
	Use:"status",
	Short:"To show only externally exposed ports",
	RunE : func(cmd *cobra.Command,args []string) error {
		if os.Getuid() != 0 {
			fmt.Println("❌ This command requires sudo. Run: sudo odin firewall status")
			os.Exit(1)
		}
		out,err := exec.Command("ufw","status").Output()
		if err!=nil {
			fmt.Errorf("Failed to Run odin firewall status : %w",err)
		}
		printHeader("Firewall Status")
		fmt.Println(string(out))
		return nil
	},
}



func init(){
	permOpenCmd.AddCommand(portCmd)
	rootCmd.AddCommand(permOpenCmd)

	fireWall.AddCommand(statusCmd)
	rootCmd.AddCommand(fireWall)
}