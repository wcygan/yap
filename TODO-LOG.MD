--- 06/11/24

- need to send the user_id back to clients so they can include it in messages
- need to send "user joined" event to clients and include their user_id to color their messages

--- New ideas from 06/10/24:

- figure out how to solve the "in-memory broadcast" problem (single producer, multiple consumers) via https://dev.to/karanpratapsingh/embedding-nats-in-go-19o. see https://www.reddit.com/r/golang/comments/14r5q62/in_memory_message_broker_any_recommendations/ or https://watermill.io/ or https://dev.to/karanpratapsingh/embedding-nats-in-go-19o or https://www.reddit.com/r/golang/comments/1bk6n74/single_producer_continuously_notify_multiple/
- maybe just use a single-direction server streaming RPC to receive all chat messages. see https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc. this will take a join server request and return a stream of messages.
- then use a separate server non-streaming RPC to send every new chat message
- join channel request from client

Rationale:

By solving the "in-memory broadcast" problem, it allows us to efficiently send chat messages to all clients connected to a chat room.

The logic looks like this:
1. client sends a message to yap-api, which then forwards it to chat-service
2. chat-service persists the message into ScyllaDB, and then publishes it to Pulsar
3. each instance of yap-api receives the message from Pulsar and then broadcasts it internally (via nats) to all clients connected to the chat room on the particular yap-api instance

This design allows horizontal scaling of yap-api and makes it such that it doesn't matter which yap-api instance the client is connected to; they will still receive all chat messages.

--- Older stuff:

- place client into chat room
- SQL to create / fetch channels
- hook up endpoints 
- send chat messages from client
- forward chat messages from yap-api to chat-service
- add pulsar & produce chat messages after persisting them
- consume the chat messages and forward them to all of the proper streaming connections (pub sub using channels in go https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/)
- send chat messages to client
- display chat messages in chat window

---

Things you need to do...

High Priority
- Use https://opensource.docs.scylladb.com/stable/cql/cqlsh.html to interact with ScyllaDB in k8s
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
  string channel_name = 2;
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
  string channel_name = 1;
  // The chat message to persist
  ChatMessage message = 2;
}
```

The client will initialize a gRPC streaming session with the API server and send a `JoinChatRequest` message to join the chat room. The API server will then add the client to the chat room and start sending messages to the client.

The client can send `ChatMessage` messages to the API server, which will then broadcast the message to all the clients in the chat room.

When the API receives a `ChatMessage` message from a client, it will forward the message to `chat-service`, which will persist the message into ScyllaDB and then publish the message on Apache Pulsar.

Each instance of the API will have a separate connection to the Pulsar cluster, and the messages will be broadcast to all the clients connected to the chat room.