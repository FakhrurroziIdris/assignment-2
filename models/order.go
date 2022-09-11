package models

import "time"

type Order struct {
	Order_ID      uint64 `gorm:"primaryKey"`
	Customer_Name string `gorm:"not null;type:varchar(255)"`
	Ordered_At    time.Time
	UpdatedAt     time.Time
	Items         []Item `gorm:"not null"`
}
