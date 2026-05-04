### REST API на Go

### Текущая реализация (что уже сделано)

- Код разделён на **слои** (handler, service, repository).
- Реализованы **все CRUD-операции**.
- Использована только **стандартная библиотека Go**.
- Правильно обрабатываются **ошибки и статус-коды**.
- Используется **Docker** (база данных поднимается в контейнере).

#### Примеры запросов (curl)

```bash
# Получить все задачи
curl http://localhost:8080/tasks/

# Получить задачу по ID
curl http://localhost:8080/tasks/1

# Создать новую задачу
curl -X POST http://localhost:8080/tasks/create \
  -H "Content-Type: application/json" \
  -d '{"title": "Записать видео про REST API", "description": "Создать обучающее видео", "completed": false}'

# Обновить задачу
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Обновлённое название","description":"Новое описание","completed":true}'
```

---

### Структура проекта

```
.
├── api/
├── cmd/
├── internal/
├── sql/
├── docker-compose.yml
├── go.mod
├── go.sum
```

---

### План доработок (что добавить дальше)

- [ ] Аутентификация (JWT, сессии)
- [ ] Валидация входных данных
- [ ] Пагинация и поиск по задачам
- [ ] Юнит-тесты и интеграционные тесты
- [ ] Документация через Swagger (OpenAPI)
- [ ] `Dockerfile` для сборки всего приложения (не только базы)
