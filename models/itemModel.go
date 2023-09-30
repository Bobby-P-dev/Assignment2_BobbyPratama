package models

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ID          int `json:"id" form:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ItemCode    string         `json:"item_code"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	OrdersID    int            `json:"OrderID" form:"OrderID"`
}
