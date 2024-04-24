package home

import (
	"fmt"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wcygan/yap/yap-cli/internal/context"
	"strings"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	focusedCreateButton = focusedStyle.Copy().Render("[ Create Chat ]")
	blurredCreateButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Create Chat"))
	focusedJoinButton   = focusedStyle.Copy().Render("[ Join Chat ]")
	blurredJoinButton   = fmt.Sprintf("[ %s ]", blurredStyle.Render("Join Chat"))
	focusedLogoutButton = focusedStyle.Copy().Render("[ Logout ]")
	blurredLogoutButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Logout"))
)

type Model struct {
	*context.Context
	taskSpinner   spinner.Model
	chatNameInput textinput.Model
	createButton  string
	joinButton    string
	logoutButton  string
	focusIndex    int
	err           error
	cursorMode    cursor.Mode
}

func InitialModel(ctx *context.Context) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}

	chatNameInput := textinput.New()
	chatNameInput.Placeholder = "Name of Chat Room"
	chatNameInput.Focus()
	chatNameInput.CharLimit = 256
	chatNameInput.Cursor.Style = cursorStyle

	return Model{
		Context:       ctx,
		taskSpinner:   taskSpinner,
		chatNameInput: chatNameInput,
		focusIndex:    0,
		createButton:  blurredCreateButton,
		joinButton:    blurredJoinButton,
		logoutButton:  blurredLogoutButton,
		cursorMode:    cursor.CursorBlink,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			s := msg.String()

			if s == "enter" {
				if m.isCreateChatFocused() {
					// Create chat button pressed
					// TODO: Implement create chat functionality
				} else if m.isJoinChatFocused() {
					// Join chat button pressed
					// TODO: Implement join chat functionality
				} else if m.isLogoutFocused() {
					m.Context.Logout()
					return m, nil
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

			m.createButton = blurredCreateButton
			m.joinButton = blurredJoinButton
			m.logoutButton = blurredLogoutButton
			if m.isTextInputFocused() {
				cmds = append(cmds, m.chatNameInput.Focus())
				m.chatNameInput.PromptStyle = focusedStyle
				m.chatNameInput.TextStyle = focusedStyle
			} else {
				m.chatNameInput.Blur()
				m.chatNameInput.PromptStyle = noStyle
				m.chatNameInput.TextStyle = noStyle
				if m.isCreateChatFocused() {
					m.createButton = focusedCreateButton
				} else if m.isJoinChatFocused() {
					m.joinButton = focusedJoinButton
				} else {
					m.logoutButton = focusedLogoutButton
				}
			}
		}

	}

	// Handle character input and blinking
	var cmd tea.Cmd
	m.chatNameInput, cmd = m.chatNameInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Hello, %s!\n\n", m.Context.GetLoginInformation().Username))

	b.WriteString(m.chatNameInput.View())
	b.WriteRune('\n')

	b.WriteString("\n\n")
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, m.createButton, m.joinButton, m.logoutButton))

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

func (m Model) isTextInputFocused() bool {
	return m.focusIndex == 0
}

func (m Model) isCreateChatFocused() bool {
	return m.focusIndex == 1
}

func (m Model) isJoinChatFocused() bool {
	return m.focusIndex == 2
}

func (m Model) isLogoutFocused() bool {
	return m.focusIndex == 3
}
