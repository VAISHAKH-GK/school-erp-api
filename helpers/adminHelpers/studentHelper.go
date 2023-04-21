package adminHelpers

import (
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
)

// Getting students according to search type
func GetStudents(search string, value string) ([]map[string]interface{}, error) {
	switch search {
	case "admissionNo":
		// Getting student by admission number
		var students, err = databaseHelpers.GetStudentByAdmissionNo(value)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "applicationNo":
		// Getting student by application number
		var students, err = databaseHelpers.GetStudentByApplicationNo(value)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "name":
		// Getting student by name
		var students, err = databaseHelpers.GetStudentByName(value)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	default:
		// Sending empty array if search type is random
		return []map[string]interface{}{}, nil
	}
}
