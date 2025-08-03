package models

import "time"

type Subject struct {
	ID            uint   `gorm:"primaryKey"`
	MataPelajaran string `gorm:"size:100;not null"`
	SchoolLevel   string `gorm:"type:ENUM('SD', 'SMP', 'SMA');not null"`
	CreatedByID   uint
	CreatedBy     Teacher    `gorm:"foreignKey:CreatedByID"`
	Students      []*Student `gorm:"many2many:subject_students;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
