package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteClass(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	if err := database.DB.First(&class, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "class tidak ditemukan")
		return
	}

	userClaims := c.MustGet("user").(jwt.MapClaims)
	loggedInTeacherID := uint(userClaims["user_id"].(float64))

	if class.GuruPengampuID != loggedInTeacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "akses ditolak: Anda bukan pembuat class ini")
		return
	}

	if err := database.DB.Delete(&models.Class{}, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal menghapus class")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil menghapus class", nil)
}