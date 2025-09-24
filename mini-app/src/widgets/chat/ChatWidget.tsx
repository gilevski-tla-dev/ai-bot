import { Box, Text } from "@radix-ui/themes";
import {
  useSendMessage,
  useChatHistory,
  useUserStats,
} from "../../features/chat/api/useChat";
import { MessageList } from "../../entities/message/ui/MessageList";
import { ChatInput } from "../../features/chat/ui/ChatInput";
import { UserStats } from "../../features/chat/ui/UserStats";

export const ChatWidget = () => {
  const { data: messages = [] } = useChatHistory();
  const { data: stats } = useUserStats();
  const sendMessageMutation = useSendMessage();

  const handleSendMessage = async (message: string) => {
    try {
      await sendMessageMutation.mutateAsync({ message });

      // Добавляем сообщение пользователя и ответ ассистента в локальное состояние
      // Это будет обновлено автоматически через invalidation в useSendMessage
    } catch (error) {
      console.error("Failed to send message:", error);
    }
  };

  const isLoading = sendMessageMutation.isPending;
  const isDisabled = stats?.messagesRemaining === 0;

  return (
    <Box
      style={{
        height: "100vh",
        display: "flex",
        flexDirection: "column",
        backgroundColor: "var(--color-background)",
      }}
    >
      {/* Header */}
      <Box
        style={{
          padding: "1rem",
          borderBottom: "1px solid var(--gray-6)",
          backgroundColor: "var(--color-background)",
        }}
      >
        <Text size="5" weight="bold" style={{ textAlign: "center" }}>
          DeepSeek AI Chat
        </Text>
      </Box>

      {/* User Stats */}
      {stats && <UserStats stats={stats} />}

      {/* Messages area */}
      <Box style={{ flex: 1, overflow: "hidden" }}>
        <MessageList messages={messages} isLoading={isLoading} />
      </Box>

      {/* Input area */}
      <Box
        style={{
          padding: "1rem",
          backgroundColor: "var(--color-background)",
          borderTop: "1px solid var(--gray-6)",
        }}
      >
        <ChatInput
          onSendMessage={handleSendMessage}
          isLoading={isLoading}
          disabled={isDisabled}
        />
      </Box>
    </Box>
  );
};
