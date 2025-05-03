package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UpdateClass(c *gin.Context) {
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

	var req requests.ClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data tidak valid: "+err.Error())
		return
	}

	class.NameKelas = req.NameKelas
	class.ClassLevel = req.ClassLevel
	class.SchoolLevel = req.SchoolLevel

	if err := database.DB.Save(&class).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal update class")
		return
	}

	if err := database.DB.Preload("GuruPengampu").First(&class, class.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil data class setelah update")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil update class", class)
}
