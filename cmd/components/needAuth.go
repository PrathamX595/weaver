package components

import (
    tea "github.com/charmbracelet/bubbletea"
)

type NeedAuth struct {
    need  bool
    ready bool
}

func InitAuth() NeedAuth {
    return NeedAuth{
        need:  false,
        ready: false,
    }
}

func (m NeedAuth) Init() tea.Cmd {
    return waitForReady()
}

func (m NeedAuth) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case readyMsg:
        m.ready = true
        return m, nil
	case tea.KeyMsg:
		if !m.ready {
            return m, nil
        }
		switch msg.String() {
		case "y":
			m.need = true
			return m, tea.Quit
		case "n":
			m.need = false
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m NeedAuth) View() string {
    if !m.ready {
		var s string
		for i := 0; i < 3; i++ {
			s += "."
		}
        return s
    }
    return "Do you need authentication? (y/n)"
}

func (m NeedAuth) NeedsAuth() bool {
    return m.need
}