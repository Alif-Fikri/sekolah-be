package requests

type ClassRequest struct {
	NameKelas    string `json:"name_kelas" binding:"required"`
	ClassLevel   string `json:"class_level" binding:"required"`
	SchoolLevel  string `json:"school_level" binding:"required,oneof=SD SMP SMA"`
}
