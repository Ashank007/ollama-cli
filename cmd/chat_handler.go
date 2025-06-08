package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


func ChatLoop(session ChatSession) {
    fmt.Printf("🤖 Chatting with %s — session: '%s' (type 'exit' to quit)\n", session.Model, session.Name)
    reader := bufio.NewReader(os.Stdin)

    history := session.Messages

    for {
        fmt.Print("You: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            break
        }

				if strings.HasPrefix(input, "search:") {
				query := strings.TrimSpace(strings.TrimPrefix(input, "search:"))
				fmt.Println("🌐 Searching for:", query)
				results, err := SearchWeb(query)
				if err != nil {
					fmt.Println("❌", err)
					continue
				}
				fmt.Println("🔍 Top Results:\n", results)
				input = "Here's some information I found:\n" + results
			}


        history = append(history, Message{Role: "user", Content: input})

        reply, err := SendChatRequest(session.Model, history)
        if err != nil {
            fmt.Println("❌ Error:", err)
            continue
        }

        history = append(history, Message{Role: "assistant", Content: reply})
    }

    session.Messages = history

    if session.Timestamp.IsZero() {
        session.Timestamp = time.Now()
    }

    SaveChat(session)
    fmt.Println("💾 Chat saved.")
}






