package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Summary struct {
	Name      string
	Framework string
	Auth      []string
	Confirmation bool
	ready     bool
}

type FinalSummary struct {
	Name      string
	Framework string
	Auth      []string
}

func InitSummary(name string, framework string, auth []string) Summary {
	return Summary{
		Name:      name,
		Framework: framework,
		Auth:      auth,
		ready:     false,
	}
}

func (m Summary) Init() tea.Cmd {
	return waitForReady()
}

func (m Summary) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.Confirmation = true
			return m, tea.Quit
		case "n":
			m.Confirmation = false
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Summary) View() string {
    if !m.ready {
		var s string
		for i := 0; i < 3; i++ {
			s += "."
		}
        return s
    }

    s := "\n=== Final Config ===\n\n"
    s += fmt.Sprintf("Project name: %s\n", m.Name)
    s += fmt.Sprintf("Framework: %s\n", m.Framework)

    s += "Authentication:\n"
    if len(m.Auth) == 0 {
        s += "   - None\n"
    } else {
        for _, auth := range m.Auth {
            s += fmt.Sprintf("   - %s\n", auth)
        }
    }

    s += "\nLook good? (y/n): "
    return s
}

func (m Summary) GetConfirmation() bool {
	return m.Confirmation
}