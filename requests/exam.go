package requests

import "encoding/json"

type CreateExamRequest struct {
	Name       string `json:"name" binding:"required"`
	SubjectID  uint   `json:"subject_id" binding:"required"`
	ClassID    uint   `json:"class_id" binding:"required"`
	TypeSeal   string `json:"type_seal" binding:"required,oneof=Wide Paper"`
	UploadType string `json:"upload_type" binding:"required,oneof=manual event"`
	Token      string `json:"token" binding:"required"`
}

type CreateExamResultRequest struct {
	StudentID  uint            `json:"student_id" binding:"required"`
	ExamID     uint            `json:"exam_id" binding:"required"`
	Score      float64         `json:"score" binding:"required"`
	Confidence float64         `json:"confidence" binding:"required"`
	Answer     json.RawMessage `json:"answer" binding:"required"`
}

type UpdateExamRequest struct {
	Name       string `json:"name"`
	SubjectID  uint   `json:"subject_id"`
	ClassID    uint   `json:"class_id"`
	TypeSeal   string `json:"type_seal" binding:"omitempty,oneof=Wide Paper"`
	UploadType string `json:"upload_type" binding:"omitempty,oneof=manual event"`
	Token      string `json:"token"`
}
