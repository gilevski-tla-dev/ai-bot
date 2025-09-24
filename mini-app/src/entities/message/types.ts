export interface Message {
  id: number;
  user_id: number;
  role: "user" | "assistant";
  content: string;
  created_at: string;
}

export interface ChatRequest {
  message: string;
}

export interface ChatResponse {
  message: string;
  timestamp: string;
}

export interface UserStats {
  userID: number;
  messagesToday: number;
  dailyLimit: number;
  messagesRemaining: number;
}
