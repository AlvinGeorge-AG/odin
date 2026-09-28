package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Show all environment variables",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader(" 📦 ENV")
		data := os.Environ()
		for _, line := range data {
			parts := strings.SplitN(line, "=", 2)
			fmt.Printf("%-20s = %s\n", parts[0], parts[1])
		}
		return nil
	},
}

var envFindCmd = &cobra.Command{
	Use:   "env-find [term]",
	Short: "Find environment variable by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		searchTerm := args[0]
		data := os.Environ()
		printHeader(" 📦 Search Env")
		found := false
		for _, line := range data {
			parts := strings.SplitN(line, "=", 2)
			if strings.Contains(strings.ToLower(parts[0]), strings.ToLower(searchTerm)) {
				fmt.Printf("%-20s = %s\n", parts[0], parts[1])
				found = true
			}
		}
		if !found {
			fmt.Printf("No environment variables found matching '%s'\n", searchTerm)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(envFindCmd)
}