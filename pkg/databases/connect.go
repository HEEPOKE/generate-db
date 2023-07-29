package databases

import (
	"fmt"

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

	db, err := gorm.Open(driver, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
