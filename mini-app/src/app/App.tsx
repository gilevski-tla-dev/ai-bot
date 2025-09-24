import { Theme } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import { QueryProvider } from "./providers/QueryProvider";
import { ChatPage } from "../pages/chat/ChatPage";

export const App = () => {
  return (
    <QueryProvider>
      <Theme appearance="dark" accentColor="sky" radius="full">
        <ChatPage />
      </Theme>
    </QueryProvider>
  );
};
