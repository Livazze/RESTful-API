package routes

import (
	"db-api/handlers"

	"github.com/gofiber/fiber/v2"
)

// Endpoint'leri /api/students altında topluyoruz
func SetupRoutes(app *fiber.App, handler *handlers.StudentHandler) {
	api := app.Group("/api/students")

	api.Get("", handler.GetAll)        // /api/students
	api.Get("/:id", handler.GetByID)   // /api/students/:id
	api.Post("", handler.Create)       // /api/students
	api.Put("/:id", handler.Update)    // /api/students/:id
	api.Delete("/:id", handler.Delete) // /api/students/:id
}

//yönlendirme tablosu gibi
