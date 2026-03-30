package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var env = &cobra.Command{
	Use:   "env",
	Short: "",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show all env variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader("ENV")
		data := os.Environ()
		for _, line := range data {
			parts := strings.SplitN(line, "=", 2)
			fmt.Printf("%-20s = %s\n", parts[0], parts[1])
		}
		return nil
	},
}

var searchCmd = &cobra.Command{
	Use:   "find [term]",
	Short: "Find the given env variable if present",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("please provide a search term")
		}
		if len(args) > 1 {
			return fmt.Errorf("please provide only one search term")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader("Search Env")
		searchTerm := args[0]
		data := os.Environ()
		for _, line := range data {
			parts := strings.SplitN(line, "=", 2)
			if strings.Contains(strings.ToLower(parts[0]), strings.ToLower(searchTerm)) {
				fmt.Printf("%-20s = %s\n", parts[0], parts[1])
			}
		}
		return nil
	},
}

func init() {
	env.AddCommand(showCmd)
	env.AddCommand(searchCmd)
	rootCmd.AddCommand(env)
}
