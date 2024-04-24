package application

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wcygan/yap/yap-cli/internal/chat"
	"github.com/wcygan/yap/yap-cli/internal/context"
	"github.com/wcygan/yap/yap-cli/internal/home"
	"github.com/wcygan/yap/yap-cli/internal/login"
)

// Model is the top-level model for the application. It contains all necessary state for the application to function.
type Model struct {
	*context.Context             // The context for the application
	login            login.Model // The login page
	home             home.Model  // The home page
	chat             chat.Model  // The chat page
}

// InitialModel creates the initial model for the application. It will cause the application to start on the login page.
func InitialModel(host string) Model {
	ctx := context.InitialContext(host)
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

	// Handle global keybindings
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	// Process the pages
	switch m.Context.GetCurrentPage() {
	case context.LoginPage:
		loginModel, cmd := m.login.Update(msg)
		m.login = loginModel.(login.Model)
		return m, cmd
	default:
		panic("page is not implemented")

	}
	return m, cmd
}

func (m Model) View() string {
	switch m.Context.GetCurrentPage() {
	case context.LoginPage:
		return m.login.View()
	default:
		return "Hello, world!"
		//	no-op until we implement the other pages
	}
}
