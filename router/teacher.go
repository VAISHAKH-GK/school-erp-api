package router

import (
	"github.com/FulgurCode/school-erp-api/controller"
	"github.com/gin-gonic/gin"
)

func TeacherRouter(router *gin.RouterGroup) {
	// Authentication
	// signup route
	router.POST("/signup", controller.TeacherSignup)
	// signup-otp route
	router.GET("/signup-otp", controller.TeacherSignupOTP)
	// login route
	router.POST("/login", controller.TeacherLogin)
	// checklogin route
	router.GET("/checklogin", controller.TeacherCheckLogin)
	// logout route
	router.DELETE("/logout", controller.TeacherLogout)

	// Admission
	// new-admission route
	router.POST("/new-admission", controller.TeacherNewAdmissionRoute)
	// edit-student route
	router.PUT("/edit-student", controller.TeacherEditStudent)
	// upload-student-photo route
	router.POST("/upload-student-photo", controller.TeacherUploadStudentPhoto)
	// import-students route
	router.POST("/import-students", controller.TeacherImportStudents)
	// get-admitted-students route
	router.GET("/get-admitted-students", controller.TeacherGetAdmittedStudents)
	// verify-student route
	router.PATCH("/verify-student", controller.TeacherVerifyStudent)
	// students-to-verify route
	router.GET("/students-to-verify", controller.TeacherStudentsToVerify)
	// course-language-report route
	router.GET("/course-language-report", controller.TeacherCourseLanguageReport)
	// course-status-report
	router.GET("/course-status-report", controller.TeacherCourseStatusReport)
	// course-gender-report
	router.GET("/course-gender-report", controller.TeacherCourseGenderReport)
	// course-category-report
	router.GET("/course-category-report", controller.TeacherCourseCategoryReport)
	// course-caste-report
	router.GET("/course-caste-report", controller.TeacherCourseCasteReport)

	// Students
	// get-students route
	router.GET("/get-students", controller.TeacherGetStudentsRoute)
	// get-student route
	router.GET("/get-student", controller.TeacherGetStudent)
	// get-student-photo route
	router.GET("/get-student-photo", controller.TeacherGetStudentPhoto)

	// Permission
	router.GET("/have-duty", controller.TeacherHaveDuty)
}
