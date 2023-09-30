package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	ID           int    `json:"id" form:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	CustomerName string `json:"customer_name" form:"customer_name" gorm:"not null"`
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Items        []Items
}
