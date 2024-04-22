package main

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/wcygan/yap/yap-cli/internal/ui"
	"os"
)

func main() {
	p := tea.NewProgram(ui.InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		os.Exit(1)
	}
}
