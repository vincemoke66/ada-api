package instructorHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/ada-api/database"
	"github.com/vincemoke66/ada-api/internals/model"
)

// GetInstructors func gets all existing instructors
// @Description Get all existing instructors
// @Tags Instructor
// @Accept json
// @Produce json
// @Success 200 {array} model.Instructor
// @router /api/instructor [get]
func GetInstructors(c *fiber.Ctx) error {
	db := database.DB
	var instructor []model.Instructor

	// find all instructors in the database
	db.Find(&instructor)

	// If no instructor is present return an error
	if len(instructor) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Instructors data found", "data": nil})
	}

	// Return instructors
	return c.JSON(fiber.Map{"status": "success", "message": "Instructors Found", "data": instructor})
}

// CreateInstructor func create a instructor
// @Description Create a Instructor
// @Tags Instructor
// @Accept json
// @Produce json
// @Param first_name body string true "first_name"
// @Param last_name body string true "last_name"
// @Param school_id body string true "school_id"
// @Success 200 {object} model.Instructor
// @router /api/instructor [post]
func CreateInstructor(c *fiber.Ctx) error {
	db := database.DB
	instructor := new(model.Instructor)

	// Parse the body to the instructor object
	err := c.BodyParser(instructor)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Return invalid school_id if empty or null
	if instructor.SchoolID == uuid.Nil.String() || instructor.SchoolID == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid School ID", "data": err})
	}

	// Create a temporary instructor data
	var storedInstructor model.Instructor

	// Find the instructor with the given school_id
	db.Find(&storedInstructor, "school_id = ?", instructor.SchoolID)

	// If instructor school id exists, return an error
	if storedInstructor.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Instructor with the same school id already exist.", "data": nil})
	}

	// Add a uuid to the new instructor
	instructor.ID = uuid.New()

	// Create the Instructor and return error if encountered
	err = db.Create(&instructor).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create instructor", "data": err})
	}

	// Return the created instructor
	return c.JSON(fiber.Map{"status": "success", "message": "Instructor created", "data": instructor})
}

// GetInstructor func get one instructor by school_id
// @Description Get one instructor by school_id
// @Tags Instructor
// @Accept json
// @Produce json
// @Success 200 {object} model.Instructor
// @router /api/instructor/{school_id} [get]
func GetInstructor(c *fiber.Ctx) error {
	db := database.DB
	var instructor model.Instructor

	// Read the param school_id
	school_id := c.Params("school_id")

	// Find the instructor with the given school_id
	db.Find(&instructor, "school_id = ?", school_id)

	// If no such instructor present, return an error
	if instructor.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Instructor not found", "data": nil})
	}

	// Return the instructor with the specified school_id
	return c.JSON(fiber.Map{"status": "success", "message": "Instructor Found", "data": instructor})
}

// UpdateInstructor update an instructor by school_id
// @Description Update a Instructor by school_id
// @Tags Instructor
// @Accept json
// @Produce json
// @Param first_name body string true "first_name"
// @Param last_name body string true "last_name"
// @Success 200 {object} model.Instructor
// @router /api/instructor/{school_id} [put]
func UpdateInstructor(c *fiber.Ctx) error {
	// Create a struct for updating only writable values
	type updateInstructor struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	db := database.DB
	var instructor model.Instructor

	// Read the param school_id
	school_id := c.Params("school_id")

	// Find the instructor with the given school_id param
	db.Find(&instructor, "school_id = ?", school_id)

	// If no such instructor, return an error
	if instructor.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Instructor not found", "data": nil})
	}

	// Store the body containing the updated data
	var updateStudentData updateInstructor
	err := c.BodyParser(&updateStudentData)

	// Return parsing error if encountered
	if err != nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the instructor
	instructor.FirstName = updateStudentData.FirstName
	instructor.LastName = updateStudentData.LastName

	// Save the Changes
	db.Save(&instructor)

	// Return the updated instructor
	return c.JSON(fiber.Map{"status": "success", "message": "Instructor Updated", "data": instructor})
}

// DeleteInstructor delete an instructor by school_id
// @Description Delete a Instructor by school_id
// @Tags Instructor
// @Accept json
// @Produce json
// @Success 200
// @router /api/instructor/{school_id} [delete]
func DeleteInstructor(c *fiber.Ctx) error {
	db := database.DB
	var instructor model.Instructor

	// Read the param school_id
	school_id := c.Params("school_id")

	// Find the instructor with the given school_id param
	db.Find(&instructor, "school_id = ?", school_id)

	// If no such instructor present return an error
	if instructor.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Instructor not found", "data": nil})
	}

	// Delete the instructor
	err := db.Delete(&instructor, "school_id = ?", school_id).Error

	// Return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete instructor", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Instructor Deleted"})
}
