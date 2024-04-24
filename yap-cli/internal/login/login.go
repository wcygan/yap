package login

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wcygan/yap/yap-cli/internal/context"
	"strings"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	focusedButton       = focusedStyle.Copy().Render("[ Login ]")
	blurredButton       = blurredStyle.Copy().Render("[ Login ]")
	focusedCreateButton = focusedStyle.Copy().Render("[ Create Account ]")
	blurredCreateButton = blurredStyle.Copy().Render("[ Create Account ]")
)

type Model struct {
	*context.Context
	taskSpinner   spinner.Model
	usernameInput textinput.Model
	passwordInput textinput.Model
	inputs        []textinput.Model
	focusIndex    int
	loginButton   string
	createButton  string
	err           error
}

func InitialModel(ctx *context.Context) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}

	usernameInput := textinput.NewModel()
	usernameInput.Placeholder = "Username"
	usernameInput.Focus()
	usernameInput.CharLimit = 256
	passwordInput := textinput.NewModel()
	passwordInput.Placeholder = "Password"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.CharLimit = 256
	return Model{
		Context:       ctx,
		taskSpinner:   taskSpinner,
		usernameInput: usernameInput,
		passwordInput: passwordInput,
		inputs:        []textinput.Model{usernameInput, passwordInput},
		focusIndex:    0,
		loginButton:   blurredButton,
		createButton:  blurredCreateButton,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				if m.focusIndex == 2 {
					// TODO: Trigger login
				} else if m.focusIndex == 3 {
					// TODO: Trigger account creation
				}
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > 3 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = 3
			}

			cmds := make([]tea.Cmd, 2)
			for i := 0; i <= 1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = blurredStyle
				m.inputs[i].TextStyle = blurredStyle
			}

			if m.focusIndex == 2 {
				m.loginButton = focusedButton
				m.createButton = blurredCreateButton
			} else if m.focusIndex == 3 {
				m.loginButton = blurredButton
				m.createButton = focusedCreateButton
			} else {
				m.loginButton = blurredButton
				m.createButton = blurredCreateButton
			}

			return m, tea.Batch(cmds...)
		}

	// Handle text input changes
	case tea.WindowSizeMsg:
		m.usernameInput.Width = msg.Width - 4
		m.passwordInput.Width = msg.Width - 4
	}

	// Update text inputs
	m.usernameInput, _ = m.usernameInput.Update(msg)
	m.passwordInput, _ = m.passwordInput.Update(msg)

	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	// Render username input
	b.WriteString("Username\n")
	b.WriteString(m.usernameInput.View())
	b.WriteString("\n\n")

	// Render password input
	b.WriteString("Password\n")
	b.WriteString(m.passwordInput.View())
	b.WriteString("\n\n")

	// Render buttons
	b.WriteString("\n\n")
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, m.loginButton, m.createButton))

	// Render errors
	if m.err != nil {
		b.WriteString("\n\n")
		b.WriteString(m.err.Error())
	}

	return b.String()
}
