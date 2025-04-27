package models

import "time"

type Session struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	TeacherID uint       `gorm:"not null" json:"teacher_id"`
	Token     string     `gorm:"type:varchar(512);unique;not null" json:"token"`
	Role      string     `gorm:"type:enum('guru','siswa');not null" json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
