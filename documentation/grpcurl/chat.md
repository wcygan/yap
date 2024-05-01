# gRPC references for the chat endpoints

## Persist Chat Message

```
grpcurl -proto proto/chat/v1/chat.proto -d '{"room_name": "general", "message": {"user_id": "user123", "message": "Hello!", "timestamp": 1686012345}}' -plaintext localhost:50052 chat.MessagePersistenceService/PersistChatMessage
```
