package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		case "down":
			m.count--
		}
	}
	return m, nil
}
func (m model) View() string {
	return fmt.Sprintf("Counter: %d\n\nPress q to quit.\n", m.count)
}
