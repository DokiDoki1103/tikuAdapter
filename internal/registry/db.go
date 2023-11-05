package registry

import (
	"fmt"
	l "gorm.io/gorm/logger"

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

// CloseDB close db
func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

// RegisterDB 注册数据库
func RegisterDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
			PrepareStmt: true,
			Logger:      l.Default.LogMode(l.Warn),
		})
		if err != nil {
			logger.FatalLog(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
		dao.SetDefault(db)
	})
	err := db.AutoMigrate(&entity.Tiku{})
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}
	return db
}
