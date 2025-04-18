package registry

import (
	"fmt"
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/internal/dao"
	"gorm.io/driver/mysql"
	l "gorm.io/gorm/logger"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"gorm.io/gorm"
	"sync"
)

var dbName = "tiku.db"
var db *gorm.DB
var once sync.Once
// 用于判断是否在使用 MySQL
var isMySQL bool

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
		var err error  //修复  undefined: err

	once.Do(func() {
		var conn gorm.Dialector
		if config.Mysql != "" {
			conn = mysql.Open(config.Mysql)
			isMySQL = true
		} else if os.Getenv("SQL_DSN") != "" {
			conn = mysql.Open(os.Getenv("SQL_DSN"))
			isMySQL = true
		} else {
			conn = sqlite.Open(dbName)
			isMySQL = false
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
	
	//以后能不能把sqlite3和Mysql数据的注册代码分开啊
	//autoincrement 是 SQLite 的自增语法
	//AUTO_INCREMENT 是 Mysql 的自增语法
	
	if isMySQL==false {
	err = db.AutoMigrate(&entity.Tiku{})
	} else {
	err = db.AutoMigrate(&entity.TikuMysql{})
	}
	
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}
	
	//以后能不能把sqlite3和Mysql数据的注册代码分开啊
	//autoincrement 是 SQLite 的自增语法
	//AUTO_INCREMENT 是 Mysql 的自增语法
	
	if isMySQL==false {
	err = db.AutoMigrate(&entity.User{})
	} else {
	err = db.AutoMigrate(&entity.UsersMysql{})
	}
	
	if err != nil {
		logger.FatalLog(fmt.Errorf("auto migrate fail: %w", err))
	}
	
	//以后能不能把sqlite3和Mysql数据的注册代码分开啊
	//autoincrement 是 SQLite 的自增语法
	//AUTO_INCREMENT 是 Mysql 的自增语法
	
	if isMySQL==false {
	err = db.AutoMigrate(&entity.Log{})
	} else {
	err = db.AutoMigrate(&entity.LogsMysql{})
	}
	
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
