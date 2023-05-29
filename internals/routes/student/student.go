package studentRoutes

import (
	"github.com/gofiber/fiber/v2"
	studentHandler "github.com/vincemoke66/ada-api/internals/handlers/student"
)

func SetupStudentRoutes(router fiber.Router) {
	student := router.Group("/student")

	// Create a student
	student.Post("/", studentHandler.CreateStudent)
	// Read all students
	student.Get("/", studentHandler.GetStudents)
	// Read a student through rfid
	student.Get("/:rfid", studentHandler.GetStudentThroughRFID)
	// Read a student through school id
	student.Get("/:school_id", studentHandler.GetStudent)
	// Update student
	student.Put("/:school_id", studentHandler.UpdateStudent)
	// Delete a student
	student.Delete("/:school_id", studentHandler.DeleteStudent)
}
