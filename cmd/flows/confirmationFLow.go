package flows

import (
	"weaver/cmd/components"

	tea "github.com/charmbracelet/bubbletea"
)

func ConformationFlow() bool {

	p := tea.NewProgram(components.InitConfirm(), tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		panic(err)
	}
	if l, ok := finalModel.(components.Confirm); ok {
		return l.GetConfVal()
	}
	return false
}
