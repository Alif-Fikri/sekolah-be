package validators

import (
	"errors"
	"regexp"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"strings"
)

func ValidateEmail(email string) error {
	var teacher models.Teacher
	if err := database.DB.Where("email = ?", email).First(&teacher).Error; err == nil {
		return errors.New("email sudah digunakan")
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if !re.MatchString(email) {
		return errors.New("email tidak valid, harus berakhiran '.com'")
	}

	if !strings.HasSuffix(email, ".com") {
		return errors.New("email harus berakhiran '.com'")
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

func ValidateLoginInput(input requests.RegisterGuruRequest) error {
	if input.Email == "" {
		return errors.New("email wajib diisi")
	}
	if input.Password == "" {
		return errors.New("password wajib diisi")
	}
	return nil
}

func ValidateNIK(nik string) error {
	var teacher models.Teacher
	if err := database.DB.Where("nik = ?", nik).First(&teacher).Error; err == nil {
		return errors.New("NIK sudah digunakan")
	}
	if len(nik) != 16 {
		return errors.New("NIK harus terdiri dari 16 digit angka")
	}
	isNumeric := regexp.MustCompile(`^\d{16}$`).MatchString
	if !isNumeric(nik) {
		return errors.New("NIK hanya boleh berisi angka")
	}
	return nil
}
