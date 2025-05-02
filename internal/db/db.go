package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

// Инициализация подключения к базе данных
func InitDB() {
	log.Println("Loading .env file")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	} else {
		log.Println(".env file loaded successfully")
	}

	// Чтение параметров подключения из .env файла
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	sslmode := "disable"

	// Формирование строки подключения
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPass, dbName, dbHost, dbPort, sslmode)

	var sqlDB *sql.DB
	var err error

	// Retry логика: пытаемся подключиться к PostgreSQL до 10 раз
	for i := 0; i < 10; i++ {
		sqlDB, err = sql.Open("postgres", connStr)
		if err == nil {
			pingErr := sqlDB.Ping()
			if pingErr == nil {
				log.Println("Connected to PostgreSQL database")
				break
			}
			log.Println("Postgres not ready yet, retrying...")
		} else {
			log.Println("SQL Open error:", err)
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to DB after retries: %v", err)
	}

	// Инициализация миграций
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v", err)
	}
	log.Println("Current working dir:", cwd)

	// Путь к миграциям
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
	log.Println("Migrations path:", migrationsURL)

	// Применение миграций
	m, err := migrate.NewWithDatabaseInstance(migrationsURL, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	// Применение миграции (если нет изменений, миграция не будет применена)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// Подключение к базе данных с использованием GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB
}
