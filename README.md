# Real-time Chat Application

## Architecture
- Microservices architecture with services for:
  - Authentication (auth service)
  - Messaging (chat service) 
  - Group management (group service)
- gRPC for inter-service communication
- buf for protobuf code generation
- PostgreSQL database for persistence
- Kubernetes for deployment and scaling

## Authentication
- JWT tokens for authentication
- Auth service responsible for issuing and validating tokens
- Tokens passed in gRPC metadata for each request
- Refresh tokens used to get new access tokens

## Messaging
- Real-time messaging implemented using gRPC streaming
- Client opens bidirectional stream with server
- Messages encrypted with symmetric key, key encrypted with recipient's public key
- Messages stored encrypted in the database

## Group Chats
- Group service manages group membership 
- Chat service handles group messaging
- Messages sent to group chat are fanned out to each member's message stream

## Deployment
- Each service deployed as a Kubernetes Deployment
- Kubernetes Services used for internal service discovery
- Ingress used to expose services externally
- Helm charts used to manage deployment configuration
- Horizontal Pod Autoscalers for scaling based on CPU usage

## Testing
- Unit tests for each service
- Integration tests for key workflows (user registration, sending messages, group chats)
- End-to-end tests that test the full flow from client to server
- CI/CD pipeline to run tests on each commit

## Client
- Go client that communicates with backend over gRPC
- Handles user login/registration flows
- Establishes gRPC streams for real-time messaging
- Encrypts/decrypts messages client-side

## Next Steps
1. Design gRPC APIs for each service
2. Implement auth service and client login/registration
3. Implement messaging service and client-side messaging
4. Implement group service and group messaging
5. Set up Kubernetes cluster and deploy services 
6. Implement end-to-end encryption
7. Optimize performance and add additional features