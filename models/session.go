package models

import "time"

type Session struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeacherID *uint      `json:"teacher_id,omitempty"`
	StudentID *uint      `json:"student_id,omitempty"`
	Token     string     `gorm:"type:varchar(512);unique;not null" json:"token"`
	Role      string     `gorm:"type:enum('guru','siswa');not null" json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
