package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var (
	bannerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#3498DB"))
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1C40F")).Bold(true)
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#7F8C8D")).Italic(true)
)

type Textmodel struct {
	textInput textinput.Model
	err       error
	ready     bool
	escaped   bool
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
		escaped:   false,
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
		case tea.KeyEsc:
			m.escaped = true
			return m, tea.Quit
		case tea.KeyEnter, tea.KeyCtrlC:
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
		var dots strings.Builder
		dots.WriteString("\n\n  ")
		for i := 0; i < 3; i++ {
			dots.WriteString(".")
		}
		dots.WriteString("\n\n")
		return dots.String()
	}

	banner := bannerStyle.Render(`
    ____    __    ____  _______       ___   ____    ____  _______  ______      
    \   \  /  \  /   / |   ____|     /   \  \   \  /   / |   ____||   _  \     
     \   \/    \/   /  |  |__       /  ^  \  \   \/   /  |  |__   |  |_)  |    
      \            /   |   __|     /  /_\  \  \      /   |   __|  |      /     
       \    /\    /    |  |____   /  _____  \  \    /    |  |____ |  |\  \
        \__/  \__/     |_______| /__/     \__\  \__/     |_______||__| \__|
        
    `)

	var view strings.Builder
	view.WriteString(banner)
	view.WriteString("\n")
	view.WriteString(titleStyle.Render("Name your project"))
	view.WriteString("\n\n")
	view.WriteString(m.textInput.View())
	view.WriteString("\n\n")
	view.WriteString(helpStyle.Render("(esc to exit program, enter to confirm)"))
	view.WriteString("\n")

	return view.String()
}

func (m Textmodel) GetProjName() string {
	return m.textInput.Value()
}

func (m Textmodel) WasEscaped() bool {
	return m.escaped
}
