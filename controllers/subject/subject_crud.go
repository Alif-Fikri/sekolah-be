package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"
	"strconv"

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
