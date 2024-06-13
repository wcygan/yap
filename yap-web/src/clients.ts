import {createConnectTransport} from "@bufbuild/connect-web";
import {createPromiseClient} from "@bufbuild/connect";
import {ChatRoomService, MessagingService} from "./generated/chat/v1/chat_connectweb"
import {AuthService} from "./generated/auth/v1/auth_connectweb"

const transport = createConnectTransport({
    baseUrl: "http://localhost:50050",
});

const messagingClient = createPromiseClient(MessagingService, transport);
const chatRoomClient = createPromiseClient(ChatRoomService, transport);
const authClient = createPromiseClient(AuthService, transport);

export { messagingClient, chatRoomClient, authClient };