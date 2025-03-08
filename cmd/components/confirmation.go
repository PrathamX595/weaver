package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Confirm struct {
	confirm bool
	ready   bool
}

func InitConfirm() Confirm {
	return Confirm{
		confirm: false,
		ready:   false,
	}
}

func (m Confirm) Init() tea.Cmd {
	return waitForReady()
}

func (m Confirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.confirm = true
			return m, tea.Quit
		case "n":
			m.confirm = false
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Confirm) View() string {
	if !m.ready {
		var s string
		for i := 0; i < 3; i++ {
			s += "."
		}
		return s
	}
	return "Are you sure? (y/n)"
}

func (m Confirm) GetConfVal() bool {
	return m.confirm
}
