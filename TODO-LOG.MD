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

import "google/protobuf/empty.proto";

package chat;

// A service for client streaming chat messages
service ClientStreamingService {
  // A client stream for sending and receiving chat messages
  rpc ChatStream (stream ChatPacket) returns (stream ChatPacket) {}
}

// A service for persisting chat messages
service MessagePersistenceService {
  // A unary rpc to submit chat messages for persistence
  rpc PersistChatMessage (MessageToPersist) returns (google.protobuf.Empty) {}
}

// User-facing packet that can be streamed over the wire
message ChatPacket {
  oneof packet_type {
    // A request to join a chat room
    JoinChatRequest join_request = 1;
    // A chat message
    ChatMessage chat_message = 2;
  }
}

// A request to join a chat room
message JoinChatRequest {
  // The user's unique identifier
  string user_id = 1;
  // The name of the chat room to join
  string room_name = 2;
}

// A chat message
message ChatMessage {
  // The user's unique identifier
  string user_id = 1;
  // The message content
  string message = 2;
  // The timestamp of the message
  int64 timestamp = 3;
}

// A message that will be persisted for a specific chat room
message MessageToPersist {
  // The name of the chat room to which the message belongs
  string room_name = 1;
  // The chat message to persist
  ChatMessage message = 2;
}
```

The client will initialize a gRPC streaming session with the API server and send a `JoinChatRequest` message to join the chat room. The API server will then add the client to the chat room and start sending messages to the client.

The client can send `ChatMessage` messages to the API server, which will then broadcast the message to all the clients in the chat room.

When the API receives a `ChatMessage` message from a client, it will forward the message to `chat-service`, which will persist the message into ScyllaDB and then publish the message on Apache Pulsar.

Each instance of the API will have a separate connection to the Pulsar cluster, and the messages will be broadcast to all the clients connected to the chat room.