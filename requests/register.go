package requests

type RegisterGuruRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	NIK         string `json:"nik" binding:"required,len=16,numeric"`
	NUPTK       string `json:"nuptk"`
	SchoolName  string `json:"school_name"`
	SchoolLevel string `json:"school_level" binding:"required,oneof=SD SMP SMA"`
	IsSLB       bool   `json:"is_slb"`
}