package flows

import (
    "github.com/PrathamX595/weaver/cmd/components"
    tea "github.com/charmbracelet/bubbletea"
)

func TextFlow() string {
    p := tea.NewProgram(components.InitText(), tea.WithAltScreen())
    finalModel, err := p.Run()
    if err != nil {
        panic(err)
    }
    if tm, ok := finalModel.(components.Textmodel); ok {
        return tm.GetProjName()
    }
    return ""
}