package requests

type LoginStudentRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterStudentRequest struct {
	Name        string `json:"name" binding:"required"`
	NISN        string `json:"nisn" binding:"required"`
	SchoolLevel string `json:"school_level" binding:"required,oneof=SD SMP SMA"`
	IsSLB       bool   `json:"is_slb"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date" binding:"required"` // Format: YYYY-MM-DD
	Gender      string `json:"gender" binding:"required,oneof=L P"`
	Address     string `json:"address"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone"`
}