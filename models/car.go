package models

import "time"

type Car struct {
	CarID     string `gorm:"primaryKey"`
	BrandID   uint
	Model     string  `gorm:"not null;type:varchar(50)"`
	Price     float32 `gorm:"not null;type:float"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Cars = []Car{}
