package buildingHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/ada-api/database"
	"github.com/vincemoke66/ada-api/internals/model"
)

// GetBuildings func gets all existing buildings
// @Description Get all existing buildings
// @Tags Building
// @Accept json
// @Produce json
// @Success 200 {array} model.Building
// @router /api/building [get]
func GetBuildings(c *fiber.Ctx) error {
	db := database.DB
	var building []model.Building

	// find all buildings in the database
	db.Find(&building)

	// If no building is present return an error
	if len(building) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Buildings data found", "data": nil})
	}

	// Return buildings
	return c.JSON(fiber.Map{"status": "success", "message": "Buildings Found", "data": building})
}

// CreateBuilding func create a building
// @Description Create a Building
// @Tags Building
// @Accept json
// @Produce json
// @Param name body string true "name"
// @Param abbrv body string true "abbrv"
// @Success 200 {object} model.Building
// @router /api/building [post]
func CreateBuilding(c *fiber.Ctx) error {
	db := database.DB
	building := new(model.Building)

	// Parse the body to the building object
	err := c.BodyParser(building)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Return invalid name if empty or null
	if building.Name == uuid.Nil.String() || building.Name == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Building Name", "data": err})
	}

	// Create a temporary building data
	var storedBuilding model.Building

	// Find the building with the given name param
	db.Find(&storedBuilding, "name = ?", building.Name)

	// If building name exists, return an error
	if storedBuilding.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Building with the same name already exist.", "data": nil})
	}

	// Add a uuid to the new building
	building.ID = uuid.New()

	// Create the Building
	err = db.Create(&building).Error
	// Return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create building", "data": err})
	}

	// Return the created building
	return c.JSON(fiber.Map{"status": "success", "message": "Building created", "data": building})
}

// GetBuilding func get one building by name
// @Description Get one building by name
// @Tags Building
// @Accept json
// @Produce json
// @Success 200 {object} model.Building
// @router /api/building/{name} [get]
func GetBuilding(c *fiber.Ctx) error {
	db := database.DB
	var building model.Building

	// Read the param building_name
	building_name := c.Params("name")

	// Find the building with the given name
	db.Find(&building, "name = ?", building_name)

	// If no such building present, return an error
	if building.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Building not found", "data": nil})
	}

	// Return the building with the specified name
	return c.JSON(fiber.Map{"status": "success", "message": "Building Found", "data": building})
}

// UpdateBuilding update an building by name
// @Description Update a Building by name
// @Tags Building
// @Accept json
// @Produce json
// @Param name body string true "name"
// @Param abbrv body string true "abbrv"
// @Success 200 {object} model.Building
// @router /api/building/{name} [put]
func UpdateBuilding(c *fiber.Ctx) error {
	// Create a struct for updating only writable values
	type updateBuilding struct {
		Name  string `json:"name"`
		Abbrv string `json:"abbrv"`
	}

	db := database.DB
	var building model.Building

	// Read the building name param
	building_name := c.Params("name")

	// Find the building with the given school_id param
	db.Find(&building, "name = ?", building_name)

	// If no such building, return an error
	if building.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Building not found", "data": nil})
	}

	// Store the body containing the updated data
	var updateBuildingData updateBuilding
	err := c.BodyParser(&updateBuildingData)

	// Return parsing error if encountered
	if err != nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the building
	building.Name = updateBuildingData.Name
	building.Abbrv = updateBuildingData.Abbrv

	// Save the Changes
	db.Save(&building)

	// Return the updated building
	return c.JSON(fiber.Map{"status": "success", "message": "Building Updated", "data": building})
}

// DeleteBuilding delete a building by name
// @Description Delete a Building by name
// @Tags Building
// @Accept json
// @Produce json
// @Success 200
// @router /api/building/{name} [delete]
func DeleteBuilding(c *fiber.Ctx) error {
	db := database.DB
	var building model.Building

	// Read the building name param
	building_name := c.Params("name")

	// Find the building with the given name param
	db.Find(&building, "name = ?", building_name)

	// If no such building present return an error
	if building.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Building not found", "data": nil})
	}

	// Delete the building
	err := db.Delete(&building, "name = ?", building_name).Error

	// Return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete building", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Building Deleted"})
}
