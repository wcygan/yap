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
	stream        chatpb.ClientStreamingService_ChatStreamClient // the stream which we will use to send and receive messages
	streamContext ctx.Context                                    // the context for the stream
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
		stream:      nil,
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
		ChannelName: m.Context.GetChannelName(),
	}

	stream, err := m.Context.GetChatClient().ChatStream(m.streamContext)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Create a ChatPacket with the JoinChatRequest
	chatPacket := &chatpb.ChatPacket{
		PacketType: &chatpb.ChatPacket_JoinRequest{
			JoinRequest: joinChatRequest,
		},
	}
	err = stream.Send(chatPacket)
	// --- TODO: MOVE THIS CODE INTO CHAT ROOM
	// Do things like cancel the current context,
	// start a new context
	// start a new chat stream & update the model's send end
	// spawn a new goroutine that listens for messages & appends them to the messages slice
	return m, nil
}

func (m Model) StopCurrentChatRoom() {
	// Do things like cancel the current context
	// stop the current chat stream
	// close the send end
	m.Context.SetChannelName("")
	m.Context.SetCurrentPage(context.HomePage)
	m.streamContext.Done()
}
