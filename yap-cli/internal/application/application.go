package application

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wcygan/yap/yap-cli/internal/chat"
	"github.com/wcygan/yap/yap-cli/internal/context"
	"github.com/wcygan/yap/yap-cli/internal/home"
	"github.com/wcygan/yap/yap-cli/internal/login"
)

type Model struct {
	*context.Context
	login login.Model
	home  home.Model
	chat  chat.Model
}

func InitialModel() Model {
	ctx := context.InitialContext()
	return Model{
		Context: ctx, // Share the context with each page
		login:   login.InitialModel(ctx),
		home:    home.InitialModel(ctx),
		chat:    chat.InitialModel(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.Context.GetCurrentPage() {
	case context.Login:
		m.login, cmd = m.login.Update(msg)
	default:
		panic("page is not implemented")

	}
	return m, cmd
}

func (m Model) View() string {
	switch m.Context.GetCurrentPage() {
	case context.Login:
		return m.login.View()
	default:
		//	no-op until we implement the other pages
	}

	panic("unhandled default case")
}
