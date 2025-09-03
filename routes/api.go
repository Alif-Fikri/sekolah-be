package routes

import (
	attendanceController "sekolah-be/controllers/attendance"
	authController "sekolah-be/controllers/auth"
	classController "sekolah-be/controllers/class"
	studentController "sekolah-be/controllers/student"
	subjectController "sekolah-be/controllers/subject"
	teacherController "sekolah-be/controllers/teacher"
	"sekolah-be/middlewares"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {

	api := r.Group("/api")
	api.POST("/login", authController.Login)
	api.POST("/logout", authController.Logout)

	guru := api.Group("/guru")
	{
		guru.POST("/register", teacherController.RegisterTeacher)
		guru.GET("/:id", middlewares.AuthMiddleware(), teacherController.GetTeacherByID)
		guru.GET("/all", middlewares.AuthMiddleware(), teacherController.GetAllTeachers)
	}

	subject := api.Group("/subject", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("guru"))
	{
		subject.POST("/", subjectController.CreateSubject)
		subject.GET("/", subjectController.GetAllSubjects)
		subject.GET("/:id", subjectController.GetSubjectByID)
		subject.PUT("/:id", subjectController.UpdateSubject)
		subject.DELETE("/:id", subjectController.DeleteSubject)
		subject.POST("/assign-student", subjectController.AssignStudentsToSubject)
	}

	class := api.Group("/class", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("guru"))
	{
		class.POST("/", classController.CreateClass)
		class.GET("/", classController.GetAllClasses)
		class.GET("/:id", classController.GetClassByID)
		class.PUT("/:id", classController.UpdateClass)
		class.DELETE("/:id", classController.DeleteClass)
		class.POST("/assign-student", classController.AssignStudentsToClass)
	}

	student := api.Group("/student")
	{
		studentCreate := student.Group("/", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("guru"))
		{
			studentCreate.POST("/register", studentController.RegisterStudent)
			studentCreate.POST("/import", studentController.ImportStudents)
		}
	}

	attendance := api.Group("/attendance")
	{
		attendance.POST("/class", attendanceController.CreateClassAttendance)
		attendance.POST("/subject", attendanceController.CreateSubjectAttendance)
		attendance.POST("/exportcsv", attendanceController.ExportAttendanceCSV)
	}
}
