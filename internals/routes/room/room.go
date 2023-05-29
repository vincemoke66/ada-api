package roomRoutes

import (
	"github.com/gofiber/fiber/v2"
	roomHandler "github.com/vincemoke66/ada-api/internals/handlers/room"
)

func SetupStudentRoutes(router fiber.Router) {
	room := router.Group("/room")

	// Create a room
	room.Post("/", roomHandler.CreateRoom)
	// Read all rooms
	room.Get("/", roomHandler.GetRooms)
	// Read all rooms on a building
	room.Get("/:building_name", roomHandler.GetRoomsOnBuilding)
	// Read a room
	room.Get("/:name", roomHandler.GetRoom)
	// Update room
	room.Put("/:name", roomHandler.UpdateRoom)
	// Delete a room
	room.Delete("/:name", roomHandler.DeleteRoom)
}
