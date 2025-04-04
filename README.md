1. Введение в проект (Overview)
Название проекта: Movie CRUD API

Цель проекта: Создание REST API для управления фильмами с аутентификацией пользователей.

Технологии: Go, Gin, GORM, PostgreSQL.

Описание: Этот проект представляет собой RESTful API для управления фильмами, позволяя пользователям добавлять, обновлять, удалять и просматривать фильмы через HTTP запросы. Используются фреймворк Gin для реализации API и GORM для работы с базой данных PostgreSQL.

2. Технологии (Technologies Used)
Go (Golang): Язык программирования, используемый для разработки высокопроизводительных приложений.

Gin: Легковесный и быстрый HTTP-фреймворк для Go, идеально подходящий для создания API.
 
GORM: ORM для Go, упрощает взаимодействие с базами данных, такими как PostgreSQL, с помощью моделей и миграций.

PostgreSQL: Реляционная база данных для хранения данных о фильмах и пользователях.

3. Установка и настройка (Setup and Installation)
Клонирование репозитория:

bash
Копировать
Редактировать
git clone https://github.com/your-username/movie-crud-api.git
cd movie-crud-api
Установка зависимостей:

bash
Копировать
Редактировать
go mod tidy
Запуск сервера:

bash
Копировать
Редактировать
go run main.go
4. Основные маршруты API (Endpoints)
GET /api/movies/: Получить список всех фильмов.

POST /api/movies/: Создать новый фильм.

GET /api/movies/:id: Получить информацию о фильме по ID.

PUT /api/movies/:id: Обновить информацию о фильме.

DELETE /api/movies/:id: Удалить фильм по ID.

Для пользователя:

GET /api/users/: Получить список всех пользователей.

POST /api/users/: Создать нового пользователя.

GET /api/users/:id: Получить информацию о пользователе по ID.

PUT /api/users/:id: Обновить информацию о пользователе.

DELETE /api/users/:id: Удалить пользователя по ID.

5. Примеры запросов и ответов (Usage)
POST /api/movies/ Пример запроса:

json
Копировать
Редактировать
{
  "title": "Inception",
  "description": "A mind-bending thriller",
  "release_date": "2010-07-16",
  "genre": "Sci-Fi"
}
Пример ответа:

json
Копировать
Редактировать
{
  "id": 1,
  "title": "Inception",
  "description": "A mind-bending thriller",
  "release_date": "2010-07-16",
  "genre": "Sci-Fi"
}
GET /api/movies/:id Пример запроса: GET /api/movies/1 Пример ответа:

json
Копировать
Редактировать
{
  "id": 1,
  "title": "Inception",
  "description": "A mind-bending thriller",
  "release_date": "2010-07-16",
  "genre": "Sci-Fi"
}
6. Работа с базой данных и GORM (Database and GORM)
Подключение к PostgreSQL: В файле main.go:

go
Копировать
Редактировать
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
Миграции с GORM:

go
Копировать
Редактировать
db.AutoMigrate(&Movie{}, &User{})
CRUD-операции:

Создание:

go
Копировать
Редактировать
db.Create(&movie)
Чтение:

go
Копировать
Редактировать
db.First(&movie, 1)
Обновление:

go
Копировать
Редактировать
db.Model(&movie).Update("Genre", "Action")
Удаление:

go
Копировать
Редактировать
db.Delete(&movie)
7. Пример работы с JSON (Working with JSON)
Gin автоматически обрабатывает JSON-запросы и ответы. Например:

go
Копировать
Редактировать
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
8. Преимущества использования Gin и GORM
Gin:

Быстрый, с низкой нагрузкой на ресурсы.

Легко масштабируемый.

Прост в использовании, благодаря встроенным middleware для логирования и обработки ошибок.

GORM:

Удобен для работы с базой данных через модели.

Поддерживает миграции, чтобы изменения в структуре данных были простыми и безопасными.

Поддерживает работу с реляционными связями между моделями.

9. Заключение (Conclusion)
Этот проект представляет собой простую и эффективную систему для управления фильмами и пользователями. Он использует мощные инструменты, такие как Gin для создания API и GORM для работы с базой данных. Проект позволяет легко управлять фильмами и пользователями, а также поддерживает полный набор операций CRUD.

