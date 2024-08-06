package routes

import (
	"qrcode_statistics/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/health", handlers.Health)

	users := apiGroup.Group("/users")
	users.Get(":id", handlers.GetMemberById)
	users.Post("", handlers.CreateMember)
	users.Put(":id", handlers.UpdateMember)
	users.Delete(":id", handlers.DeleteMember)

	event := apiGroup.Group("/event")
	event.Get(":id", handlers.GetEventById)
	event.Post("", handlers.CreateEvent)
	event.Put(":id", handlers.UpdateEvent)
	event.Delete(":id", handlers.DeleteEvent)
}
