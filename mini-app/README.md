# 🤖 Telegram Mini App - DeepSeek AI Chat

Современное мини-приложение для Telegram с интеграцией DeepSeek AI через OpenRouter API.

## 🏗️ Архитектура

Проект построен с использованием **Feature-Sliced Design (FSD)** архитектуры:

```
src/
├── app/                 # Инициализация приложения
│   ├── providers/       # React Query Provider
│   └── App.tsx          # Главный компонент
├── pages/              # Страницы
│   └── chat/           # Страница чата
├── widgets/            # Виджеты
│   └── chat/           # Виджет чата
├── features/           # Фичи
│   ├── chat/           # Логика чата
│   └── auth/           # Аутентификация
├── entities/           # Сущности
│   ├── message/        # Модель сообщения
│   └── user/           # Модель пользователя
└── shared/             # Общие компоненты
    ├── api/            # API клиент
    ├── ui/             # UI компоненты
    └── lib/            # Утилиты
```

## 🛠️ Технологии

- **React 19** - UI библиотека
- **TypeScript** - типизация
- **Vite** - сборщик
- **Radix UI** - компоненты
- **TanStack Query** - управление состоянием сервера
- **Axios** - HTTP клиент
- **Feature-Sliced Design** - архитектура

## 🚀 Функциональность

### ✅ Реализовано:

- 💬 **Чат с ИИ** - отправка сообщений и получение ответов
- 📝 **История сообщений** - отображение предыдущих диалогов
- 📊 **Статистика пользователя** - лимиты и использование
- 🔐 **Аутентификация** - через Telegram WebApp API
- ⚡ **Оптимистичные обновления** - мгновенный отклик UI
- 🔄 **Автоматическое обновление** - синхронизация с сервером
- 📱 **Адаптивный дизайн** - для мобильных устройств

### 🎯 API интеграция:

- `POST /api/chat` - отправка сообщения
- `GET /api/history` - получение истории
- `GET /api/stats` - статистика пользователя
- `GET /health` - проверка здоровья

## 📦 Установка и запуск

```bash
# Установка зависимостей
npm install

# Разработка
npm run dev

# Сборка
npm run build

# Линтинг
npm run lint
```

## 🔧 Конфигурация

### Переменные окружения:

```bash
VITE_API_BASE_URL=/api  # Базовый URL API
```

### TanStack Query настройки:

- **Retry**: 2 попытки для queries, 1 для mutations
- **Stale Time**: 5 минут для queries
- **Refetch**: отключен при фокусе окна

## 📱 Telegram WebApp интеграция

Приложение автоматически:

- Получает данные пользователя из Telegram
- Отправляет `initData` для аутентификации
- Адаптируется под тему Telegram
- Использует нативные возможности WebApp

## 🎨 UI/UX

- **Темная тема** - соответствует Telegram
- **Адаптивный дизайн** - для всех устройств
- **Плавные анимации** - для лучшего UX
- **Индикаторы загрузки** - обратная связь
- **Обработка ошибок** - понятные сообщения

## 🔄 Состояние приложения

### TanStack Query кэш:

```typescript
// Ключи запросов
chatKeys = {
  all: ["chat"],
  history: () => [...chatKeys.all, "history"],
  stats: () => [...chatKeys.all, "stats"],
};
```

### Автоматические обновления:

- После отправки сообщения обновляется история и статистика
- Кэш инвалидируется для актуальных данных
- Оптимистичные обновления для мгновенного отклика

## 🚀 Развертывание

Приложение автоматически собирается в Docker контейнере:

```bash
# Сборка образа
docker build -t mini-app .

# Запуск
docker run -p 80:80 mini-app
```

## 📊 Производительность

- **Bundle size**: ~337KB (gzipped: ~109KB)
- **First load**: оптимизирован для быстрой загрузки
- **Caching**: агрессивное кэширование запросов
- **Lazy loading**: компоненты загружаются по требованию

## 🔒 Безопасность

- **Аутентификация**: через Telegram WebApp
- **Валидация**: всех входящих данных
- **CORS**: настроен для Telegram доменов
- **Rate limiting**: на уровне API

## 🧪 Тестирование

```bash
# Запуск тестов
npm test

# Покрытие
npm run test:coverage
```

## 📈 Мониторинг

- **React Query DevTools** - в режиме разработки
- **Error boundaries** - обработка ошибок
- **Logging** - детальные логи для отладки

## 🤝 Разработка

### Структура компонентов:

```typescript
// Entity
export interface Message {
  id: number;
  user_id: number;
  role: "user" | "assistant";
  content: string;
  created_at: string;
}

// Feature hook
export const useSendMessage = () => {
  return useMutation({
    mutationFn: (data: ChatRequest) => apiClient.sendMessage(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: chatKeys.history() });
    },
  });
};
```

### Добавление новых фич:

1. Создать папку в `features/`
2. Добавить API хуки в `api/`
3. Создать UI компоненты в `ui/`
4. Экспортировать через `index.ts`

## 📚 Документация

- [TanStack Query](https://tanstack.com/query/latest)
- [Radix UI](https://www.radix-ui.com/)
- [Feature-Sliced Design](https://feature-sliced.design/)
- [Telegram WebApp](https://core.telegram.org/bots/webapps)
