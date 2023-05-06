package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Getting student details from database usign query
func GetStudents(query map[string]interface{}) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting students from database
	var result, err = db.Collection("students").Find(context.Background(), query)
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

func GetStudentsByName(name string, status string) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting students from database
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"$text": bson.M{"$search": name}})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		if status != "" && student["status"] == status {
			students = append(students, student)
		} else {
			students = append(students, student)
		}
	}
	return students, err
}

// Getting student details using student id
func GetStudent(studentId primitive.ObjectID) (map[string]interface{}, error) {
	// database
	var db = connections.Db
	// get student and return
	var student map[string]interface{}
	var err = db.Collection("students").FindOne(context.Background(), bson.M{"_id": studentId}).Decode(&student)
	return student, err
}
