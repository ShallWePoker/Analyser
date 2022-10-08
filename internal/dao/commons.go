package dao

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/database"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/loggers"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"reflect"
)

var db = database.Db
var logger = loggers.Log

func InitTables() {
	createTableIfNotExists(models.SysUser{})
}

func createTableIfNotExists(modelType interface{}) {
	val := reflect.Indirect(reflect.ValueOf(modelType))
	modelName := val.Type().Name()

	logger.Info(fmt.Sprintf("Migrating Table of %s ...", modelName))
	err := db.AutoMigrate(modelType)
	if err != nil {
		logger.Error(fmt.Sprintf("AutoMigrate failed with error: %v", err))
	}
}