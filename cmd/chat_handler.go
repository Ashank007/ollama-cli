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






