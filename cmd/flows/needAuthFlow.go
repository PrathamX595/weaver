package flows

import (
	"github.com/PrathamX595/weaver/cmd/components"

	tea "github.com/charmbracelet/bubbletea"
)

func NeedAuthFlow() bool {
	p := tea.NewProgram(components.InitAuth(), tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		panic(err)
	}
	if a, ok := finalModel.(components.NeedAuth); ok {
		return a.NeedsAuth()
	}
	return false
}
