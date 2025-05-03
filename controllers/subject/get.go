package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetSubjectByID(c *gin.Context) {
	idParam := c.Param("id")
	subjectID, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var subject models.Subject
	if err := database.DB.Preload("CreatedBy").First(&subject, subjectID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Subject tidak ditemukan")
		return
	}

	claims := c.MustGet("user").(jwt.MapClaims)
	teacherID := uint(claims["user_id"].(float64))
	if subject.CreatedByID != teacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "akses ditolak: bukan subject Anda")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "data subject ditemukan", subject)
}

func GetAllSubjects(c *gin.Context) {
	claims := c.MustGet("user").(jwt.MapClaims)
	createdBy := uint(claims["user_id"].(float64))

	var subjects []models.Subject
	if err := database.DB.Preload("CreatedBy").Where("created_by_id = ?", createdBy).Find(&subjects).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil data subject")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "data subject ditemukan", subjects)
}
