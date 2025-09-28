import { Theme } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import { QueryProvider } from "./providers/QueryProvider";
import { ChatPage } from "../pages/chat/ChatPage";
import { NotificationContainer } from "../shared/ui/NotificationContainer";

export const App = () => {
  return (
    <QueryProvider>
      <Theme appearance="dark" accentColor="sky" radius="full">
        <NotificationContainer />
        <ChatPage />
      </Theme>
    </QueryProvider>
  );
};
