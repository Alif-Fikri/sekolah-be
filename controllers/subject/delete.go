package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteSubject(c *gin.Context) {
	id := c.Param("id")

	var subject models.Subject
	if err := database.DB.First(&subject, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "subject tidak ditemukan")
		return
	}

	userClaims := c.MustGet("user").(jwt.MapClaims)
	loggedInTeacherID := uint(userClaims["user_id"].(float64))

	if subject.CreatedByID != loggedInTeacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "akses ditolak: Anda bukan pembuat subject ini")
		return
	}

	if err := database.DB.Delete(&subject).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal menghapus subject")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil menghapus subject", nil)
}
