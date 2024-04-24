package chat

import (
	"github.com/wcygan/yap/yap-cli/internal/application"
)

type Model struct {
	*application.Context
}

func InitialModel(ctx *application.Context) Model {
	return Model{
		Context: ctx,
	}
}
