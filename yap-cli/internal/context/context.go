package context

import "sync"

// Context contains the shared state of the application
type Context struct {
	sync.RWMutex                       // The context is shared, so we need to protect it
	host             string            // The host server address
	loginInformation *LoginInformation // The user's login information
	currentPage      Page              // The current page the user is on
}

// LoginInformation contains information for the authentication lifecycle
type LoginInformation struct {
	Username     string // The user's username
	AccessToken  string // The token which authenticates the user's requests
	RefreshToken string // The token which refreshes the access token as needed
}

type Page = int

const (
	LoginPage Page = iota
	HomePage
	ChatPage
)

func InitialContext(host string) *Context {
	return &Context{
		host:             host,      // Set the host server address for gRPC communication
		loginInformation: nil,       // No login information to start; the user must log in
		currentPage:      LoginPage, // Start on the login page
	}
}

func (c *Context) GetHost() string {
	c.RLock()
	defer c.RUnlock()
	return c.host
}

func (c *Context) SetCurrentPage(page Page) {
	c.Lock()
	defer c.Unlock()
	c.currentPage = page
}

func (c *Context) GetCurrentPage() Page {
	c.RLock()
	defer c.RUnlock()
	return c.currentPage
}

func (c *Context) SetLoginInformation(loginInformation *LoginInformation) {
	c.Lock()
	defer c.Unlock()
	c.loginInformation = loginInformation
}

func (c *Context) GetLoginInformation() *LoginInformation {
	c.RLock()
	defer c.RUnlock()
	return c.loginInformation
}

func (c *Context) Logout() {
	c.Lock()
	defer c.Unlock()
	c.loginInformation = nil
	c.currentPage = LoginPage
}
