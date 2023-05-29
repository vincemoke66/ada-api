package router

import (
	"github.com/gofiber/fiber/v2"
	buildingRoutes "github.com/vincemoke66/ada-api/internals/routes/building"
	instructorRoutes "github.com/vincemoke66/ada-api/internals/routes/instructor"
	keyRoutes "github.com/vincemoke66/ada-api/internals/routes/key"
	recordRoutes "github.com/vincemoke66/ada-api/internals/routes/record"
	roomRoutes "github.com/vincemoke66/ada-api/internals/routes/room"
	studentRoutes "github.com/vincemoke66/ada-api/internals/routes/student"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	studentRoutes.SetupStudentRoutes(api)
	instructorRoutes.SetupStudentRoutes(api)
	buildingRoutes.SetupStudentRoutes(api)
	roomRoutes.SetupStudentRoutes(api)
	keyRoutes.SetupStudentRoutes(api)
	recordRoutes.SetupStudentRoutes(api)
}
