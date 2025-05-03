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

func UpdateSubject(c *gin.Context) {
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
	var req requests.SubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data tidak valid: "+err.Error())
		return
	}

	subject.MataPelajaran = req.MataPelajaran
	subject.SchoolLevel = req.SchoolLevel

	if err := database.DB.Save(&subject).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal update subject")
		return
	}

	if err := database.DB.Preload("CreatedBy").First(&subject, subject.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal memuat data subject setelah update")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil update subject", subject)
}
