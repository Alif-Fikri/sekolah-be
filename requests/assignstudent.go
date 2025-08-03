package requests

type AssignStudentToClassRequest struct {
	StudentIDs []uint `json:"student_ids"`
	ClassID    uint   `json:"class_id"`
}

type AssignStudentToSubjectRequest struct {
	StudentIDs []uint `json:"student_ids"`
	SubjectID  uint   `json:"subject_id"`
}
