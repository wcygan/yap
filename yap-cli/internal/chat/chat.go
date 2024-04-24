package chat

import (
	"github.com/wcygan/yap/yap-cli/internal/context"
)

type Model struct {
	*context.Context
}

func InitialModel(ctx *context.Context) Model {
	return Model{
		Context: ctx,
	}
}
