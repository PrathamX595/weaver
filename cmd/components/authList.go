package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type AuthList struct {
    choices  []string
    cursor   int
    selected map[int]struct{}
    ready    bool
}

func InitAuthList() AuthList {
    return AuthList{
        choices:  []string{"Google", "GitHub", "discord"},
        selected: make(map[int]struct{}),
        ready:    false,
    }
}

func (m AuthList) Init() tea.Cmd {
    return waitForReady()
}

func (m AuthList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case readyMsg:
        m.ready = true
        return m, nil
	case tea.KeyMsg:
		if !m.ready {
            return m, nil
        }
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "c":
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
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}
	}
	return m, nil
}

func (m AuthList) View() string {
	if !m.ready {
		var s string
		for i := 0; i < 3; i++ {
			s += "."
		}
        return s
    }
	s := "Which Auths shoud be included\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress c to confirm.\n"
	s += "\nPress esc to cancel.\n"
	return s
}

func (m AuthList) GetSelectedAuthVals() []string {
	var selected []string
	for i := range m.selected {
		selected = append(selected, m.choices[i])
	}
	return selected
}
