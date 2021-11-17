package db

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ Repo = (*dbRepo)(nil)

const logLevel = "info"
const enableFileLog = true

type Repo interface {
	i()
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
}

type dbRepo struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

func New() (Repo, error) {
	cfg := configs.Get().MySQL
	dbr, err := dbConnect(cfg.Read.User, cfg.Read.Pass, cfg.Read.Addr, cfg.Read.Name)
	if err != nil {
		return nil, err
	}

	dbw, err := dbConnect(cfg.Write.User, cfg.Write.Pass, cfg.Write.Addr, cfg.Write.Name)
	if err != nil {
		return nil, err
	}

	return &dbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDbR() *gorm.DB {
	return d.DbR
}

func (d *dbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

func (d *dbRepo) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *dbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

type myLogger struct {
	*log.Logger
}

func (l *myLogger) Printf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf(format, v...))
}

func NewMyLogger(logger *log.Logger) *myLogger {
	return &myLogger{Logger: logger}
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if enableFileLog {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename: "./logs/gorm.log",
			//MaxSize:    global.App.Config.Log.MaxSize,
			//MaxBackups: global.App.Config.Log.MaxBackups,
			//MaxAge:     global.App.Config.Log.MaxAge,
			//Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return NewMyLogger(log.New(writer, "\r\n", log.LstdFlags))
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch logLevel {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold: 200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:      logMode,                // 日志级别
	})
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: getGormLogger(),
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	cfg := configs.Get().MySQL.Base

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * cfg.ConnMaxLifeTime)

	// 使用插件
	err = db.Use(&TracePlugin{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
