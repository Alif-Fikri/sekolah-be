package Controller

import (
	"encoding/csv"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginStudent(c *gin.Context) {
	var input requests.LoginStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Data login tidak valid: "+err.Error())
		return
	}

	var student models.Student
	if err := database.DB.Where("name = ?", input.Name).First(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Nama siswa tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(input.Password)); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Password salah")
		return
	}

	tokenString, err := utils.GenerateToken(student.ID, "siswa")
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	session := models.Session{
		StudentID: &student.ID,
		Token:     tokenString,
		Role:      "siswa",
	}
	if err := database.DB.Create(&session).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal menyimpan sesi login")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login berhasil", gin.H{
		"token": tokenString,
	})
}

func RegisterStudent(c *gin.Context) {
	var input requests.RegisterStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data siswa tidak valid: "+err.Error())
		return
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "format tanggal lahir tidak valid (YYYY-MM-DD)")
		return
	}

	passwordRaw := birthDate.Format("02012006")
	passwordHash, err := utils.HashPassword(passwordRaw)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal generate password")
		return
	}

	student := models.Student{
		Name:        input.Name,
		NISN:        input.NISN,
		Password:    passwordHash,
		SchoolLevel: input.SchoolLevel,
		IsSLB:       input.IsSLB,
		BirthPlace:  input.BirthPlace,
		BirthDate:   birthDate,
		Gender:      input.Gender,
		Address:     input.Address,
		Email:       input.Email,
		Phone:       input.Phone,
	}

	if err := database.DB.Create(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "gagal mendaftarkan siswa: "+err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "siswa berhasil didaftarkan", gin.H{
		"password_awal": passwordRaw,
	})
}

func LogoutStudent(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required")
        return
    }

    tokenParts := strings.Split(authHeader, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
        utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token format")
        return
    }
    tokenString := tokenParts[1]

    if err := database.DB.Where("token = ?", tokenString).Delete(&models.Session{}).Error; err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to logout")
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Logout successful", nil)
}

func ImportStudents(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "File CSV wajib diunggah")
		return
	}

	src, err := file.Open()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuka file")
		return
	}
	defer src.Close()

	reader := csv.NewReader(src)
	rows, err := reader.ReadAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Format CSV tidak valid")
		return
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 10 {
			continue
		}

		birthDate, err := time.Parse("2006-01-02", row[5])
		if err != nil {
			continue
		}

		passwordRaw := birthDate.Format("02012006")
		passwordHash, err := utils.HashPassword(passwordRaw)
		if err != nil {
			continue
		}

		isSLB, _ := strconv.ParseBool(row[3])

		student := models.Student{
			Name:        row[0],
			NISN:        row[1],
			Password:    passwordHash,
			SchoolLevel: row[2],
			IsSLB:       isSLB,
			BirthPlace:  row[4],
			BirthDate:   birthDate,
			Gender:      row[6],
			Address:     row[7],
			Email:       row[8],
			Phone:       row[9],
		}

		database.DB.Create(&student)
	}

	utils.SuccessResponse(c, http.StatusOK, "Data siswa berhasil diimpor", nil)
}
