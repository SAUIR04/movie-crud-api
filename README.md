
---

# 🎬 **Movie CRUD API**

**Movie CRUD API** — это RESTful API, предназначенное для управления фильмами с аутентификацией пользователей. Это API позволяет добавлять, обновлять, удалять и просматривать фильмы с использованием технологий Go, Gin и GORM, а также PostgreSQL в качестве базы данных.

## 🚀 **Технологии (Technologies Used)**

- **Go (Golang)**: Высокопроизводительный язык программирования для разработки эффективных веб-приложений.
- **Gin**: Легковесный и быстрый HTTP-фреймворк для Go, идеально подходящий для создания API.
- **GORM**: ORM для Go, который упрощает взаимодействие с базой данных (PostgreSQL).
- **PostgreSQL**: Реляционная база данных для хранения данных о фильмах и пользователях.

## 🔧 **Установка и настройка (Setup and Installation)**

1. Клонировать репозиторий:
    ```bash
    git clone https://github.com/your-username/movie-crud-api.git
    cd movie-crud-api
    ```

2. Установить зависимости:
    ```bash
    go mod tidy
    ```

3. Запуск сервера:
    ```bash
    go run main.go
    ```

Сервер будет запущен на [http://localhost:8080](http://localhost:8080).

## 🛠 **Основные маршруты API (Endpoints)**

### Фильмы

- **GET /api/movies**: Получить список всех фильмов.
- **POST /api/movies**: Создать новый фильм.
- **GET /api/movies/:id**: Получить информацию о фильме по ID.
- **PUT /api/movies/:id**: Обновить информацию о фильме.
- **DELETE /api/movies/:id**: Удалить фильм по ID.

### Пользователи

- **GET /api/users**: Получить список всех пользователей.
- **POST /api/users**: Создать нового пользователя.
- **GET /api/users/:id**: Получить информацию о пользователе по ID.
- **PUT /api/users/:id**: Обновить информацию о пользователе.
- **DELETE /api/users/:id**: Удалить пользователя по ID.

## 🔍 **Примеры запросов и ответов (Usage Examples)**

### 1. POST /api/movies

**Запрос**:
```json
{
  "title": "Inception",
  "description": "A mind-bending thriller",
  "release_date": "2010-07-16",
  "genre": "Sci-Fi"
}
```

**Ответ**:
```json
{
  "id": 1,
  "title": "Inception",
  "description": "A mind-bending thriller",
  "release_date": "2010-07-16",
  "genre": "Sci-Fi"
}
```

### 2. GET /api/movies/:id

**Запрос**:
```
GET /api/movies/1
```

**Ответ**:
```json
{
  "id": 1,
  "title": "Inception",
  "description": "A mind-bending thriller",
  "release_date": "2010-07-16",
  "genre": "Sci-Fi"
}
```

## 🗄 **Работа с базой данных и GORM (Database and GORM)**

### Подключение к PostgreSQL:
В файле `main.go`:

```go
import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func main() {
    dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Ошибка подключения к БД:", err)
    }
    log.Println("Подключение успешно!")
}
```

### Миграции с GORM:
```go
db.AutoMigrate(&Movie{}, &User{})
```

### CRUD-операции:

**Создание**:
```go
db.Create(&movie)
```

**Чтение**:
```go
db.First(&movie, 1)
```

**Обновление**:
```go
db.Model(&movie).Update("Genre", "Action")
```

**Удаление**:
```go
db.Delete(&movie)
```

## 📄 **Работа с JSON в Gin**

Gin автоматически обрабатывает JSON-запросы и ответы. Пример:

```go
type Movie struct {
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Genre       string `json:"genre"`
}

r.POST("/api/movies/", func(c *gin.Context) {
    var movie Movie
    if err := c.ShouldBindJSON(&movie); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db.Create(&movie)
    c.JSON(200, gin.H{"message": "Movie created successfully", "movie": movie})
})
```

## 💡 **Преимущества использования Gin и GORM**

### Gin:
- Высокая скорость.
- Легкость в использовании и масштабируемость.
- Встроенные middleware для логирования и обработки ошибок.

### GORM:
- Удобство в работе с базой данных через модели.
- Миграции для безопасных изменений структуры данных.
- Поддержка реляционных связей между моделями.

## ✅ **Заключение (Conclusion)**

Этот проект предоставляет API для управления фильмами и пользователями, используя лучшие инструменты Go, такие как **Gin** и **GORM**. Система поддерживает полный набор операций CRUD для управления фильмами и пользователями, обеспечивая простоту и эффективность работы с данными.

---

