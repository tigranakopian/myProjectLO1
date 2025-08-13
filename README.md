# myProjectLO1
myProject-lo1
# Task API

Простое REST API на Go для управления задачами с асинхронным логированием.

## 🔧 Сборка и запуск

```bash
go build -o task-api ./cmd/server
./task-api

## 📦 Примеры запросов (PowerShell)

### 🔨 Создание задачи

Invoke-RestMethod -Uri "http://localhost:8081/tasks" -Method POST -Body '{"id":"1","title":"Test","status":"pending"}' -ContentType "application/json"



📄 Получение задачи по ID
Invoke-RestMethod -Uri "http://localhost:8081/tasks/1" -Method GET

📋 Получение всех задач
Invoke-RestMethod -Uri "http://localhost:8081/tasks" -Method GET

🔍 Фильтрация по статусу
Invoke-RestMethod -Uri "http://localhost:8081/tasks?status=pending" -Method GET



