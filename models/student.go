package models

import (
	"time"
)

type Student struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"size:100;not null"`
	NISN        string         `gorm:"size:20;unique;not null"`
	Password    string         `gorm:"not null"`
	ClassID     uint           `gorm:"not null"`
	Class       Class          `gorm:"foreignKey:ClassID"`
	SchoolLevel string         `gorm:"type:ENUM('SD','SMP','SMA')"`
	IsSLB       bool
	BirthPlace  string         `gorm:"size:100"`
	BirthDate   time.Time
	Gender      string         `gorm:"type:ENUM('L','P')"` // L: Laki-laki, P: Perempuan
	Address     string         `gorm:"type:text"`
	Email       string         `gorm:"size:100;unique"`
	Phone       string         `gorm:"size:20"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
