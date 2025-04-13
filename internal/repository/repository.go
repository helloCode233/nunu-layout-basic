package repository

import (
	"fmt"
	"github.com/go-nunu/nunu-layout-basic/pkg/config"
	"github.com/go-nunu/nunu-layout-basic/pkg/helper/path"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pkgLog "github.com/go-nunu/nunu-layout-basic/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gorm.io/gorm/schema"
	"strconv"
)

type Repository struct {
	db     *gorm.DB
	logger *pkgLog.Logger
	conf   *config.Configuration
}

// @wire:DB
func NewRepository(conf *config.Configuration, logger *pkgLog.Logger) *Repository {
	repository := &Repository{
		logger: logger,
		conf:   conf,
	}
	switch conf.Database.Driver {

	case "mysql":
		repository.db = repository.InitMySqlGorm(conf, logger.Logger)

	case "sqlite":
		repository.db = repository.InitSqLiteGorm(conf, logger.Logger)

	default:
		panic("unknown database driver: " + conf.Database.Driver)
	}
	return repository
}

func (d *Repository) InitMySqlGorm(conf *config.Configuration, gLog *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Database.UserName,
		conf.Database.Password,
		conf.Database.Host,
		strconv.Itoa(conf.Database.Port),
		conf.Database.Database,
		conf.Database.Charset,
	)
	if db, err := gorm.Open(mysql.Open(dsn), getGormConfig(conf, gLog)); err != nil {
		gLog.Error("failed opening connection to err:", zap.Any("err", err))
		panic("failed to connect database")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.Database.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.Database.MaxOpenConns)
		err := db.AutoMigrate()
		if err != nil {
			print(err)
		}
		return db
	}

}
func getLogger(conf *config.Configuration) logger.Interface {
	var writer io.Writer
	var logMode logger.LogLevel

	if conf.Database.EnableFileLogWriter {
		logFileDir := conf.Log.RootDir
		if !filepath.IsAbs(logFileDir) {
			logFileDir = filepath.Join(path.RootPath(), logFileDir)
		}
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   filepath.Join(logFileDir, conf.Database.LogFilename),
			MaxSize:    conf.Log.MaxSize,
			MaxBackups: conf.Log.MaxBackups,
			MaxAge:     conf.Log.MaxAge,
			Compress:   conf.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}

	switch conf.Database.LogMode {
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

	return logger.New(

		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                        // 慢查询 SQL 阈值
			Colorful:                  !conf.Database.EnableFileLogWriter, // 禁用彩色打印
			IgnoreRecordNotFoundError: false,                              // 忽略ErrRecordNotFound（记录未找到）错误
			LogLevel:                  logMode,                            // Log lever
		},
	)
}
func getGormConfig(conf *config.Configuration, gLog *zap.Logger) *gorm.Config {
	var gorm_conf *gorm.Config
	if gLog != nil {
		gorm_conf = &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: conf.Database.TablePrefix,
			},
			DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
			Logger:                                   getLogger(conf), // 使用自定义 Logger
		}
	} else {
		gorm_conf = &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: conf.Database.TablePrefix,
			},
			DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
			//Logger:                                   getLogger(conf), // 使用自定义 Logger
		}
	}
	return gorm_conf
}
func (d *Repository) InitSqLiteGorm(conf *config.Configuration, gLog *zap.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(conf.Database.Database+".db"), getGormConfig(conf, gLog))
	if err != nil {
		gLog.Error("failed opening connection to err:", zap.Any("err", err))
		panic("failed to connect database")
	} else {
		err := db.AutoMigrate()
		if err != nil {
			print(err.Error())
		}
	}
	return db
}
