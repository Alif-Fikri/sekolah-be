package models

import (
	"time"
)

type Student struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	NISN        string `gorm:"size:20;unique;not null"`
	Password    string `gorm:"not null"`
	SchoolLevel string `gorm:"type:ENUM('SD','SMP','SMA')"`
	IsSLB       bool
	BirthPlace  string `gorm:"size:100"`
	BirthDate   time.Time
	Gender      string     `gorm:"type:ENUM('L','P')"`
	Address     string     `gorm:"type:text"`
	Email       string     `gorm:"size:100;unique"`
	Phone       string     `gorm:"size:20"`
	Classes     []*Class   `gorm:"many2many:class_students;"`
	Subjects    []*Subject `gorm:"many2many:subject_students;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
