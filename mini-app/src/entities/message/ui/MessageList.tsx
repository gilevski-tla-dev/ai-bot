import { Box, ScrollArea, Text } from "@radix-ui/themes";
import { useEffect, useRef } from "react";
import { Message } from "../types";
import { MessageBubble } from "./MessageBubble";

interface MessageListProps {
  messages: Message[];
  isLoading?: boolean;
}

export const MessageList = ({ messages, isLoading }: MessageListProps) => {
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    if (messages && messages.length > 0) {
      scrollToBottom();
    }
  }, [messages]);

  if (!messages || messages.length === 0) {
    return (
      <Box
        style={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          height: "100%",
          textAlign: "center",
          padding: "2rem",
        }}
      >
        <Text size="4" weight="bold" mb="2">
          Добро пожаловать в DeepSeek AI!
        </Text>
        <Text size="3" color="gray" style={{ maxWidth: "300px" }}>
          Начинайте общаться с ИИ. Задавайте вопросы, получайте ответы.
        </Text>
      </Box>
    );
  }

  return (
    <ScrollArea style={{ height: "100%" }}>
      <Box style={{ padding: "1rem", minHeight: "100%" }}>
        {messages?.map((message) => (
          <MessageBubble key={message.id} message={message} />
        ))}
        {isLoading && (
          <Box
            style={{
              alignSelf: "flex-start",
              maxWidth: "80%",
            }}
          >
            <Box
              style={{
                padding: "0.75rem 1rem",
                borderRadius: "12px",
                backgroundColor: "var(--gray-3)",
              }}
            >
              <Text size="2" color="gray">
                Думаю...
              </Text>
            </Box>
          </Box>
        )}
        <div ref={messagesEndRef} />
      </Box>
    </ScrollArea>
  );
};
