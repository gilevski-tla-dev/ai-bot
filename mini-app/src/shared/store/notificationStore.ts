import { create } from "zustand";

export interface Notification {
  id: string;
  message: string;
  type: "error" | "warning" | "info";
  timestamp: number;
}

interface NotificationState {
  notifications: Notification[];
  addNotification: (
    message: string,
    type?: "error" | "warning" | "info"
  ) => void;
  removeNotification: (id: string) => void;
  clearAllNotifications: () => void;
}

export const useNotificationStore = create<NotificationState>((set, get) => ({
  notifications: [],

  addNotification: (
    message: string,
    type: "error" | "warning" | "info" = "error"
  ) => {
    const id = Math.random().toString(36).substr(2, 9);
    const notification: Notification = {
      id,
      message,
      type,
      timestamp: Date.now(),
    };

    set((state) => {
      const newNotifications = [...state.notifications, notification];

      // Ограничиваем количество уведомлений до 3
      if (newNotifications.length > 3) {
        // Удаляем самое старое уведомление
        newNotifications.sort((a, b) => a.timestamp - b.timestamp);
        newNotifications.shift();
      }

      return { notifications: newNotifications };
    });

    // Автоматически удаляем уведомление через 3 секунды
    setTimeout(() => {
      get().removeNotification(id);
    }, 100000);
  },

  removeNotification: (id: string) => {
    set((state) => ({
      notifications: state.notifications.filter(
        (notification) => notification.id !== id
      ),
    }));
  },

  clearAllNotifications: () => {
    set({ notifications: [] });
  },
}));
