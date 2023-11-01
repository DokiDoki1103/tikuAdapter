package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/itihey/tikuAdapter/api"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

var dbName = "tiku.db"
var db *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			logger.FatalLog(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
		dao.SetDefault(db)
	})
	err := db.AutoMigrate(&model.Tiku{})
	if err != nil {
		logger.FatalLog(fmt.Printf("Error: AutoMigrate(&model.Tiku{}) fail: %s", err))
	}
}

func closeDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	return err
}
func main() {
	logger.SetupGinLog()
	http.HandleFunc("/", api.Handler)
	defer func() {
		err := closeDB()
		if err != nil {
			logger.FatalLog(err)
		}
	}()
	if err := http.ListenAndServe(":8060", nil); err != nil {
		logger.FatalLog(err)
	}
}
