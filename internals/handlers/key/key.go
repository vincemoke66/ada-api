package keyHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/ada-api/database"
	"github.com/vincemoke66/ada-api/internals/model"
)

// GetKeys func gets all existing keys
// @Description Get all existing keys
// @Tags Key
// @Accept json
// @Produce json
// @Success 200 {array} model.Key
// @router /api/key [get]
func GetKeys(c *fiber.Ctx) error {
	db := database.DB
	var keys []model.Key

	// find all keys in the database
	db.Find(&keys)

	// If no key is present return an error
	if len(keys) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Keys data found", "data": nil})
	}

	// Else return keys
	return c.JSON(fiber.Map{"status": "success", "message": "Keys Found", "data": keys})
}

// GetKeysUsingBuildingName func gets all keys based on the given building name
// @Description Gets all keys based on the given building name
// @Tags Key
// @Accept json
// @Produce json
// @Success 200 {array} model.Key
// @router /api/key/{building_name} [get]
func GetKeysUsingBuildingName(c *fiber.Ctx) error {
	db := database.DB
	var keys []model.Key

	// Read the param key_building_name
	key_building_name := c.Params("building_name")

	// Create a temporary building data
	var storedBuilding model.Building
	// Find the building with the given building name
	db.Find(&storedBuilding, "name = ?", key_building_name)
	// If building does not exist, return an error
	if storedBuilding.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Building does not exist.", "data": nil})
	}

	// find all keys on the given building name in the database
	db.Find(&keys, "building_id = ?", storedBuilding.ID)

	// If no key is present return an error
	if len(keys) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Keys data found", "data": nil})
	}

	// Else return keys
	return c.JSON(fiber.Map{"status": "success", "message": "Keys Found", "data": keys})
}

// GetKeysUsingBuildingName func gets all keys based on the given rfid
// @Description Gets all keys based on the given rfid
// @Tags Key
// @Accept json
// @Produce json
// @Success 200 {array} model.Key
// @router /api/key/{rfid} [get]
func GetKeyUsingRFID(c *fiber.Ctx) error {
	db := database.DB
	var key model.Key

	// Read the param key_rfid
	key_rfid := c.Params("rfid")

	// find all keys on the given rfid in the database
	db.Find(&key, "rfid = ?", key_rfid)
	// If key does not exist, return an error
	if key.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Key does not exist.", "data": nil})
	}

	// Else return keys
	return c.JSON(fiber.Map{"status": "success", "message": "Key Found", "data": key})
}

// CreateKey func creates a key
// @Description Creates a Key
// @Tags Key
// @Accept json
// @Produce json
// @Param rfid body string true "rfid"
// @Param status body string true "status"
// @Param building_name body string true "building_name"
// @Param room_name body string true "room_name"
// @Success 200 {object} model.Key
// @router /api/key [post]
func CreateKey(c *fiber.Ctx) error {
	db := database.DB
	key := new(model.Key)

	type KeyToAdd struct {
		RFID         string          `json:"rfid"`
		Status       model.KeyStatus `json:"status"`
		BuildingName string          `json:"building_name"`
		RoomName     string          `json:"room_name"`
	}

	key_to_add := new(KeyToAdd)

	// Parse the body to the key object
	err := c.BodyParser(key_to_add)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Return invalid school_id if empty or null
	if key_to_add.RFID == uuid.Nil.String() || key_to_add.RFID == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	// Create a temporary building data
	var storedBuidling model.Building
	db.Find(&storedBuidling, "name = ?", key_to_add.BuildingName)
	// If building name does not exists, return an error
	if storedBuidling.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Building does not exist.", "data": nil})
	}

	// Create a temporary room data
	var storedRoom model.Room
	db.Find(&storedRoom, "name = ?", key_to_add.RoomName)
	// If room name does not exists, return an error
	if storedRoom.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room does not exist.", "data": nil})
	}
	var keys []model.Key
	// Check if room id already exist in keys table
	db.Find(&keys, "room_id = ?", storedRoom.ID)
	if len(keys) != 0 {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room already has a key.", "data": nil})
	}

	// Create a temporary key data
	var storedKey model.Key

	// Find the key with the given school_id
	db.Find(&storedKey, "rfid = ?", key_to_add.RFID)

	// If key school id exists, return an error
	if storedKey.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Key with the same rfid already exist.", "data": nil})
	}

	// Add a uuid to the new key
	key.ID = uuid.New()

	key.RFID = key_to_add.RFID
	key.Status = key_to_add.Status
	key.BuildingID = storedBuidling.ID
	key.RoomID = storedRoom.ID
	key.BuildingName = storedBuidling.Name
	key.RoomName = storedRoom.Name
	key.RoomFloor = storedRoom.Floor

	// Create the Key and return error if encountered
	err = db.Create(&key).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create key", "data": err})
	}

	// Return the created key
	return c.JSON(fiber.Map{"status": "success", "message": "Key created", "data": key})
}

// GetStudent func get one student by school_id
// @Description Get one student by school_id
// @Tags Student
// @Accept json
// @Produce json
// @Success 200 {object} model.Student
// @router /api/student/{school_id} [get]
// func GetStudent(c *fiber.Ctx) error {
// 	db := database.DB
// 	var student model.Student

// 	// Read the param school_id
// 	school_id := c.Params("school_id")

// 	// Find the student with the given school_id
// 	db.Find(&student, "school_id = ?", school_id)

// 	// If no such student present, return an error
// 	if student.ID == uuid.Nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Student not found", "data": nil})
// 	}

// 	// Return the student with the specified school_id
// 	return c.JSON(fiber.Map{"status": "success", "message": "Student Found", "data": student})
// }

// UpdateKey update a key by rfid
// @Description Update a Key by rfid
// @Tags Key
// @Accept json
// @Produce json
// @Param building_name body string true "building_name"
// @Param room_name body string true "room_name"
// @Param status body string true "status"
// @Success 200 {object} model.Key
// @router /api/key/{rfid} [put]
func UpdateKey(c *fiber.Ctx) error {
	// Create a struct for updating only writable values
	type KeyToUpdate struct {
		BuildingName string          `json:"building_name"`
		RoomName     string          `json:"room_name"`
		Status       model.KeyStatus `json:"status"`
	}

	db := database.DB
	var key model.Key

	// Read the param school_id
	rfid := c.Params("rfid")

	// Find the key with the given school_id param
	db.Find(&key, "rfid = ?", rfid)

	// If no such key, return an error
	if key.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Key not found", "data": nil})
	}

	// Store the body containing the updated data
	var key_to_update KeyToUpdate
	err := c.BodyParser(&key_to_update)
	// Return parsing error if encountered
	if err != nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create a temporary building data
	var storedBuidling model.Building
	db.Find(&storedBuidling, "name = ?", key_to_update.BuildingName)
	// If building does not exists, return an error
	if storedBuidling.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Building does not exist.", "data": nil})
	}

	// Create a temporary room data
	var storedRoom model.Room
	db.Find(&storedRoom, "name = ?", key_to_update.RoomName)
	// If room does not exists, return an error
	if storedRoom.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room does not exist.", "data": nil})
	}

	// Edit the key
	key.BuildingID = storedBuidling.ID
	key.RoomID = storedRoom.ID
	key.Status = key_to_update.Status

	// Save the Changes
	db.Save(&key)

	// Return the updated key
	return c.JSON(fiber.Map{"status": "success", "message": "Key Updated", "data": key})
}

// DeleteKey delete a key by rfid
// @Description Delete a Key by rfid
// @Tags Key
// @Accept json
// @Produce json
// @Success 200
// @router /api/key/{rfid} [delete]
func DeleteKey(c *fiber.Ctx) error {
	db := database.DB
	var key model.Key

	// Read the param key_rfid
	key_rfid := c.Params("rfid")

	// Find the key with the given rfid param
	db.Find(&key, "rfid = ?", key_rfid)

	// If no such key present return an error
	if key.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Key not found", "data": nil})
	}

	// Delete the key
	err := db.Delete(&key, "rfid = ?", key_rfid).Error

	// Return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete key", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Key Deleted"})
}
