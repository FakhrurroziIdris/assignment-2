package models

import (
	"time"
)

type Brand struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Cars      []Car
}

// func (b *Brand) BeforeCreate(tx *gorm.DB) (err Error) {
// 	if len(b.Name) < 3 || len(b.Name) > 50 {
// 		err = errors.New("Brand name length is invalid")
// 	}
// 	return
// }
