package Controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/utils"
)

type LoginStudentRequest struct {
	NISN     string `json:"nisn" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginStudent(c *gin.Context) {
	var input LoginStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Data login tidak valid: "+err.Error())
		return
	}

	var student models.Student
	if err := database.DB.Where("nisn = ?", input.NISN).First(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "NISN tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(input.Password)); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Password salah")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id": student.ID,
		"role":       "siswa",
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login berhasil", gin.H{
		"token": tokenString,
	})
}
