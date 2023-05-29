package instructorRoutes

import (
	"github.com/gofiber/fiber/v2"
	instructorHandler "github.com/vincemoke66/ada-api/internals/handlers/instructor"
)

func SetupStudentRoutes(router fiber.Router) {
	instructor := router.Group("/instructor")

	// Create a instructor
	instructor.Post("/", instructorHandler.CreateInstructor)
	// Read all instructor
	instructor.Get("/", instructorHandler.GetInstructors)
	// Read a instructor
	instructor.Get("/:school_id", instructorHandler.GetInstructor)
	// Update instructor
	instructor.Put("/:school_id", instructorHandler.UpdateInstructor)
	// Delete a instructor
	instructor.Delete("/:school_id", instructorHandler.DeleteInstructor)
}
