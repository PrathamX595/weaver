package flows

import (
	"github.com/PrathamX595/weaver/cmd/components"

	tea "github.com/charmbracelet/bubbletea"
)

func SummarizeFlow(name, framework string, auth []string) bool {
	p := tea.NewProgram(components.InitSummary(name, framework, auth), tea.WithAltScreen())
	finalModel , err := p.Run() 
	if err != nil {
		panic(err)
	}
	if l, ok := finalModel.(components.Summary); ok {
		return l.GetConfirmation()
	}
	return false
}
