import axios, { AxiosInstance } from "axios";
import {
  ChatRequest,
  ChatResponse,
  Message,
  UserStats,
} from "../../entities/message/types";
import { useNotificationStore } from "../store/notificationStore";

class ApiClient {
  private client: AxiosInstance;
  private notificationStore: ReturnType<typeof useNotificationStore.getState>;

  constructor() {
    this.client = axios.create({
      baseURL: "/api",
      timeout: 60000, // Увеличиваем до 60 секунд для сложных запросов
      headers: {
        "Content-Type": "application/json",
      },
    });

    this.notificationStore = useNotificationStore.getState();
    this.setupInterceptors();
  }

  private setupInterceptors() {
    // Request interceptor для добавления Telegram WebApp данных
    this.client.interceptors.request.use(
      (config) => {
        const tg = (window as any).Telegram?.WebApp;
        if (tg?.initData) {
          config.headers["X-Telegram-Init-Data"] = tg.initData;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    // Response interceptor для обработки ошибок
    this.client.interceptors.response.use(
      (response) => response,
      (error) => {
        const status = error.response?.status;
        const errorData = error.response?.data;

        // Определяем тип уведомления на основе статуса
        let notificationType: "error" | "warning" | "info" = "error";
        if (status === 429) {
          notificationType = "warning"; // Too Many Requests
        } else if (status >= 500) {
          notificationType = "error"; // Server errors
        } else if (status >= 400) {
          notificationType = "error"; // Client errors
        }

        // Извлекаем сообщение об ошибке
        let errorMessage = "Произошла ошибка";
        if (errorData?.error) {
          errorMessage = errorData.error;
        } else if (error.message) {
          errorMessage = error.message;
        }

        // Показываем уведомление
        this.notificationStore.addNotification(errorMessage, notificationType);

        if (status === 401) {
          // Обработка ошибки аутентификации
          console.error("Authentication error:", errorData);
        }

        return Promise.reject(error);
      }
    );
  }

  // Chat API
  async sendMessage(data: ChatRequest): Promise<ChatResponse> {
    const response = await this.client.post<ChatResponse>("/chat", data);
    return response.data;
  }

  async getChatHistory(): Promise<Message[]> {
    try {
      const response = await this.client.get<{
        messages: Message[];
        count: number;
      }>("/history");
      return response.data.messages || [];
    } catch (error) {
      console.error("Failed to get chat history:", error);
      return [];
    }
  }

  async getUserStats(): Promise<UserStats> {
    const response = await this.client.get<{
      daily_messages: number;
      daily_limit: number;
      remaining: number;
    }>("/stats");

    // Преобразуем формат данных API в формат фронтенда
    return {
      userID: 0, // Будет установлено из контекста Telegram
      messagesToday: response.data.daily_messages,
      dailyLimit: response.data.daily_limit,
      messagesRemaining: response.data.remaining,
    };
  }

  // Health check
  async healthCheck(): Promise<{ status: string; service: string }> {
    const response = await this.client.get("/health");
    return response.data;
  }
}

export const apiClient = new ApiClient();
