package models

import "time"

type Class struct {
	ID             uint   `gorm:"primaryKey"`
	NameKelas      string `gorm:"size:50;not null"`
	ClassLevel     string `gorm:"size:50;not null"`
	SchoolLevel    string `gorm:"type:ENUM('SD', 'SMP', 'SMA');not null"`
	GuruPengampuID uint
	GuruPengampu   Teacher    `gorm:"foreignKey:GuruPengampuID"`
	Students       []*Student `gorm:"many2many:class_students;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
