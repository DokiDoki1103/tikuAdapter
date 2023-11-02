package registry

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"gorm.io/gorm"
	"sync"
)

var dbName = "tiku.db"
var db *gorm.DB
var once sync.Once

func closeDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
func DB() *dao.Query {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			logger.FatalLog(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
		dao.SetDefault(db)
	})
	err := db.AutoMigrate(&entity.Tiku{})
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}
	return dao.Q
}
