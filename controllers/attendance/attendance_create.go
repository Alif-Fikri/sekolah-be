package controllers

import (
	"net/http"
	"time"

	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
)

func CreateClassAttendance(c *gin.Context) {
	var req requests.ClassAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	attendance := models.ClassAttendance{
		StudentID:      req.StudentID,
		ClassID:        req.ClassID,
		Status:         req.Status,
		AttendanceTime: time.Now(),
	}

	if err := database.DB.Create(&attendance).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal menyimpan absensi kelas")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Absensi kelas berhasil dicatat", attendance)
}

func CreateSubjectAttendance(c *gin.Context) {
	var req requests.SubjectAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	attendance := models.SubjectAttendance{
		StudentID:      req.StudentID,
		SubjectID:      req.SubjectID,
		Status:         req.Status,
		AttendanceTime: time.Now(),
	}

	if err := database.DB.Create(&attendance).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal menyimpan absensi mapel")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Absensi mapel berhasil dicatat", attendance)
}