package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2, 1, 2)
	btnStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).Padding(0, 2)
	errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

func (m Model) View() string {
	var s string
	switch m.currentScreen {
	case "login":
		s += "Username: " + m.username + "\n"
		s += "Password: " + m.password + "\n"
		s += btnStyle.Render("[Tab] Switch Input [Enter] Login")
		if m.loginError {
			s += errStyle.Render("\nLogin Failed. Try again.")
		}
	case "home":
		s += "Welcome " + m.username + "!\n"
		s += btnStyle.Render("[1] Create Chat [2] Join Chat [3] Logout")
	case "chat":
		s += "Chat: " + m.chatroom + "\n"
		s += "Message: " + m.message + "\n"
		s += btnStyle.Render("[Enter] Send Message [Home] Return to Home")
	}
	return docStyle.Render(s)
}
