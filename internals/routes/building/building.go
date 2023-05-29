package buildingRoutes

import (
	"github.com/gofiber/fiber/v2"
	buildingHandler "github.com/vincemoke66/ada-api/internals/handlers/building"
)

func SetupStudentRoutes(router fiber.Router) {
	building := router.Group("/building")

	// Create a building
	building.Post("/", buildingHandler.CreateBuilding)
	// Read all building
	building.Get("/", buildingHandler.GetBuildings)
	// Read a building
	building.Get("/:name", buildingHandler.GetBuilding)
	// Update building
	building.Put("/:name", buildingHandler.UpdateBuilding)
	// Delete a building
	building.Delete("/:name", buildingHandler.DeleteBuilding)
}
