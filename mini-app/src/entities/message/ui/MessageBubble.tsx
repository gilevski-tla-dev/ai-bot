import { Box, Text } from "@radix-ui/themes";
import { Message } from "../types";

interface MessageBubbleProps {
  message: Message;
}

export const MessageBubble = ({ message }: MessageBubbleProps) => {
  const isUser = message.role === "user";

  return (
    <Box
      style={{
        alignSelf: isUser ? "flex-end" : "flex-start",
        maxWidth: "80%",
        marginBottom: "0.5rem",
      }}
    >
      <Box
        style={{
          padding: "0.75rem 1rem",
          borderRadius: "12px",
          backgroundColor: isUser ? "var(--accent-9)" : "var(--gray-3)",
          color: isUser ? "var(--accent-contrast)" : "var(--gray-12)",
        }}
      >
        <Text size="2">{message.content}</Text>
      </Box>
    </Box>
  );
};
