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
	if err := database.DB.
		Preload("GuruPengampu").
		Where("guru_pengampu_id = ?", guruID).
		Find(&classes).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil data class")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "data class ditemukan", classes)
}

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

func DeleteClass(c *gin.Context) {
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

	if err := database.DB.Delete(&models.Class{}, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal menghapus class")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil menghapus class", nil)
}
