package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const chatDir = ".chats"

func chatSessionFilename(session ChatSession) string {
    return fmt.Sprintf("%s_%d.json", session.Name, session.Timestamp.Unix())
}

func SaveChat(session ChatSession) error {
    if _, err := os.Stat(chatDir); os.IsNotExist(err) {
        os.Mkdir(chatDir, 0755)
    }
    fileName := chatSessionFilename(session)
    path := filepath.Join(chatDir, fileName)

    data, err := json.MarshalIndent(session, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(path, data, 0644)
}


func LoadChatSessions() ([]ChatSession, error) {
	var sessions []ChatSession
	files, err := ioutil.ReadDir(chatDir)
	if err != nil {
		return sessions, err
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(chatDir, file.Name()))
		if err != nil {
			continue
		}
		var session ChatSession
		if json.Unmarshal(data, &session) == nil {
			sessions = append(sessions, session)
		}
	}
	return sessions, nil
}

func DeleteChat(index int) error {
	files, err := ioutil.ReadDir(chatDir)
	if err != nil {
		return err
	}
	if index < 0 || index >= len(files) {
		return fmt.Errorf("invalid index")
	}
	return os.Remove(filepath.Join(chatDir, files[index].Name()))
}

func ListChatFiles() ([]string, error) {
	files, err := ioutil.ReadDir(chatDir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, f := range files {
		names = append(names, f.Name())
	}
	return names, nil
}

