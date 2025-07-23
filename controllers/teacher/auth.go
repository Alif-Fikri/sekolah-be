package controllers

import (
	"net/http"

	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"
	"sekolah-be/requests"
	"sekolah-be/validators"
	"github.com/gin-gonic/gin"
)

func LoginTeacher(c *gin.Context) {
	var input requests.LoginGuruRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data login tidak valid: "+err.Error())
		return
	}

	var teacher models.Teacher
	if err := database.DB.Where("email = ?", input.Email).First(&teacher).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "email belum terdaftar")
		return
	}

	if !utils.CheckPasswordHash(input.Password, teacher.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "password salah")
		return
	}

	tokenString, err := utils.GenerateToken(teacher.ID, teacher.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal menghasilkan token")
		return
	}

	session := models.Session{
		TeacherID: teacher.ID,
		Token:     tokenString,
		Role:      teacher.Role,
	}
	if err := database.DB.Create(&session).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal menyimpan sesi login")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Berhasil login", gin.H{
		"token": tokenString,
	})
}

func RegisterTeacher(c *gin.Context) {
	var input requests.RegisterGuruRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data registrasi tidak valid: "+err.Error())
		return
	}

	if err := validators.ValidateEmail(input.Email); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePassword(input.Password); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengenkripsi password")
		return
	}

	if err := validators.ValidateNIK(input.NIK); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	teacher := models.Teacher{
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		NIK:         input.NIK,
		NUPTK:       input.NUPTK,
		SchoolName:  input.SchoolName,
		SchoolLevel: input.SchoolLevel,
		IsSLB:       input.IsSLB,
		Role:        "guru",
	}

	if err := database.DB.Create(&teacher).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "gagal mendaftarkan guru: "+err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "guru berhasil terdaftar", nil)
}

func LogoutTeacher(c *gin.Context) {
	tokenValue, exists := c.Get("tokenString")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "token tidak ditemukan dalam context")
		return
	}

	tokenString, ok := tokenValue.(string)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "token invalid")
		return
	}

	if err := database.DB.Where("token = ?", tokenString).Delete(&models.Session{}).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal logout")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil logout", nil)
}
