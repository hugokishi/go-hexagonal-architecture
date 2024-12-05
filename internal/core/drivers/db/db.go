package db

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
)

func New(connection string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger:                   logger.Default.LogMode(getLoggerLevel()),
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "gorm"}).Fatal(err)
	}

	return db
}

func getLoggerLevel() logger.LogLevel {
	loggerLevel := os.Getenv("DB_LOGGER_MODE")

	switch loggerLevel {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		return logger.Silent
	}
}

func IsError(err error) bool {
	return err != nil &&
		!strings.Contains(err.Error(), "operation was canceled") &&
		!strings.Contains(err.Error(), "context canceled")
}
