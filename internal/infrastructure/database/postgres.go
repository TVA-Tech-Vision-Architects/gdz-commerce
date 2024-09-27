package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/B6137151/GDZ-Commerce/migrations"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase() (*gorm.DB, error) {
	envPath := findEnvFile()
	if envPath == "" {
		return nil, fmt.Errorf(".env file not found")
	}

	err := godotenv.Load(envPath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	if err := migrations.CreateStoreTable(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB instance: %w", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	return db, nil
}

func findEnvFile() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		filePath := filepath.Join(dir, ".env")
		if _, err := os.Stat(filePath); err == nil {
			return filePath
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
