package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-juno/juno/example/juno/internal/constant"
	"github.com/go-juno/juno/example/juno/internal/model"

	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 连接数据库
func NewMysqlDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		constant.Config.Database.Mysql.User,
		constant.Config.Database.Mysql.Password,
		constant.Config.Database.Mysql.Host,
		constant.Config.Database.Mysql.Port,
		constant.Config.Database.Mysql.Database)
	dia := mysql.Open(dsn)

	config := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		config,
	)
	db, err = gorm.Open(dia, &gorm.Config{Logger: newLogger})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if constant.RELEASE {
		err = db.AutoMigrate(
			&model.Greeting{},
		)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
	}
	return
}
