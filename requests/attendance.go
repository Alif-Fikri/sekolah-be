package requests

type ClassAttendanceRequest struct {
	StudentID uint   `json:"student_id" binding:"required"`
	ClassID   uint   `json:"class_id" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=hadir izin sakit alpha"`
}

type SubjectAttendanceRequest struct {
	StudentID uint   `json:"student_id" binding:"required"`
	SubjectID uint   `json:"subject_id" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=hadir izin sakit alpha"`
}

type AttendanceExportRequest struct {
	ClassID   uint   `json:"class_id" binding:"required"`
	SubjectID uint   `json:"subject_id" binding:"required"`
	RangeType string `json:"range_type" binding:"required"`
	Date      string `json:"date"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
