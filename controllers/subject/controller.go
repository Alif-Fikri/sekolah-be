package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
)

func AssignStudentsToSubject(c *gin.Context) {
	var req requests.AssignStudentToClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Data tidak valid")
		return
	}

	var class models.Class
	if err := database.DB.First(&class, req.ClassID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Kelas tidak ditemukan")
		return
	}

	var students []models.Student
	if err := database.DB.Where("id IN ?", req.StudentIDs).Find(&students).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data siswa")
		return
	}

	if err := database.DB.Model(&class).Association("Students").Append(students); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal menambahkan siswa ke kelas")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Siswa berhasil ditambahkan ke kelas", nil)
}
