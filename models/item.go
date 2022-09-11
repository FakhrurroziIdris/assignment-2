package models

import "time"

type Item struct {
	Item_ID     uint64  `gorm:"primaryKey"`
	Item_Code   string  `gorm:"not null"`
	Description string  `gorm:"type:varchar(255)"`
	Quantity    float32 `gorm:"not null;type:float"`
	Order_ID    uint64  `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
