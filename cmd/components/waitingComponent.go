package components

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type readyMsg struct{}

func waitForReady() tea.Cmd {
	return tea.Tick(5*time.Millisecond, func(t time.Time) tea.Msg {
        return readyMsg{}
    })
}
