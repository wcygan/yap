Things you need to do...

High Priority
- Implement the backend flow for chatting
- Figure out how to use ingress for yap-api instead of port-forwarding

Low Priority

- Hash the passwords lol
- Work on the Chat CLI
  - Add three screens:
    - Login (DONE)
    - Home (ask AI to clone it similar to the login page)
    - Chat (ask AI to integrate the example from https://github.com/charmbracelet/bubbletea/blob/master/examples/chat/main.go with your program)
  - Add page transitions according to the UI spec
  - Add a way to create a new user -- API Call
  - Add a way to login -- API Call
    - Store username & access token in CLI state (model?)
  - Add a way to create a new chat -- API Call
  - Add a way to connect to a chat -- API Call
  - Add a way to send and receive messages (grpc streaming) -- API Calls
  - Add a way to logout
- Display the username in the chat window with a color
  - This can be done by hashing the username and using the first 6 characters of the hash as the color code
- Publish CLI binaries to GitHub / Add installation instructions
- Publish a Docker image of the CLI to Docker Hub


## How to implement chat room streaming

We can have the API server handle the chat room streaming. The client can connect to the API server using gRPC and send messages to the chat room. The API server can then broadcast the messages to all the clients connected to the chat room.

```protobuf
syntax = "proto3";

package chat;

service ChatService {
  rpc ChatStream (stream ChatPacket) returns (stream ChatPacket) {}
}

message ChatPacket {
  oneof packet_type {
    JoinChatRequest join_request = 1;
    ChatMessage chat_message = 2;
  }
}

message JoinChatRequest {
  string user_id = 1;
  string room_name = 2;
}

message ChatMessage {
  string user_id = 1;
  string message = 2;
  int64 timestamp = 3;
}
```

The client will initialize a gRPC streaming session with the API server and send a `JoinChatRequest` message to join the chat room. The API server will then add the client to the chat room and start sending messages to the client.

The client can send `ChatMessage` messages to the API server, which will then broadcast the message to all the clients in the chat room.

When the API receives a `ChatMessage` message from a client, it will forward the message to `chat-service`, which will persist the message into ScyllaDB and then publish the message on Apache Pulsar.

Each instance of the API will have a separate connection to the Pulsar cluster, and the messages will be broadcast to all the clients connected to the chat room.