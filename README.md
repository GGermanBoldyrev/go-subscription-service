# Subscription Service API

REST-сервис для агрегации данных об онлайн-подписках пользователей  
Тестовое задание для Junior Golang Developer (Effective Mobile)

---

## Возможности

- CRUD-операции для подписок (создать, получить, обновить, удалить, получить список)
- Подсчет суммарной стоимости подписок за выбранный период с фильтрами по пользователю и сервису
- Swagger-документация (OpenAPI 3)
- Логгирование всех HTTP-запросов (middleware)
- Миграции схемы БД при старте
- Настройка через `.env`
- Docker Compose для запуска всего стека
- Swagger документация генерируется при сборке проекта
---

## Быстрый старт

1. **Клонируй репозиторий:**
    ```bash
    git clone https://github.com/yourusername/go-subscription-service.git
    cd go-subscription-service
    ```

2. **Создай файл `.env` в корне:**
    ```env
    DB_URL=postgres://user:password@db:5432/subscription?sslmode=disable
    PORT=8080
    ```

3. **Запусти сервис и БД через Docker Compose:**
    ```bash
    docker compose -f docker/docker-compose.yml up --build
    ```

---

## Документация API (Swagger)

После запуска открой:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---