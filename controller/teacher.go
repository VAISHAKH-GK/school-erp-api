package controller

import (
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/teacherHelpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// POST request on '/api/teacher/signup'
func TeacherSignup(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	var teacher, err = databaseHelpers.GetTeacherWithEmail(data["email"].(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(401, "No teacher with this email")
			return
		}
		c.JSON(500, "Request failed")
		return
	}
	var exists = teacherHelpers.UserExists(teacher)
	if !exists {
		c.JSON(409, "Account already made")
		return
	}
	err = teacherHelpers.SignUpSetOTP(c, data)
	if err != nil {
		c.JSON(500, "Network issue")
		return
	}
	c.JSON(200, "OTP sended to the email adress")
}

// GET request on '/api/teacher/signup-otp'
func TeacherSignupOTP(c *gin.Context) {
	// comparing otp
	var result = teacherHelpers.CompareOtp(c)
	// Checking if password is correct and sending response
	if !result {
		c.JSON(401, "Incorrect OTP")
		return
	}
	teacherHelpers.CreateTeacherUser(c)
	c.JSON(200, "Teacher account created")
}

// POST request '/api/teacher/login'
func TeacherLogin(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Checking for username
	var teacher, err = databaseHelpers.GetTeacherWithEmail(data["email"].(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(401, "Wrong username or password")
			return
		}
		c.JSON(500, "Request failed")
	}
	// Comparing password and sending response
	var result = helpers.ComparePassword(teacher["password"].(string), data["password"].(string))
	if !result {
		c.JSON(401, "Wrong username or password")
		return
	}
	// storing id and sending response if password is correct
	teacherHelpers.LoginWithSesssion(c, teacher)
	c.JSON(200, "Login Successful")
}

// GET request on '/api/teacher/checklogin'
func TeacherCheckLogin(c *gin.Context) {
	// checking if logged in as teacher and sending response
	var isLoggedIn = teacherHelpers.CheckLogin(c)
	c.JSON(200, isLoggedIn)
}

// DELETE request on '/api/teacher/logout'
func TeacherLogout(c *gin.Context) {
	// clearing 'teacher' session
	teacherHelpers.Logout(c)
	// Response for the request
	c.JSON(200, "Loggged Out")
}

// GET request on '/api/teacher/get-admitted-students'
func TeacherGetAdmittedStudents(c *gin.Context) {
	// Checking if logged in
	if !teacherHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In as teacher")
		return
	}
	// Getting admitted student details and sending response
	var students, err = databaseHelpers.GetAdmittedStudents()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, students)
}

// PATCH request on '/api/teacher/verify-student'
func TeacherVerifyStudent(c *gin.Context) {
	// Checking if logged in
	if !teacherHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In as teacher")
		return
	}
	// Getting student id
	var studentId, _ = primitive.ObjectIDFromHex(c.Query("studentId"))
	// verifying student and sending response
	var err = databaseHelpers.VerifyStudent(studentId)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Student verifyed")
}

// GET request on '/api/teacher/students-to-verify'
func TeacherStudentsToVerify(c *gin.Context) {
	// Checking if logged in
	if !teacherHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In as teacher")
		return
	}
	// Getting students that is remain to verify and sending response
	var students, err = databaseHelpers.GetStudentsToVerify()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, students)
}

// GET request on '/api/teacher/get-student'
func TeacherGetStudent(c *gin.Context) {
	// Checking if logged in
	if !teacherHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In as teacher")
		return
	}
	// Getting object id of student
	var studentId, err = primitive.ObjectIDFromHex(c.Query("studentId"))
	helpers.CheckNilErr(err)
	// Getting student using id
	student, err := databaseHelpers.GetStudent(studentId)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, student)
}
