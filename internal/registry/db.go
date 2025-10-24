package registry

import (
	"fmt"
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/internal/dao"
	"gorm.io/driver/mysql"
	l "gorm.io/gorm/logger"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"gorm.io/gorm"
	"sync"
)

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
func RegisterDB(config configs.Config) *gorm.DB {
	once.Do(func() {
		var err error
		var conn gorm.Dialector
		if config.Mysql != "" {
			conn = mysql.Open(config.Mysql)
		} else if os.Getenv("SQL_DSN") != "" {
			conn = mysql.Open(os.Getenv("SQL_DSN"))
		} else {
			// 使用 SQLite，从配置读取路径，默认为 tiku.db（向后兼容）
			dbPath := config.Database.Path
			if dbPath == "" {
				dbPath = "tiku.db"
			}

			// 确保数据库文件所在目录存在
			dbDir := filepath.Dir(dbPath)
			if dbDir != "." && dbDir != "" {
				if err := os.MkdirAll(dbDir, 0755); err != nil {
					logger.FatalLog(fmt.Errorf("create database directory fail: %w", err))
				}
			}

			conn = sqlite.Open(dbPath)
		}
		db, err = gorm.Open(conn, &gorm.Config{
			PrepareStmt: true,
			Logger:      l.Default.LogMode(l.Info),
		})
		if err != nil {
			logger.FatalLog(fmt.Errorf("open db fail: %w", err))
		}
		dao.SetDefault(db)
	})
	err := db.AutoMigrate(&entity.Tiku{})
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}

	err = db.AutoMigrate(&entity.Log{})
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}

	// 首先需要注册一个默认用户
	user := entity.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Nickname: "管理员",
	}
	_ = dao.User.Create(&user)
	return db
}
