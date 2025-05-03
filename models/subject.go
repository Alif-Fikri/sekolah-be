package models

import "time"

type Subject struct {
	ID            uint   `gorm:"primaryKey"`
	MataPelajaran string `gorm:"size:100;not null"`
	SchoolLevel   string `gorm:"type:ENUM('SD', 'SMP', 'SMA');not null"`
	CreatedByID   uint
	CreatedBy     Teacher `gorm:"foreignKey:CreatedByID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
