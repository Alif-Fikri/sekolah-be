package models

import "time"
import "gorm.io/datatypes"

type Exam struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	SubjectID  uint   `gorm:"not null"`
	ClassID    uint   `gorm:"not null"`
	TypeSeal   string `gorm:"type:ENUM('Wide','Paper');not null"`
	UploadType string `gorm:"type:ENUM('manual','event');not null"`
	Token      string `gorm:"size:100;not null"`
	CreatedBy  uint   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ExamResult struct {
	ID         uint           `gorm:"primaryKey"`
	StudentID  uint           `gorm:"not null"`
	ExamID     uint           `gorm:"not null"`
	Score      float64        `gorm:"type:float"`
	Confidence float64        `gorm:"type:float"`
	Answer     datatypes.JSON `gorm:"type:json"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}