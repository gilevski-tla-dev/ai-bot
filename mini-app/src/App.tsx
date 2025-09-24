import { Theme, TextArea, Button, Flex, Box, Text } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import { useState } from "react";

function App() {
  const [inputValue, setInputValue] = useState("");

  const handleSubmit = () => {
    setInputValue("");
  };

  return (
    <Theme appearance="dark" accentColor="sky" radius="full">
      <Box
        style={{
          height: "100vh",
          display: "flex",
          flexDirection: "column",
          backgroundColor: "var(--color-background)",
          padding: "1rem",
        }}
      >
        {/* Main content area */}
        <Box
          style={{
            flex: 1,
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Text size="6" weight="bold" mb="4" style={{ textAlign: "center" }}>
            Deepseek AI
          </Text>
          <Text
            size="3"
            color="gray"
            style={{ textAlign: "center", maxWidth: "300px" }}
          >
            Начинайте общаться и уменьшайте количество рутины
          </Text>
        </Box>

        {/* Fixed input at bottom */}
        <Box
          style={{
            bottom: 0,
            padding: "1rem",
            backgroundColor: "var(--color-background)",
            borderTop: "1px solid var(--gray-6)",
            backdropFilter: "blur(10px)",
            zIndex: 1000,
          }}
        >
          <Flex direction="column" gap="2">
            <TextArea
              value={inputValue}
              onChange={(e) => setInputValue(e.target.value)}
              placeholder="Введите сообщение..."
              style={{
                flex: 1,
                minHeight: "60px",
                maxHeight: "120px",
                resize: "vertical",
              }}
              size="3"
              rows={5}
              resize="vertical"
            />
            <Button
              onClick={handleSubmit}
              disabled={!inputValue.trim()}
              size="3"
              style={{ minWidth: "60px" }}
            >
              Отправить
            </Button>
          </Flex>
        </Box>
      </Box>
    </Theme>
  );
}

export default App;
