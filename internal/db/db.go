package db

import (
	"backend/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() error {
	cfg := config.Load()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	log.Printf("Connecting to database with DSN: host=%s port=%s user=%s dbname=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	if err := enableUUIDExtension(db); err != nil {
		log.Fatalf("Failed to enable uuid-ossp extension: %v", err)
	}
	return Migrate()
}

func GetDB() *gorm.DB {
	return db
}

func enableUUIDExtension(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
}
