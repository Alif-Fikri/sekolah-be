package routes

import (
	classController "sekolah-be/controllers/class"
	studentController "sekolah-be/controllers/student"
	subjectController "sekolah-be/controllers/subject"
	teacherController "sekolah-be/controllers/teacher"
	"sekolah-be/middlewares"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	guru := r.Group("/guru")
	{
		guru.POST("/register", teacherController.RegisterTeacher)
		guru.POST("/login", teacherController.LoginTeacher)
		guru.POST("/logout", middlewares.AuthMiddleware(), teacherController.LogoutTeacher)
		guru.GET("/:id", middlewares.AuthMiddleware(), teacherController.GetTeacherByID)
		guru.GET("/all", middlewares.AuthMiddleware(), teacherController.GetAllTeachers)
	}

	subject := r.Group("/subject", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("guru"))
	{
		subject.POST("/", subjectController.CreateSubject)
		subject.GET("/", subjectController.GetAllSubjects)
		subject.GET("/:id", subjectController.GetSubjectByID)
		subject.PUT("/:id", subjectController.UpdateSubject)
		subject.DELETE("/:id", subjectController.DeleteSubject)
	}

	class := r.Group("/class", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("guru"))
	{
		class.POST("/", classController.CreateClass)
		class.GET("/", classController.GetAllClasses)
		class.GET("/:id", classController.GetClassByID)
		class.PUT("/:id", classController.UpdateClass)
		class.DELETE("/:id", classController.DeleteClass)
	}

	student := r.Group("/student")
	{
		student.POST("/login", studentController.LoginStudent)
		student.POST("/logout", studentController.LogoutStudent)

		studentCreate := student.Group("/", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("guru"))
		{
			studentCreate.POST("/register", studentController.RegisterStudent)
			studentCreate.POST("/import", studentController.ImportStudents)
		}
	}
}
