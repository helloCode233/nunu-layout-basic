package repository

import (
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	//rdb    *redis.Client
	logger *log.Logger
}

// @wire:Repository
func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		//rdb:    rdb,
		logger: logger,
	}
}

// @wire:Repository
func NewDb() *gorm.DB {
	// TODO: init db
	//db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//return db
	return &gorm.DB{}
}
