package validators

import (
	"errors"
	"regexp"
	"sekolah-be/database"
	"sekolah-be/models"
)

type RegisterGuruRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	NIDN        string `json:"nidn"`
	SchoolName  string `json:"school_name"`
	SchoolLevel string `json:"school_level" binding:"required,oneof=SD SMP SMA"`
	IsSLB       bool   `json:"is_slb"`
}

type LoginGuruRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func EmailUnique(email string) error {
	var teacher models.Teacher
	if err := database.DB.Where("email = ?", email).First(&teacher).Error; err == nil {
		return errors.New("email sudah digunakan")
	}
	return nil
}

func ValidatePassword(password string) error {
	var (
		hasNumber    = regexp.MustCompile(`[0-9]`).MatchString
		hasUppercase = regexp.MustCompile(`[A-Z]`).MatchString
	)

	if !hasNumber(password) || !hasUppercase(password) {
		return errors.New("password harus mengandung minimal 1 angka dan 1 huruf kapital")
	}
	return nil
}

func ValidateLoginInput(input LoginGuruRequest) error {
	if input.Email == "" {
		return errors.New("email wajib diisi")
	}
	if input.Password == "" {
		return errors.New("password wajib diisi")
	}
	return nil
}
