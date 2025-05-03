package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"
	"sekolah-be/requests"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateClass(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	guruID := uint(user["user_id"].(float64))

	var req requests.ClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data tidak valid: "+err.Error())
		return
	}

	class := models.Class{
		NameKelas:      req.NameKelas,
		ClassLevel:     req.ClassLevel,
		SchoolLevel:    req.SchoolLevel,
		GuruPengampuID: guruID,
	}

	if err := database.DB.Create(&class).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal membuat class")
		return
	}

	if err := database.DB.Preload("GuruPengampu").First(&class, class.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal memuat data class setelah create")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil membuat class", class)
}