package models

import (
	"time"
)

type ClassAttendance struct {
	ID             uint      `gorm:"primaryKey"`
	StudentID      uint      `gorm:"not null"`
	ClassID        uint      `gorm:"not null"`
	AttendanceTime time.Time `gorm:"not null"`
	Status         string    `gorm:"type:ENUM('hadir','izin','sakit','alpha');not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Student Student
	Class   Class
}

type SubjectAttendance struct {
	ID             uint      `gorm:"primaryKey"`
	StudentID      uint      `gorm:"not null"`
	SubjectID      uint      `gorm:"not null"`
	AttendanceTime time.Time `gorm:"not null"` 
	Status         string    `gorm:"type:ENUM('hadir','izin','sakit','alpha');not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Student Student
	Subject Subject
}
