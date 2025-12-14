package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	DB *gorm.DB
}

func NewDatabaseConfig(databaseURL string) (*DatabaseConfig, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection established successfully")

	return &DatabaseConfig{DB: db}, nil
}

func (d *DatabaseConfig) GetDB() *gorm.DB {
	return d.DB
}

func (d *DatabaseConfig) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
