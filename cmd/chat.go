package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/spf13/cobra"
)

var model string

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start, resume or delete chat sessions",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewScanner(os.Stdin)
		for {
			fmt.Println("\n🧠 What would you like to do?")
			fmt.Println("[1] Start a new chat")
			fmt.Println("[2] Continue a previous chat")
			fmt.Println("[3] Delete a chat")
			fmt.Println("[4] Exit")
			fmt.Print("Choose an option: ")

			var option int
			_, err := fmt.Scanln(&option)
			if err != nil {
				fmt.Println("❌ Invalid input.")
				reader.Scan()
				continue
			}

			switch option {
			case 1:
				startNewChat(reader)
			case 2:
				resumeChat()
			case 3:
				deleteChatInteractive()
			case 4:
				return
			default:
				fmt.Println("❌ Invalid choice.")
			}
		}
	},
}

func startNewChat(reader *bufio.Scanner) {
	fmt.Print("Enter new chat name: ")
	name := readLine(reader)

	m := model
	if m == "" {
		var err error
		m, err = PickModel()
		if err != nil {
			fmt.Println("❌", err)
			return
		}
	}
    session := ChatSession{
        Name:      name,
        Model:     m,
        Messages:  []Message{},
        Timestamp: time.Now(),
    }
    ChatLoop(session)

}

func resumeChat() {
	sessions, err := LoadChatSessions()
	if err != nil || len(sessions) == 0 {
		fmt.Println("📭 No chats available.")
		return
	}

	for i, s := range sessions {
		fmt.Printf("[%d] %s | Model: %s | %s\n", i+1, s.Name, s.Model, s.Timestamp.Format("Jan 2 15:04"))
	}

	fmt.Print("Enter chat number to resume: ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(sessions) {
		fmt.Println("❌ Invalid choice.")
		return
	}

	selected := sessions[choice-1]

	 fmt.Println("\n🕰️ Previous chat history:")
    for _, msg := range selected.Messages {
        fmt.Printf("%s: %s\n", msg.Role, msg.Content)
    }

  fmt.Println("\n🚀 Resuming chat...")
  ChatLoop(selected)

}

func deleteChatInteractive() {
	files, err := ListChatFiles()
	if err != nil || len(files) == 0 {
		fmt.Println("📭 No chats to delete.")
		return
	}

	for i, f := range files {
		fmt.Printf("[%d] %s\n", i+1, f)
	}

	fmt.Print("Enter chat number to delete: ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(files) {
		fmt.Println("❌ Invalid choice.")
		return
	}

	err = DeleteChat(choice - 1)
	if err != nil {
		fmt.Println("❌ Failed to delete:", err)
	} else {
		fmt.Println("✅ Chat deleted.")
	}
}

func readLine(reader *bufio.Scanner) string {
	if reader.Scan() {
		return strings.TrimSpace(reader.Text())
	}
	return ""
}


