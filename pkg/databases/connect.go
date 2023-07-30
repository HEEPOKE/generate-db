package databases

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HEEPOKE/generate-db/pkg/config"
	"github.com/HEEPOKE/generate-db/pkg/enums"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDriver(dbType string) (gorm.Dialector, error) {
	switch dbType {
	case string(enums.POSTGRESQL):
		return postgres.Open(config.Cfg.DB_DSN), nil
	case string(enums.MYSQL):
		return mysql.Open(config.Cfg.DB_DSN), nil
	case string(enums.SQLSERVER):
		return sqlserver.Open(config.Cfg.DB_DSN), nil
	default:
		return nil, fmt.Errorf("unsupported DB_TYPE specified: %s", dbType)
	}
}

func ConnectDB() (*gorm.DB, error) {
	driver, err := getDriver(config.Cfg.DB_TYPE)
	if err != nil {
		return nil, err
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(driver, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
