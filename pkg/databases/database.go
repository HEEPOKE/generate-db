package databases

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Cfg.POSTGRES_HOST, config.Cfg.POSTGRES_USER, config.Cfg.POSTGRES_PASSWORD, config.Cfg.POSTGRES_DATABASE, config.Cfg.POSTGRES_PORT, config.Cfg.POSTGRES_SSL, config.Cfg.POSTGRES_TIMEZONE)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Generate{})

	return db, nil
}
