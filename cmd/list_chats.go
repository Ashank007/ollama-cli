package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listChatsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved chat sessions",
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := LoadChatSessions()
		if err != nil {
			fmt.Println("❌ Error loading chats:", err)
			return
		}

		if len(sessions) == 0 {
			fmt.Println("📭 No chat sessions found.")
			return
		}

		fmt.Println("🗃️ Saved Chats:")
		for i, s := range sessions {
			fmt.Printf("[%d] %s | Model: %s | %s\n", i+1, s.Name, s.Model, s.Timestamp.Format("Jan 2 15:04"))
		}
	},
}

func init() {
	rootCmd.AddCommand(listChatsCmd)
}

