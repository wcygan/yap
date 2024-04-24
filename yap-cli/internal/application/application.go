package application

import (
	"github.com/wcygan/yap/yap-cli/internal/chat"
	"github.com/wcygan/yap/yap-cli/internal/home"
	"github.com/wcygan/yap/yap-cli/internal/login"
)

type Model struct {
	*Context
	login login.Model
	home  home.Model
	chat  chat.Model
}

func InitialModel() Model {
	ctx := &Context{
		loginInformation: nil,   // No login information to start; the user must log in
		currentPage:      Login, // Start on the login page
	}
	return Model{
		Context: ctx, // Share the context with each page
		login:   login.InitialModel(ctx),
		home:    home.InitialModel(ctx),
		chat:    chat.InitialModel(ctx),
	}
}
