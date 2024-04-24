package home

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/wcygan/yap/yap-cli/internal/context"
)

type Model struct {
	*context.Context
	taskSpinner     spinner.Model
	createChatInput textinput.Model
	joinChatInput   textinput.Model
}

func InitialModel(ctx *context.Context) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}
	return Model{
		Context:     ctx,
		taskSpinner: taskSpinner,
	}
}
