package login

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wcygan/yap/yap-cli/internal/context"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton       = focusedStyle.Copy().Render("[ Login ]")
	blurredButton       = fmt.Sprintf("[ %s ]", blurredStyle.Render("Login"))
	focusedCreateButton = focusedStyle.Copy().Render("[ Create Account ]")
	blurredCreateButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Create Account"))
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
	cursorMode    cursor.Mode
}

func InitialModel(ctx *context.Context) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}

	usernameInput := textinput.New()
	usernameInput.Placeholder = "Username"
	usernameInput.Focus()
	usernameInput.CharLimit = 256
	usernameInput.Cursor.Style = cursorStyle

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Password"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.EchoCharacter = '•'
	passwordInput.CharLimit = 256
	passwordInput.Cursor.Style = cursorStyle

	return Model{
		Context:       ctx,
		taskSpinner:   taskSpinner,
		usernameInput: usernameInput,
		passwordInput: passwordInput,
		inputs:        []textinput.Model{usernameInput, passwordInput},
		focusIndex:    0,
		loginButton:   blurredButton,
		createButton:  blurredCreateButton,
		cursorMode:    cursor.CursorBlink,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.err = nil
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyCtrlR:
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			s := msg.String()

            if s == "enter" && m.focusIndex == len(m.inputs) {
                // Login button pressed
                // TODO: Perform login action
                m.err = fmt.Errorf("incorrect login credentials")
                return m, nil
            } else if s == "enter" && m.focusIndex == len(m.inputs)+1 {
                // Create account button pressed
                // TODO: Perform create account action
                // Stub the API call and assume an error occurred
                m.err = fmt.Errorf("account already exists")
                return m, nil
            }

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs)+1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) + 1
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			if m.focusIndex == len(m.inputs) {
				m.loginButton = focusedButton
				m.createButton = blurredCreateButton
			} else if m.focusIndex == len(m.inputs)+1 {
				m.loginButton = blurredButton
				m.createButton = focusedCreateButton
			} else {
				m.loginButton = blurredButton
				m.createButton = blurredCreateButton
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	b.WriteString("\n\n")
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, m.loginButton, m.createButton))

	b.WriteString("\n\n")
	b.WriteString(blurredStyle.Render("1. navigate with tab, shift+tab, and enter"))
	b.WriteString(blurredStyle.Render("\n2. cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(blurredStyle.Render(" (ctrl+r to change style)\n"))
	

	if m.err != nil {
		b.WriteString("\n\n")
		b.WriteString(m.err.Error())
	} else {
		b.WriteString("\n\n")
	}

	return b.String()
}