package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase connects to the PostgreSQL database.
func ConnectDatabase() (*gorm.DB, error) {
	dsn := "user=testuser password=testpassword dbname=testdb sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
