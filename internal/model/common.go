package model

import (
	"github.com/go-nunu/nunu-layout-basic/pkg/helper/time"
	"gorm.io/gorm"
)

type Timestamps struct {
	CreatedAt time.LocalTime `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.LocalTime `gorm:"column:updated_at" json:"updatedAt"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
type ID struct {
	ID uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
}
