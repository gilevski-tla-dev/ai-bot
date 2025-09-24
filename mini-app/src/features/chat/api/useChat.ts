import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "../../../shared/api/client";
import { ChatRequest } from "../../../entities/message/types";

// Query keys
export const chatKeys = {
  all: ["chat"] as const,
  history: () => [...chatKeys.all, "history"] as const,
  stats: () => [...chatKeys.all, "stats"] as const,
};

// Хук для отправки сообщения
export const useSendMessage = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: ChatRequest) => apiClient.sendMessage(data),
    onSuccess: () => {
      // Инвалидируем кэш истории и статистики
      queryClient.invalidateQueries({ queryKey: chatKeys.history() });
      queryClient.invalidateQueries({ queryKey: chatKeys.stats() });
    },
    onError: (error) => {
      console.error("Failed to send message:", error);
    },
  });
};

// Хук для получения истории сообщений
export const useChatHistory = () => {
  return useQuery({
    queryKey: chatKeys.history(),
    queryFn: () => apiClient.getChatHistory(),
    staleTime: 1000 * 60 * 5, // 5 минут
    refetchOnWindowFocus: false,
  });
};

// Хук для получения статистики пользователя
export const useUserStats = () => {
  return useQuery({
    queryKey: chatKeys.stats(),
    queryFn: () => apiClient.getUserStats(),
    staleTime: 1000 * 60 * 2, // 2 минуты
    refetchOnWindowFocus: false,
  });
};

// Хук для проверки здоровья API
export const useHealthCheck = () => {
  return useQuery({
    queryKey: ["health"],
    queryFn: () => apiClient.healthCheck(),
    staleTime: 1000 * 60 * 5, // 5 минут
    refetchInterval: 1000 * 60 * 5, // Проверяем каждые 5 минут
  });
};
