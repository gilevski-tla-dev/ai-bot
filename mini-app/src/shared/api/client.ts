import axios, { AxiosInstance } from "axios";
import {
  ChatRequest,
  ChatResponse,
  Message,
  UserStats,
} from "../../entities/message/types";

class ApiClient {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: "/api",
      timeout: 30000,
      headers: {
        "Content-Type": "application/json",
      },
    });

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
        if (error.response?.status === 401) {
          // Обработка ошибки аутентификации
          console.error("Authentication error:", error.response.data);
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
    const response = await this.client.get<Message[]>("/history");
    return response.data;
  }

  async getUserStats(): Promise<UserStats> {
    const response = await this.client.get<UserStats>("/stats");
    return response.data;
  }

  // Health check
  async healthCheck(): Promise<{ status: string; service: string }> {
    const response = await this.client.get("/health");
    return response.data;
  }
}

export const apiClient = new ApiClient();
