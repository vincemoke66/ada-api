package recordRoutes

import (
	"github.com/gofiber/fiber/v2"
	recordHandler "github.com/vincemoke66/ada-api/internals/handlers/record"
)

func SetupStudentRoutes(router fiber.Router) {
	record := router.Group("/record")

	// Create a record
	record.Post("/", recordHandler.CreateRecord)

	// Read all records
	record.Get("/", recordHandler.GetAllRecords)
}
