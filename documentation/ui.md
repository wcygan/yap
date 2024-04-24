We will build the TUI with https://github.com/charmbracelet/bubbletea. The main idea is that
it is essentially an MVC application with a view, model, and controller according to the following design:

MODEL
1. Current Page
    1. Login
    2. Home
    3. Chat
2. Login Status
    1. LOGGED OUT
    2. LOGGED IN
        1. Username
        2. Access Token
        3. Refresh Token
3. Current Chat Room
    1. Room Name
    2. Message
        1. User
        2. Text
    3. grpc connection

VIEW
1. When logged out, show the login page
2. When logged in, show the home page
3. When in a chat room, show the chat room page
4. When logging in or creating an account, show error messages for failure
5. When joining a non-existent chat room, show error messages for failure

CONTROLLER
1. Logging In
    1. Spinner
    2. Transition to homepage
2. Join group chat
    1. Spinner
    2. Transition to chat room
4. Logging out
    1. Transition to login
5. "Leave Chat"
    1. Transition to homepage

PAGES:

1. Login
    1. Username
    2. Password
    3. Login Button
    4. Create Account Button
2. Home
   1. "Create Chat" Button
   2. "Create Chat" Text Field
   3. "Join Chat" Button
   4. "Join Chat" Text Field
   5. "Logout" Button
3. Chat
   1. Chat Room Name
   2. Chat Room Messages
   3. Chat Room Message Input
   4. Chat Room Send Button
   5. Chat Room Leave Button