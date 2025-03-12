package flows

import (
    "github.com/PrathamX595/weaver/cmd/components"
    tea "github.com/charmbracelet/bubbletea"
)

func ListFlow() string {
    p := tea.NewProgram(components.InitList(), tea.WithAltScreen())
    finalModel, err := p.Run()
    if err != nil {
        panic(err)
    }
    if l, ok := finalModel.(components.List); ok {
        return l.GetSelectedVal()
    }
    return ""
}