package controllers

import (
	"net/http"
	"strings"

	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Data login tidak valid: "+err.Error())
		return
	}

	var teacher models.Teacher
	if err := database.DB.Where("email = ?", input.Email).First(&teacher).Error; err == nil {
		if !utils.CheckPasswordHash(input.Password, teacher.Password) {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Password salah")
			return
		}

		tokenString, err := utils.GenerateToken(teacher.ID, teacher.Role)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal menghasilkan token")
			return
		}

		session := models.Session{
			TeacherID: &teacher.ID,
			Token:     tokenString,
			Role:      teacher.Role,
		}
		if err := database.DB.Create(&session).Error; err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal menyimpan sesi login")
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "Login guru berhasil", gin.H{
			"token": tokenString,
			"role":  teacher.Role,
		})
		return
	}

	var student models.Student
	if err := database.DB.Where("email = ?", input.Email).First(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Email tidak ditemukan")
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

	utils.SuccessResponse(c, http.StatusOK, "Login siswa berhasil", gin.H{
		"token": tokenString,
		"role":  "siswa",
	})
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Token tidak ditemukan")
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Format token tidak valid")
		return
	}
	tokenString := tokenParts[1]

	if err := database.DB.Where("token = ?", tokenString).Delete(&models.Session{}).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal logout")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Logout berhasil", nil)
}
