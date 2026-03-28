package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
	"os"
)


var cleanCmd = &cobra.Command{
	Use:"clean",
	Short:"",
}

var aptCmd = &cobra.Command{
	Use:"apt",
	Short:"This command removes packages that were automatically installed to satisfy dependencies of other packages but are no longer required by any currently installed software",
	RunE : func(cmd *cobra.Command,args []string) error {
		if os.Getuid() != 0 {
			fmt.Println("❌ This command requires sudo. Run: sudo odin clean apt")
			os.Exit(1)
		}
		out,err := exec.Command("sh","-c","apt autoremove && sudo apt clean").Output()
		if err!=nil {
			return fmt.Errorf("Failed to Run Odin Info! : %w",err)
		}
		printHeader("Clean");
		fmt.Println(string(out))
		return nil;
	},
}



var cacheCmd  = &cobra.Command{
	Use:"cache",
	Short:"Clears the Cached Files",
	RunE : func(cmd *cobra.Command,args []string) error {
		out,err := exec.Command("sh","-c","rm -rf ~/.cache/thumbnails/* && rm -rf ~/.cache/*").Output()
		if err!=nil {
			return fmt.Errorf("Failed to Run Odin clean cache! : %w",err)
		}
		printHeader("Cache Clean");
		fmt.Println(string(out))
		return nil;
	},
}


func init(){
	cleanCmd.AddCommand(aptCmd)
	cleanCmd.AddCommand(cacheCmd)
	rootCmd.AddCommand(cleanCmd)
}
