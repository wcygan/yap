package home

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/wcygan/yap/yap-cli/internal/application"
)

type Model struct {
	*application.Context
	taskSpinner spinner.Model
}

func InitialModel(ctx *application.Context) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}
	return Model{
		Context:     ctx,
		taskSpinner: taskSpinner,
	}
}
