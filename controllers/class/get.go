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

func GetClassByID(c *gin.Context) {
	idParam := c.Param("id")
	classID, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var class models.Class
	if err := database.DB.Preload("GuruPengampu").First(&class, classID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Kelas tidak ditemukan")
		return
	}

	user := c.MustGet("user").(jwt.MapClaims)
	teacherID := uint(user["user_id"].(float64))
	if class.GuruPengampuID != teacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "akses ditolak: bukan kelas Anda")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "data kelas ditemukan", class)
}

func GetAllClasses(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	guruID := uint(user["user_id"].(float64))

	var classes []models.Class
	if err := database.DB.Where("guru_pengampu_id = ?", guruID).Find(&classes).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil data class")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "data class ditemukan", classes)
}
