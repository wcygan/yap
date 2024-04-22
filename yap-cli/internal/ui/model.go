package ui

import (
	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	currentScreen  string
	username       string
	password       string
	chatroom       string
	message        string
	inputFocus     string // "username", "password", "chatroom", "message"
	loginError     bool
	accountCreated bool
}

func InitialModel() Model {
	return Model{currentScreen: "login", inputFocus: "username"}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m.handleEnter()
		case tea.KeyTab:
			return m.switchFocus(), nil
		}

		// Handle character input based on the focused input field
		if msg.Type == tea.KeyRunes {
			return m.handleInput(msg.Runes[0]), nil
		}
	}
	return m, nil
}

// switchFocus cycles the input focus through the available fields
func (m Model) switchFocus() tea.Model {
	switch m.currentScreen {
	case "login":
		if m.inputFocus == "username" {
			m.inputFocus = "password"
		} else {
			m.inputFocus = "username"
		}
	case "home":
		// No inputs to focus on the home screen
	case "chat":
		if m.inputFocus == "message" {
			m.inputFocus = "message" // only one input in chat
		}
	}
	return m
}

// handleInput processes keyboard input into the correct field
func (m Model) handleInput(r rune) tea.Model {
	switch m.inputFocus {
	case "username":
		m.username += string(r)
	case "password":
		m.password += string(r)
	case "chatroom":
		m.chatroom += string(r)
	case "message":
		m.message += string(r)
	}
	return m
}

// handleEnter simulates actions based on current screen and input
func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case "login":
		return m.tryLogin()
	case "home":
		// Additional logic for home screen options
	case "chat":
		// Send the message in the chatroom
	}
	return m, nil
}

// tryLogin simulates a login attempt
func (m Model) tryLogin() (tea.Model, tea.Cmd) {
	// This is where you would implement actual authentication logic
	if m.username == "user" && m.password == "pass" { // Placeholder
		m.currentScreen = "home"
	} else {
		m.loginError = true
	}
	return m, nil
}
