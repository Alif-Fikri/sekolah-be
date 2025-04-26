package controllers

import (
	"net/http"

	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"
	"sekolah-be/validators"

	"github.com/gin-gonic/gin"
)


func LoginTeacher(c *gin.Context) {
	var input validators.LoginGuruRequest

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil login",
		"token":   tokenString,
	})
}
