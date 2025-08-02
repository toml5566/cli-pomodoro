package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/toml5566/cli-pomodoro/internal/app"
)

func main() {
	p := tea.NewProgram(
		app.InitialModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running %s: %v\n", app.PROJECT_NAME, err)
		os.Exit(1)
	}
}
