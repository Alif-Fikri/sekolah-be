package controllers

import (
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"
	"sekolah-be/validators"

	"github.com/gin-gonic/gin"
)

func RegisterTeacher(c *gin.Context) {
	var input validators.RegisterGuruRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data registrasi tidak valid: "+err.Error())
		return
	}

	if err := validators.ValidateEmail(input.Email); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePassword(input.Password); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengenkripsi password")
		return
	}

	if err := validators.ValidateNIK(input.NIK); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	teacher := models.Teacher{
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		NIK:         input.NIK,
		NUPTK:       input.NUPTK,
		SchoolName:  input.SchoolName,
		SchoolLevel: input.SchoolLevel,
		IsSLB:       input.IsSLB,
		Role:        "guru",
	}

	if err := database.DB.Create(&teacher).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "gagal mendaftarkan guru: "+err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "guru berhasil terdaftar", nil)
}
