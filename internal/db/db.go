package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq" // Подключение драйвера для PostgreSQL
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	// Получаем параметры для подключения к базе данных из .env файла
	databaseName := "postgres"
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	sslmode := "disable"

	// Формируем строку подключения
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUser, dbPass, dbName, dbHost, dbPort, sslmode)

	// Открываем соединение с базой данных
	sqlDB, err := sql.Open(databaseName, connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Настроим миграции
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v", err)
	}
	log.Println("Current working dir:", cwd)

	// Указываем путь для миграций
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
	log.Println("Migrations path:", migrationsURL)

	// Выполняем миграции
	m, err := migrate.NewWithDatabaseInstance(migrationsURL, databaseName, driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	// Открываем соединение с GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB
}
