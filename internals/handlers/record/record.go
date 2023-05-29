package recordHandler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/ada-api/database"
	"github.com/vincemoke66/ada-api/internals/model"
	"gorm.io/gorm"
)

// GetAllRecords func gets all records
// @Description Gets all records
// @Tags Record
// @Accept json
// @Produce json
// @Success 200 {array} model.Record
// @router /api/record [get]
func GetAllRecords(c *fiber.Ctx) error {
	db := database.DB
	var records []model.Record

	// find all records
	db.Order("created_at DESC").Find(&records)

	// If no record is present return an error
	if len(records) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Records data found", "data": nil})
	}

	// Else return records
	return c.JSON(fiber.Map{"status": "success", "message": "Records Found", "data": records})
}

// CreateRecord func creates a record
// @Description Creates a Record
// @Tags Record
// @Accept json
// @Produce json
// @Param type body string true "type"
// @Param school_id body string true "school_id"
// @Param key_rfid body string true "key_rfid"
// @Success 200 {object} model.Record
// @router /api/record [post]
func CreateRecord(c *fiber.Ctx) error {
	db := database.DB
	record := new(model.Record)

	type RecordToAdd struct {
		StudentRFID string
		RoomName    string
	}

	record_to_add := new(RecordToAdd)
	err := c.BodyParser(record_to_add)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Read if the sent studentRFID is valid
	// Create a temporary student data
	var storedStudent model.Student
	db.Find(&storedStudent, "rfid = ?", record_to_add.StudentRFID)
	// If student does not exist, return an error
	if storedStudent.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Student does not exist.", "data": nil})
	}

	// Read if the sent roomName is valid
	// Create a temporary room data
	var storedRoom model.Room
	db.Find(&storedRoom, "name = ?", record_to_add.RoomName)
	// If room does not exist, return an error
	if storedRoom.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room does not exist.", "data": nil})
	}

	// Check if the entered room and current time has a schedule
	currentTime := time.Now()
	hasSchedule, scheduleFound, err := CheckSchedule(currentTime, record_to_add.RoomName)

	if !hasSchedule {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid input.", "data": nil})
	}

	// If the currentTime and roomName has a correct schedule
	// Create a new record
	record.ID = uuid.New()
	record.StudentID = storedStudent.ID
	record.Section = storedStudent.Section
	record.Subject = scheduleFound.Subject
	record.RoomName = storedRoom.Name

	// Create the Record and return error if encountered
	err = db.Create(&record).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create record", "data": err})
	}

	// Return the created record
	return c.JSON(fiber.Map{"status": "success", "message": "Record created", "data": record})
}

// Function to check if the input matches a schedule
func CheckSchedule(inputTime time.Time, roomName string) (bool, model.Schedule, error) {
	db := database.DB
	var schedule model.Schedule

	// Query to check if there is a matching schedule
	err := db.Where("start_time <= ? AND end_time >= ? AND room_name = ?", inputTime, inputTime, roomName).First(&schedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// No matching schedule found
			return false, model.Schedule{}, nil
		}
		// Error occurred during the query
		return false, model.Schedule{}, err
	}

	// Matching schedule found
	return true, schedule, nil
}
