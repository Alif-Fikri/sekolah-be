package Controllers

import (
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterStudent(c *gin.Context) {
	var input requests.RegisterStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "data siswa tidak valid: "+err.Error())
		return
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "format tanggal lahir tidak valid (YYYY-MM-DD)")
		return
	}

	passwordRaw := birthDate.Format("02012006")
	passwordHash, err := utils.HashPassword(passwordRaw)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal generate password")
		return
	}

	user := c.MustGet("user").(jwt.MapClaims)
	teacherID := uint(user["user_id"].(float64))

	student := models.Student{
		Name:        input.Name,
		NISN:        input.NISN,
		Password:    passwordHash,
		SchoolLevel: input.SchoolLevel,
		IsSLB:       input.IsSLB,
		BirthPlace:  input.BirthPlace,
		BirthDate:   birthDate,
		Gender:      input.Gender,
		Address:     input.Address,
		Email:       input.Email,
		Phone:       input.Phone,
		CreatedBy:   &teacherID,
	}

	if err := database.DB.Create(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "gagal mendaftarkan siswa: "+err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "siswa berhasil didaftarkan", gin.H{
		"password_awal": passwordRaw,
	})
}

func ImportStudents(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "File CSV wajib diunggah")
		return
	}

	src, err := file.Open()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuka file")
		return
	}
	defer src.Close()

	reader := csv.NewReader(src)
	rows, err := reader.ReadAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Format CSV tidak valid")
		return
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 10 {
			continue
		}

		birthDate, err := time.Parse("2006-01-02", row[5])
		if err != nil {
			continue
		}

		passwordRaw := birthDate.Format("02012006")
		passwordHash, err := utils.HashPassword(passwordRaw)
		if err != nil {
			continue
		}

		user := c.MustGet("user").(jwt.MapClaims)
		teacherID := uint(user["user_id"].(float64))

		isSLB, _ := strconv.ParseBool(row[3])

		student := models.Student{
			Name:        row[0],
			NISN:        row[1],
			Password:    passwordHash,
			SchoolLevel: row[2],
			IsSLB:       isSLB,
			BirthPlace:  row[4],
			BirthDate:   birthDate,
			Gender:      row[6],
			Address:     row[7],
			Email:       row[8],
			Phone:       row[9],
			CreatedBy: &teacherID,
		}

		database.DB.Create(&student)
	}

	utils.SuccessResponse(c, http.StatusOK, "Data siswa berhasil diimpor", nil)
}
