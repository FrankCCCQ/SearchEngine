package db

import (
	"SearchEngine/config"
	logs "SearchEngine/pkg/logger"
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

var _db *gorm.DB

func InitDB() {
	mConfig := config.Conf.MySQL
	host := mConfig.Host
	port := mConfig.Port
	database := mConfig.Database
	username := mConfig.Username
	password := mConfig.Password
	charset := mConfig.Charset
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	err := Database(dsn)
	if err != nil {
		logs.LogrusObj.Errorf("start database failed, original error: %T %v", errors.Cause(err), errors.Cause(err))
		logs.LogrusObj.Errorf("stack trace: \n%+v\n", err)
	}
}

func Database(connString string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		err = errors.Wrap(err, "failed to open mysql")
		return err
	}

	if err = db.Use(tracing.NewPlugin()); err != nil {
		err = errors.Wrap(err, "failed to use db plugin")
		return err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	migration()
	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	return _db.WithContext(ctx)
}
