package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	//data source name
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	// 	The gorm.Open function is called to open a connection to the PostgreSQL database. It takes two arguments: the first one is the GORM dialect for PostgreSQL (postgres.Open(dsn)), and the second one is a pointer to a gorm.Config instance (empty in this case, indicating default configuration).
	// If an error occurs during the connection attempt, gorm.Open will return nil for the database connection (db) and an error. Otherwise, it returns a pointer to the gorm.DB instance representing the database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
