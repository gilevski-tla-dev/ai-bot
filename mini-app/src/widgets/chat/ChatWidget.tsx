import { Box, Text } from "@radix-ui/themes";
import { useState, useMemo } from "react";
import {
  useSendMessage,
  useChatHistory,
  useUserStats,
} from "../../features/chat/api/useChat";
import { MessageList } from "../../entities/message/ui/MessageList";
import { ChatInput } from "../../features/chat/ui/ChatInput";
import { UserStats } from "../../features/chat/ui/UserStats";
import { Message } from "../../entities/message/types";

export const ChatWidget = () => {
  const { data: messages = [] } = useChatHistory();
  const { data: stats } = useUserStats();
  const sendMessageMutation = useSendMessage();

  // Локальное состояние для временных сообщений
  const [tempMessages, setTempMessages] = useState<Message[]>([]);

  // Объединяем сообщения из сервера с временными
  const allMessages = useMemo(() => {
    return [...messages, ...tempMessages];
  }, [messages, tempMessages]);

  const handleSendMessage = async (message: string) => {
    // Создаем временное сообщение пользователя
    const tempUserMessage: Message = {
      id: Date.now(), // Временный ID
      user_id: 0, // Будет обновлено с сервера
      role: "user",
      content: message,
      created_at: new Date().toISOString(),
    };

    // Сразу добавляем сообщение пользователя в локальное состояние
    setTempMessages([tempUserMessage]);

    try {
      await sendMessageMutation.mutateAsync({ message });

      // Очищаем временные сообщения после успешной отправки
      // Данные обновятся автоматически через invalidation в useSendMessage
      setTempMessages([]);
    } catch (error) {
      console.error("Failed to send message:", error);
      // Убираем временное сообщение при ошибке
      setTempMessages([]);

      // Можно добавить уведомление пользователю об ошибке
      // Например, через toast или другой UI компонент
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
        <MessageList
          messages={allMessages}
          isLoading={isLoading && tempMessages.length > 0}
        />
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
