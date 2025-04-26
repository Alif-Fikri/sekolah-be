package models

import (
	"time"
)

type Teacher struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"size:100;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	Password     string    `gorm:"not null"`
	NIDN         string    `gorm:"size:50"`
	SchoolName   string    `gorm:"size:100"`
	SchoolLevel  string    `gorm:"type:ENUM('SD', 'SMP', 'SMA');not null"`
	IsSLB        bool
	Role         string    `gorm:"type:ENUM('guru', 'siswa');default:'guru'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
