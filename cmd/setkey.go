package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setKeyCmd = &cobra.Command{
	Use:   "set-key",
	Short: "Set your SerpAPI key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Please provide your SerpAPI key.")
			return
		}
		err := SaveAPIKey(args[0])
		if err != nil {
			fmt.Println("❌ Failed to save API key:", err)
		} else {
			fmt.Println("✅ API key saved.")
		}
	},
}

