package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null; type: varchar(191)"`
	OrderAt      string `gorm:"not null; type: varchar(191)"`
	Items        []Items
}
