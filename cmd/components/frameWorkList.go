package components

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type ListModel struct {
	list     list.Model
	selected string
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func InitList() ListModel {
	items := []list.Item{
		item{title: "Chi", desc: "Lightweight, idiomatic and composable router for Go"},
		item{title: "Echo", desc: "High performance, minimalist web framework"},
		item{title: "Fiber", desc: "Express-inspired web framework built on Fasthttp"},
		item{title: "Http", desc: "Standard library http package for Go"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a framework"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().MarginLeft(2)

	return ListModel{list: l}
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			// Store selected item when Enter is pressed
			if i, ok := m.list.SelectedItem().(item); ok {
				m.selected = i.Title()
			}
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {
	return docStyle.Render(m.list.View())
}

func (m ListModel) GetSelectedVal() string {
	return m.selected
}
