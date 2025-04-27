package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
)

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
