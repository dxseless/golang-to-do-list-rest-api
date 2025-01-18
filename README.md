# To-Do List API

Это простой REST API для управления списком задач, написанный на Go. Проект использует базу данных SQLite через драйвер `modernc.org/sqlite`, который не требует `cgo`.

## Запуск проекта

1. Убедитесь, что у вас установлен Go.
2. Клонируйте репозиторий.
3. Перейдите в директорию проекта и выполните:

```bash
go mod tidy
```

```bash
go run main.go
```

Тестирование API

Примеры запросов:

Получить все задачи:
curl -X GET http://localhost:8080/todos

Создать новую задачу:
curl -X POST -H "Content-Type: application/json" -d '{"task":"купить молоко"}' http://localhost:8080/todos

Получить задачу по ID:
curl -X GET http://localhost:8080/todos/1

Обновить задачу:
curl -X PUT -H "Content-Type: application/json" -d '{"task":"купить хлеб"}' http://localhost:8080/todos/1

Удалить задачу:
curl -X DELETE http://localhost:8080/todos/1

Использование API

GET /todos - Получить все задачи.

GET /todos/{id} - Получить задачу по ID.

POST /todos - Создать новую задачу.

PUT /todos/{id} - Обновить задачу.

DELETE /todos/{id} - Удалить задачу.
