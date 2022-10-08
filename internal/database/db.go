package database

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/configs"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/loggers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB
var logger = loggers.Log

func init() {
	err := InitDb()
	if err != nil {
		logger.Fatal(err)
	}
}

func InitDb() error {
	dbConf := configs.Config.Db
	var err error
	switch dbConf.ServerType {
	case "postgres", "postgresql":
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Database, dbConf.Sslmode)
		Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		sqlDb, err := Db.DB()
		if err != nil {
			return err
		}
		sqlDb.SetMaxIdleConns(dbConf.MaxIdleConns)
		sqlDb.SetMaxOpenConns(dbConf.MaxOpenConns)
		sqlDb.SetConnMaxLifetime(2 * time.Hour)
		sqlDb.SetConnMaxIdleTime(1 * time.Hour)
		go func() {
			ticker := time.Tick(10 * time.Second)
			for {
				<-ticker
				_ = sqlDb.Ping()
			}
		}()
	default:
		logger.Fatalf("Unsupported database type")
	}
	return nil
}