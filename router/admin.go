package router

import (
	"github.com/FulgurCode/school-erp-api/controller"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.RouterGroup) {
	// Authentication
	// login route
	router.POST("/login", controller.AdminLoginRoute)
	// checklogin route
	router.GET("/checklogin", controller.AdminCheckLoginRoute)
	// logout route
	router.DELETE("/logout", controller.AdminLogoutRoute)
	// change-password route
	router.PUT("/change-password", controller.ChangeAdminPassword)

	// Admission
	// new-admission route
	router.POST("/new-admission", controller.NewAdmissionRoute)
	// import-students route
	router.POST("/import-students", controller.ImportStudents)
	// edit-student route
	router.PUT("/edit-student", controller.EditStudent)
	// upload-student-photo route
	router.POST("/upload-student-photo", controller.UploadStudentPhoto)
	// get-student-photo route
	router.GET("/get-student-photo", controller.GetStudentPhoto)
	// students-to-confirm route
	router.GET("/students-to-confirm", controller.AdminStudentsToConfirm)
	// confirm-student route
	router.PATCH("/confirm-student", controller.AdminConfirmStudent)
	// verify-student route
	router.PATCH("/verify-student", controller.AdminVerifyStudent)
	// students-to-verify route
	router.GET("/students-to-verify", controller.AdminStudentsToVerify)
	// course-language-report route
	router.GET("/course-language-report", controller.AdminCourseLanguageReport)
	// course-status-report
	router.GET("/course-status-report", controller.AdminCourseStatusReport)
	// course-gender-report
	router.GET("/course-gender-report", controller.AdminCourseGenderReport)
	// course-category-report
	router.GET("/course-category-report", controller.AdminCourseCategoryReport)
	// course-caste-report
	router.GET("/course-caste-report", controller.AdminCourseCasteReport)

	// Student
	// get-student route
	router.GET("/get-students", controller.GetStudentsRoute)
	// get-student route
	router.GET("/get-student", controller.GetStudent)

	// Teacher
	// add-teacher route
	router.POST("/add-teacher", controller.AddTeacher)
	// add-duty route
	router.POST("/add-duty", controller.AddDuty)
	// get-teacher route
	router.GET("/get-teacher", controller.GetTeacher)
	// import-teachers route
	router.POST("/import-teachers", controller.ImportTeachers)
	// get-teachers route
	router.GET("/get-teachers", controller.GetTeachers)
	// get-duties route
	router.GET("/get-duties", controller.GetDuties)
	// delete-duty route
	router.DELETE("/delete-duty", controller.DeleteDuty)
}
