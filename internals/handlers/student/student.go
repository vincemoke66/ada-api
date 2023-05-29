package studentHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/ada-api/database"
	"github.com/vincemoke66/ada-api/internals/model"
)

// GetStudents func gets all existing students
// @Description Get all existing students
// @Tags Student
// @Accept json
// @Produce json
// @Success 200 {array} model.Student
// @router /api/student [get]
func GetStudents(c *fiber.Ctx) error {
	db := database.DB
	var students []model.Student

	// find all students in the database
	db.Find(&students)

	// If no student is present return an error
	if len(students) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Students data found", "data": nil})
	}

	// Else return students
	return c.JSON(fiber.Map{"status": "success", "message": "Students Found", "data": students})
}

// CreateStudent func create a student
// @Description Create a Student
// @Tags Student
// @Accept json
// @Produce json
// @Param first_name body string true "first_name"
// @Param last_name body string true "last_name"
// @Param school_id body string true "school_id"
// @Param college body string true "college"
// @Param rfid body string true "rfid"
// @Param course body string true "course"
// @Success 200 {object} model.Student
// @router /api/student [post]
func CreateStudent(c *fiber.Ctx) error {
	db := database.DB
	student := new(model.Student)

	// Parse the body to the student object
	err := c.BodyParser(student)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Return invalid school_id if empty or null
	if student.SchoolID == uuid.Nil.String() || student.SchoolID == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid School ID", "data": err})
	}

	// Create a temporary student data
	var storedStudent model.Student
	// Find the student with the given school_id
	db.Find(&storedStudent, "school_id = ?", student.SchoolID)
	// If student school id exists, return an error
	if storedStudent.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Invalid Student Credentials.", "data": nil})
	}
	// Find the student with the given rfid
	db.Find(&storedStudent, "rfid = ?", student.RFID)
	// If student rfid exists, return an error
	if storedStudent.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Invalid Student Credentials.", "data": nil})
	}

	// Add a uuid to the new student
	student.ID = uuid.New()

	// Create the Student and return error if encountered
	err = db.Create(&student).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create student", "data": err})
	}

	// Return the created student
	return c.JSON(fiber.Map{"status": "success", "message": "Student created", "data": student})
}

// GetStudent func get one student by school_id
// @Description Get one student by school_id
// @Tags Student
// @Accept json
// @Produce json
// @Success 200 {object} model.Student
// @router /api/student/{school_id} [get]
func GetStudent(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student

	// Read the param school_id
	school_id := c.Params("school_id")

	// Find the student with the given school_id
	db.Find(&student, "school_id = ?", school_id)

	// If no such student present, return an error
	if student.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Student not found", "data": nil})
	}

	// Return the student with the specified school_id
	return c.JSON(fiber.Map{"status": "success", "message": "Student Found", "data": student})
}

// GetStudentThroughRFID func get one student by rfid
// @Description Get one student by rfid
// @Tags Student
// @Accept json
// @Produce json
// @Success 200 {object} model.Student
// @router /api/student/{rfid} [get]
func GetStudentThroughRFID(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student

	// Read the param rfid
	rfid := c.Params("rfid")

	// Find the student with the given rfid
	db.Find(&student, "rfid = ?", rfid)

	// If no such student present, return an error
	if student.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Student not found with that rfid", "data": nil})
	}

	// Return the student with the specified rfid
	return c.JSON(fiber.Map{"status": "success", "message": "Student Found", "data": student})
}

// UpdateStudent update a student by school_id
// @Description Update a Student by school_id
// @Tags Student
// @Accept json
// @Produce json
// @Param first_name body string true "first_name"
// @Param last_name body string true "last_name"
// @Param college body string true "college"
// @Param course body string true "course"
// @Success 200 {object} model.Student
// @router /api/student/{school_id} [put]
func UpdateStudent(c *fiber.Ctx) error {
	// Create a struct for updating only writable values
	type updateStudent struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		College   string `json:"college"`
		Course    string `json:"course"`
	}

	db := database.DB
	var student model.Student

	// Read the param school_id
	school_id := c.Params("school_id")

	// Find the student with the given school_id param
	db.Find(&student, "school_id = ?", school_id)

	// If no such student, return an error
	if student.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Student not found", "data": nil})
	}

	// Store the body containing the updated data
	var updateStudentData updateStudent
	err := c.BodyParser(&updateStudentData)

	// Return parsing error if encountered
	if err != nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the student
	student.FirstName = updateStudentData.FirstName
	student.LastName = updateStudentData.LastName
	student.College = updateStudentData.College
	student.Course = updateStudentData.Course

	// Save the Changes
	db.Save(&student)

	// Return the updated student
	return c.JSON(fiber.Map{"status": "success", "message": "Student Updated", "data": student})
}

// DeleteStudent delete a student by school_id
// @Description Delete a Student by school_id
// @Tags Student
// @Accept json
// @Produce json
// @Success 200
// @router /api/student/{school_id} [delete]
func DeleteStudent(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student

	// Read the param school_id
	school_id := c.Params("school_id")

	// Find the student with the given school_id param
	db.Find(&student, "school_id = ?", school_id)

	// If no such student present return an error
	if student.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Student not found", "data": nil})
	}

	// Delete the student
	err := db.Delete(&student, "school_id = ?", school_id).Error

	// Return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete student", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Student Deleted"})
}
