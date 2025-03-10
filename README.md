### **Описание проекта / Project Description**

Этот проект реализует **Telegram-бота**, который управляет подписками на уведомления. Бот позволяет пользователям подписываться и отписываться от уведомлений, а также получает уведомления через вебхуки от внешнего сервиса (например, AmoCRM) и рассылает их всем подписчикам.

Проект использует **Gin** для обработки HTTP-запросов и **SQLite** для хранения данных о подписчиках. Для развертывания используется **Docker Compose**, что упрощает настройку и запуск проекта.

---

This project implements a **Telegram bot** that manages notification subscriptions. The bot allows users to subscribe and unsubscribe from notifications, receives notifications via webhooks from an external service (e.g., AmoCRM), and distributes them to all subscribers.

The project uses **Gin** for handling HTTP requests and **SQLite** for storing subscriber data. **Docker Compose** is used for deployment, making the setup and execution easier.

---

### **Функции / Features**

- Подключение и настройка Telegram-бота / Connecting and setting up the Telegram bot.
- Обработка команд `/subscribe` и `/unsubscribe` от пользователей / Processing `/subscribe` and `/unsubscribe` commands from users.
- Получение и обработка вебхуков от внешнего сервиса / Receiving and processing webhooks from an external service.
- Хранение данных о подписчиках в базе данных **SQLite** / Storing subscriber data in an **SQLite** database.
- Рассылка уведомлений всем подписчикам / Sending notifications to all subscribers.

---

### **Технологии / Technologies**

- **Go (Gin)**
- **Telegram Bot API**
- **SQLite**
- **Docker Compose**
- **gin-gonic/gin**
- **go-telegram-bot-api**

---

### **Как запустить проект / How to Run the Project**

#### **1. Установка зависимостей / Install Dependencies**

Перед запуском убедитесь, что у вас установлен **Docker** и **Docker Compose**.

Before running, ensure that **Docker** and **Docker Compose** are installed.

#### **2. Запуск с Docker Compose / Run with Docker Compose**

1. Клонируйте репозиторий / Clone the repository:
   ```bash
   git clone https://github.com/EmelinDanila/telegram-bot-crm.git
   ```

2. Перейдите в директорию проекта / Navigate to the project directory:
   ```bash
   cd telegram-bot-crm
   ```

3. Создайте файл `.env` в корне проекта и добавьте в него ваши данные / Create a `.env` file in the project root and add your credentials:
   ```
   TELEGRAM_BOT_TOKEN=your-telegram-bot-token
   ```

4. Запустите проект с помощью Docker Compose / Start the project using Docker Compose:
   ```bash
   docker-compose up --build
   ```

5. После успешного запуска, проект будет доступен по адресу: `http://localhost:8080`.

   After successful startup, the project will be available at: `http://localhost:8080`.

---

### **Эндпоинты / Endpoints**

1. **`/send`** - Эмуляция вебхука от AmoCRM / Webhook simulation from AmoCRM.
   - Принимает JSON-данные и рассылает их всем подписчикам / Accepts JSON data and sends it to all subscribers.

2. **`/telegram`** - Обработка команд Telegram-бота / Processing Telegram bot commands.
   - Подписка на уведомления: `/subscribe` / Subscribe to notifications: `/subscribe`
   - Отписка от уведомлений: `/unsubscribe` / Unsubscribe from notifications: `/unsubscribe`

---

### **Как настроить ngrok для локальной разработки / How to Setup ngrok for Local Development**

1. Запустите локальный сервер / Start the local server:
   ```bash
   docker-compose up
   ```

2. Запустите **ngrok** для проксирования запросов в ваш локальный сервер / Start **ngrok** to proxy requests to your local server:
   ```bash
   ngrok http 8080
   ```

3. Получите URL, предоставленный **ngrok**, и используйте его для настройки вебхуков в Telegram или AmoCRM / Get the URL provided by **ngrok** and use it to configure webhooks in Telegram or AmoCRM.

---

Если у вас возникнут вопросы или потребуется помощь, не стесняйтесь обращаться! / If you have any questions or need help, feel free to ask!