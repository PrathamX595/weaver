package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type Textmodel struct {
	textInput textinput.Model
	err       error
	ready     bool
}

func InitText() Textmodel {
	ti := textinput.New()
	ti.Placeholder = "MyProj"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Textmodel{
		textInput: ti,
		err:       nil,
		ready:     false,
	}
}

func (m Textmodel) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		waitForReady(),
	)
}

func (m Textmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case readyMsg:
		m.ready = true
		return m, nil
	case tea.KeyMsg:
		if !m.ready {
			return m, nil
		}

		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Textmodel) View() string {
	if !m.ready {
		var s string
		for i := 0; i < 3; i++ {
			s += "."
		}
		return s
	}

	return fmt.Sprintf(
		"Name your project\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to cancel)",
	) + "\n"
}

func (m Textmodel) GetProjName() string {
	return m.textInput.Value()
}
