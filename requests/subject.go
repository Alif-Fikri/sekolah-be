package requests

type SubjectRequest struct {
	MataPelajaran string `json:"mata_pelajaran" binding:"required"`
	SchoolLevel   string `json:"school_level" binding:"required,oneof=SD SMP SMA"`
}
