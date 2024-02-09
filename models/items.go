package models

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null; type: varchar(20)"`
	Description string `gorm:"not null; type: varchar(191)"`
	Quantity    uint   `gorm:"not null"`
	OrderId     uint   `gorm:"not null"`
}
