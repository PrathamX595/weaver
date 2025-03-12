package flows

import (
	"github.com/PrathamX595/weaver/cmd/components"

	tea "github.com/charmbracelet/bubbletea"
)

func AuthListFlow() []string {
    p := tea.NewProgram(components.InitAuthList(), tea.WithAltScreen())
    finalModel, err := p.Run()
    if err != nil {
        panic(err)
    }
    if aList, ok := finalModel.(components.AuthList); ok {
        return aList.GetSelectedAuthVals()
    }
    return nil
}
