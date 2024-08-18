package main

import (
	"fmt"

	"github.com/arikchakma/backend-projects/github-user-activity/cmd"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Padding(1, 2).
			Bold(true).
			Render(fmt.Sprintf("Error: %s", err))
		fmt.Println(errorStyle)
	}
}
