package controllers

import (
	"encoding/csv"
	"net/http"
	"sekolah-be/database"
	"sekolah-be/models"
	"sekolah-be/requests"
	"sekolah-be/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func ExportAttendanceCSV(c *gin.Context) {
	var req requests.AttendanceExportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Data tidak valid")
		return
	}

	var start, end time.Time
	location := time.Now().Location()

	switch req.RangeType {
	case "harian":
		date, err := time.ParseInLocation("2006-01-02", req.Date, location)
		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Tanggal tidak valid")
			return
		}
		start = date
		end = date.Add(24 * time.Hour)

	case "mingguan":
		end = time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour)
		start = end.AddDate(0, 0, -7)

	case "custom":
		var err error
		start, err = time.ParseInLocation("2006-01-02", req.StartDate, location)
		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Start date tidak valid")
			return
		}
		end, err = time.ParseInLocation("2006-01-02", req.EndDate, location)
		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "End date tidak valid")
			return
		}
		end = end.Add(24 * time.Hour)

	case "bulanan":
		var err error
		start, err = time.ParseInLocation("2006-01-02", req.StartDate, location)
		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Start date tidak valid")
			return
		}
		end = start.AddDate(0, 1, 0)

	default:
		utils.ErrorResponse(c, http.StatusBadRequest, "Range type tidak dikenali")
		return
	}

	var attendances []models.SubjectAttendance
	if err := database.DB.Preload("Student").Preload("Subject").
		Where("subject_id = ? AND attendance_time BETWEEN ? AND ?", req.SubjectID, start, end).
		Find(&attendances).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data absensi")
		return
	}

	var class models.Class
	if err := database.DB.First(&class, req.ClassID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Kelas tidak ditemukan")
		return
	}

	c.Header("Content-Disposition", "attachment; filename=rekap_absensi.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	writer.Write([]string{"Nama Siswa", "Tanggal", "Jam", "Status", "Kelas", "Mata Pelajaran"})

	for _, a := range attendances {
		writer.Write([]string{
			a.Student.Name,
			a.AttendanceTime.Format("2006-01-02"),
			a.AttendanceTime.Format("15:04"),
			a.Status,
			class.NameKelas,
			a.Subject.MataPelajaran,
		})
	}
}
