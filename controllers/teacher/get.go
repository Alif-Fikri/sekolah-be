package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
)

func GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "guru tidak ditemukan")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil mengambil data guru", gin.H{"teacher": teacher})
}

func GetAllTeachers(c *gin.Context) {
	var teachers []models.Teacher

	if err := database.DB.Find(&teachers).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil data guru")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil mengambil semua data guru", gin.H{"teachers": teachers})
}
