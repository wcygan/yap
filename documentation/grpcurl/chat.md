# gRPC references for the chat endpoints

// TODO: update these with the new endpoints

## Persist Chat Message

```
grpcurl -proto proto/chat/v1/chat.proto -d '{"channel_id": "123e4567-e89b-12d3-a456-426614174000", "user_id": "f47ac10b-58cc-4372-a567-0e02b2c3d479", "message": "Hello, world!", "timestamp": 1625097600 }' -plaintext localhost:50052 chat.MessagePersistenceService/PersistChatMessage
```
