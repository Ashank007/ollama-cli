package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ollama-chat",
	Short: "CLI to chat with Ollama models",
}

func init() {
	rootCmd.AddCommand(chatCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

