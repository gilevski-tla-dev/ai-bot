import { Cross2Icon } from "@radix-ui/react-icons";
import { Box, Flex, Text } from "@radix-ui/themes";
import { useEffect, useState } from "react";
import { Notification } from "../store/notificationStore";

interface NotificationChipProps {
  notification: Notification;
  onClose: () => void;
}

export const NotificationChip = ({
  notification,
  onClose,
}: NotificationChipProps) => {
  const [isVisible, setIsVisible] = useState(false);
  const [isExiting, setIsExiting] = useState(false);

  useEffect(() => {
    // Анимация появления
    const timer = setTimeout(() => setIsVisible(true), 10);
    return () => clearTimeout(timer);
  }, []);

  const handleClose = () => {
    setIsExiting(true);
    setTimeout(onClose, 200); // Даем время на анимацию исчезновения
  };

  const getNotificationStyles = () => {
    switch (notification.type) {
      case "error":
        return {
          backgroundColor: "var(--red-9)",
          borderColor: "var(--red-8)",
          color: "var(--red-1)",
        };
      case "warning":
        return {
          backgroundColor: "var(--orange-9)",
          borderColor: "var(--orange-8)",
          color: "var(--orange-1)",
        };
      case "info":
        return {
          backgroundColor: "var(--blue-9)",
          borderColor: "var(--blue-8)",
          color: "var(--blue-1)",
        };
      default:
        return {
          backgroundColor: "var(--red-9)",
          borderColor: "var(--red-8)",
          color: "var(--red-1)",
        };
    }
  };

  const styles = getNotificationStyles();

  return (
    <Box
      style={{
        ...styles,
        padding: "12px 16px",
        borderRadius: "var(--radius-3)",
        border: "1px solid",
        borderColor: styles.borderColor,
        fontSize: "var(--font-size-2)",
        fontWeight: "var(--font-weight-medium)",
        boxShadow: "var(--shadow-3)",
        transform:
          isVisible && !isExiting ? "translateY(0)" : "translateY(-100%)",
        opacity: isVisible && !isExiting ? 1 : 0,
        transition: "all 0.2s ease-in-out",
        cursor: "pointer",
        userSelect: "none",
        width: "100%",
        wordBreak: "break-word",
      }}
      onClick={handleClose}
    >
      <Flex align="center" justify="between" gap="2">
        <Text style={{ color: styles.color, flex: 1 }}>
          {notification.message}
        </Text>
        <Box
          style={{
            color: styles.color,
            cursor: "pointer",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            width: "16px",
            height: "16px",
            borderRadius: "50%",
            backgroundColor: "rgba(255, 255, 255, 0.2)",
            transition: "background-color 0.2s ease",
          }}
          onMouseEnter={(e) => {
            e.currentTarget.style.backgroundColor = "rgba(255, 255, 255, 0.3)";
          }}
          onMouseLeave={(e) => {
            e.currentTarget.style.backgroundColor = "rgba(255, 255, 255, 0.2)";
          }}
        >
          <Cross2Icon width="10" height="10" />
        </Box>
      </Flex>
    </Box>
  );
};
