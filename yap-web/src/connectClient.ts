import { createConnectTransport } from "@bufbuild/connect-web";
import { createPromiseClient } from "@bufbuild/connect";
import { ChatRoomService } from "./generated/chat/v1/chat_connectweb"

const transport = createConnectTransport({
  baseUrl: "http://localhost:50050",
});

const client = createPromiseClient(ChatRoomService, transport);

export default client;