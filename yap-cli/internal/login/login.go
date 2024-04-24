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
			// Handle focus cycling and input activation
			if msg.String() == "tab" || msg.String() == "shift+tab" {
				if msg.String() == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}
				if m.focusIndex >= len(m.inputs) + 2 { // Include two buttons in the focus cycle
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs) + 1
				}
			}
			// Update focus and style for inputs and buttons
			for i, input := range m.inputs {
				if i == m.focusIndex {
					input.Focus()
					input.PromptStyle = focusedStyle
					input.TextStyle = focusedStyle
				} else {
					input.Blur()
					input.PromptStyle = blurredStyle
					input.TextStyle = blurredStyle
				}
			}
			if m.focusIndex == len(m.inputs) { // Login button
				m.loginButton = focusedButton
				m.createButton = blurredCreateButton
			} else if m.focusIndex == len(m.inputs) + 1 { // Create account button
				m.loginButton = blurredButton
				m.createButton = focusedCreateButton
			} else {
				m.loginButton = blurredButton
				m.createButton = blurredCreateButton
			}
			return m, nil
		}

	// Handle text input changes
	case tea.WindowSizeMsg:
		m.usernameInput.Width = msg.Width - 4
		m.passwordInput.Width = msg.Width - 4
	}

	// Update text inputs outside the switch-case for KeyMsg
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
