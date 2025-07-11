package provider

import (
	"fmt"
	"github.com/magomedcoder/ipmanager/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewPostgresClient(conf *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: conf.Postgres.GetDsn(),
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("не удалось подключиться к серверу postgres: %v", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
