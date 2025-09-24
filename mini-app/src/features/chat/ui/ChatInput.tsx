import { Button, Flex, TextArea } from "@radix-ui/themes";
import { useState } from "react";

interface ChatInputProps {
  onSendMessage: (message: string) => void;
  isLoading: boolean;
  disabled?: boolean;
}

export const ChatInput = ({
  onSendMessage,
  isLoading,
  disabled,
}: ChatInputProps) => {
  const [inputValue, setInputValue] = useState("");

  const handleSubmit = () => {
    if (!inputValue.trim() || isLoading || disabled) return;

    onSendMessage(inputValue.trim());
    setInputValue("");
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSubmit();
    }
  };

  return (
    <Flex direction="column" gap="2">
      <TextArea
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        onKeyPress={handleKeyPress}
        placeholder="Введите сообщение..."
        disabled={isLoading || disabled}
        style={{
          flex: 1,
          minHeight: "60px",
          maxHeight: "120px",
          resize: "vertical",
        }}
        size="3"
        rows={3}
        resize="vertical"
      />
      <Button
        onClick={handleSubmit}
        disabled={!inputValue.trim() || isLoading || disabled}
        size="3"
        style={{ minWidth: "100px" }}
      >
        {isLoading ? "Отправка..." : "Отправить"}
      </Button>
    </Flex>
  );
};
