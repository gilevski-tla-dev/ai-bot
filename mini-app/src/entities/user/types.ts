export interface User {
  id: number;
  user_id: number;
  username?: string;
  first_name?: string;
  last_name?: string;
  created_at: string;
}

export interface TelegramUser {
  id: number;
  first_name: string;
  last_name?: string;
  username?: string;
  language_code?: string;
}
