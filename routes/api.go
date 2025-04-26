package routes

import (
	"github.com/gin-gonic/gin"
	"sekolah-be/controllers"
)

func Api(r *gin.Engine) {
	guru := r.Group("/guru")
	{
		guru.POST("/register", controllers.RegisterTeacher)
		guru.POST("/login", controllers.LoginTeacher)
	}
}
