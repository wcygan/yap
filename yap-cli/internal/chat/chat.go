package chat

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	chatpb "github.com/wcygan/yap/generated/go/chat/v1"
	"github.com/wcygan/yap/yap-cli/internal/context"
	ctx "golang.org/x/net/context"
	"strings"
)

type (
	errMsg error
)

type Model struct {
	*context.Context
	viewport      viewport.Model
	messages      []string
	textarea      textarea.Model
	senderStyle   lipgloss.Style
	err           error
	homeButton    string
	streamContext ctx.Context
}

type StartNewChatRoomMsg struct{}

func InitialModel(ctx *context.Context) Model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(60)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(60, 5)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return Model{
		Context:     ctx,
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		homeButton:  "[ Homepage ]",
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.ShouldStartNewChatRoom() {
		// Need to figure out how to this code run immediately after the page transition
		m.messages = append(m.messages, "Starting new chat room...")
		return m.StartNewChatRoom()
	}

	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		case tea.KeyCtrlH:
			m.StopCurrentChatRoom()
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
		"Press Ctrl+H to return to the homepage",
	) + "\n\n"
}

func (m Model) StartNewChatRoom() (tea.Model, tea.Cmd) {
	m.Context.SetShouldStartNewChatRoom(false)
	m.streamContext = ctx.Background()
	joinChatRequest := &chatpb.JoinChatRequest{
		UserId:      m.Context.GetLoginInformation().UserId,
		UserName:    m.Context.GetLoginInformation().Username,
		ChannelName: m.Context.GetChannelName(),
	}

	stream, err := m.Context.GetChatRoomClient().JoinChatRoom(m.streamContext, joinChatRequest)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Spawn a new goroutine that listens for messages & appends them to the messages slice
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				m.err = err
				return
			}

			switch contents := in.Contents.(type) {
			case *chatpb.Packet_Message:
				// TODO: Add pretty colorful messages based on the sender
				msg := fmt.Sprintf("[%s] %s", contents.Message.UserName, contents.Message.Message)
				m.messages = append(m.messages, msg)
			case *chatpb.Packet_UserJoined:
				msg := fmt.Sprintf(">>> %s joined the chat room <<<", contents.UserJoined.UserName)
				m.messages = append(m.messages, msg)
			case *chatpb.Packet_UserLeft:
				msg := fmt.Sprintf("<<< %s left the chat room >>>", contents.UserLeft.UserName)
				m.messages = append(m.messages, msg)
			default:
				fmt.Println("Unknown packet contents")
			}
		}
	}()

	return m, nil
}

func (m Model) StopCurrentChatRoom() {
	// Leave the chat room
	m.Context.SetChannelName("")

	// Go back to the home page
	m.Context.SetCurrentPage(context.HomePage)

	// Close the stream
	m.streamContext.Done()

	// Clear the messages
	m.messages = []string{}
}
