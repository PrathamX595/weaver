package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
    choices  []string
    cursor   int
    selected string
    ready    bool
}

func InitList() List {
    return List{
        choices: []string{"Fiber", "Gin", "Standard Library(http)", "echo", "chi"},
        ready:   false,
    }
}

func (m List) Init() tea.Cmd {
    return waitForReady()
}

func (m List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case readyMsg:
        m.ready = true
        return m, nil
	case tea.KeyMsg:
		if !m.ready {
            return m, nil
        }
		switch msg.String() {
		case "ctrl+c", "c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if len(m.selected) >=1 {
				m.selected = ""
				m.selected = m.choices[m.cursor]
			}else{
				m.selected = m.choices[m.cursor]
			}
		}
	}
	return m, nil
}

func (m List) View() string {
	if !m.ready {
		var s string
		for i := 0; i < 3; i++ {
			s += "."
		}
        return s
    }
	s := "which framework would you like\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.selected == choice {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress c to confirm.\n"
	return s
}

func (m List) GetSelectedVal() string {
	return m.selected
}