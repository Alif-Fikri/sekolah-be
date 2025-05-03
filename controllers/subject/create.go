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

func CreateSubject(c *gin.Context) {

	claims := c.MustGet("user").(jwt.MapClaims)
	createdBy := uint(claims["user_id"].(float64))

	var req requests.SubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data tidak valid: "+err.Error())
		return
	}

	subject := models.Subject{
		MataPelajaran: req.MataPelajaran,
		SchoolLevel:   req.SchoolLevel,
		CreatedByID:   createdBy,
	}

	if err := database.DB.Create(&subject).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal membuat subject")
		return
	}

	if err := database.DB.Preload("CreatedBy").First(&subject, subject.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil subject yang baru dibuat")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil membuat subject", subject)
}
