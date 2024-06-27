import {createConnectTransport, createGrpcWebTransport} from "@bufbuild/connect-web";
import {createPromiseClient} from "@bufbuild/connect";
import {ChatRoomService, MessagingService} from "./generated/chat/v1/chat_connectweb";
import {AuthService} from "./generated/auth/v1/auth_connectweb";

const grpcWebTransport = createGrpcWebTransport({
    baseUrl: "http://127.0.0.1:8080",
});

const messagingClient = createPromiseClient(MessagingService, grpcWebTransport);
const chatRoomClient = createPromiseClient(ChatRoomService, grpcWebTransport);
const authClient = createPromiseClient(AuthService, grpcWebTransport);

export {messagingClient, chatRoomClient, authClient};