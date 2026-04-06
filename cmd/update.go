package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Odin to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := http.Get("https://api.github.com/repos/AlvinGeorge-AG/odin/releases/latest")
		if err != nil {
			fmt.Println("Error:", err)
			return fmt.Errorf("failed to check updates: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return fmt.Errorf("failed to read response: %w", err)
		}
		var release GithubRelease
		json.Unmarshal(body, &release)

		if release.TagName != VERSION {
			fmt.Printf("⚡ New version available: %s\n", release.TagName)
			return runUpdateInstall()

		} else {
			fmt.Println("✅ Odin is up to date")
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
