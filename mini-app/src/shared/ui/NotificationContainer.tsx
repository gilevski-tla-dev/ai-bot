import { Box } from "@radix-ui/themes";
import { useNotificationStore } from "../store/notificationStore";
import { NotificationChip } from "./NotificationChip";

export const NotificationContainer = () => {
  const { notifications, removeNotification } = useNotificationStore();

  if (notifications.length === 0) {
    return null;
  }

  return (
    <Box
      style={{
        position: "fixed",
        top: "0",
        left: "0",
        right: "0",
        zIndex: 9999,
        display: "flex",
        flexDirection: "column",
        gap: "8px",
        width: "100%",
        padding: "16px",
        pointerEvents: "none",
      }}
    >
      {notifications.map((notification) => (
        <Box
          key={notification.id}
          style={{
            pointerEvents: "auto",
            animation: "slideInFromTop 0.3s ease-out",
          }}
        >
          <NotificationChip
            notification={notification}
            onClose={() => removeNotification(notification.id)}
          />
        </Box>
      ))}

      <style>{`
        @keyframes slideInFromTop {
          from {
            transform: translateY(-100%);
            opacity: 0;
          }
          to {
            transform: translateY(0);
            opacity: 1;
          }
        }
      `}</style>
    </Box>
  );
};
