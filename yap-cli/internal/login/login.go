package login

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wcygan/yap/yap-cli/internal/context"
)

type Model struct {
	*context.Context
	taskSpinner   spinner.Model
	usernameInput textinput.Model
	passwordInput textinput.Model
}

func InitialModel(ctx *context.Context) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}
	return Model{
		Context:     ctx,
		taskSpinner: taskSpinner,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "Login"
}
