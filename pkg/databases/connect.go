package databases

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/generate-db/internals/core/utils"
	"github.com/HEEPOKE/generate-db/pkg/config"
	"github.com/HEEPOKE/generate-db/pkg/enums"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDriver(dbType string) (gorm.Dialector, error) {
	var dialector gorm.Dialector
	switch dbType {
	case string(enums.POSTGRESQL):
		dialector = postgres.Open(config.Cfg.DB_DSN)
	case string(enums.MYSQL):
		dialector = mysql.Open(config.Cfg.DB_DSN)
	case string(enums.SQLSERVER):
		dialector = sqlserver.Open(config.Cfg.DB_DSN)
	default:
		log.Printf("unsupported DB_TYPE specified: %v", dbType)
		return nil, fmt.Errorf("unsupported DB_TYPE specified: %s", dbType)
	}
	return dialector, nil
}

func ConnectDB() (*gorm.DB, error) {
	driver, err := getDriver(config.Cfg.DB_TYPE)
	if err != nil {
		return nil, err
	}

	newLogger := utils.ConfigureLoggerDB(logger.Error)

	db, err := gorm.Open(driver, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
