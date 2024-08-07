package routes

import (
	"os"
	"qrcode_statistics/internal/handlers"
	"qrcode_statistics/internal/middleware"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/login", handlers.Authen)
	apiGroup.Get("/health", handlers.Health)

	apiGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
	}))

	users := apiGroup.Group("/users", middleware.AuthMiddleware)
	users.Get("", handlers.GetMemberById)
	users.Post("", handlers.CreateMember)
	users.Put("", handlers.UpdateMember)
	users.Delete("", handlers.DeleteMember)

	event := apiGroup.Group("/event", middleware.AuthMiddleware)
	event.Get(":id", handlers.GetEventById)
	event.Post("", handlers.CreateEvent)
	event.Put(":id", handlers.UpdateEvent)
	event.Delete(":id", handlers.DeleteEvent)

	qrcode := apiGroup.Group("/qrcode", middleware.AuthMiddleware)
	qrcode.Get(":id", handlers.GetQrcodeById)
	qrcode.Post("", handlers.CreateQrcode)
	qrcode.Put(":id", handlers.UpdateQrcode)
	qrcode.Delete(":id", handlers.DeleteQrcode)

	statistics := apiGroup.Group("/statistics", middleware.AuthMiddleware)
	statistics.Post("", handlers.Statistics)
}
