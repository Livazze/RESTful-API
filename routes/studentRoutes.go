package routes

import (
	"db-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler *handlers.StudentHandler) {
	// Endpoint'leri /api/students altında topluyoruz
	api := app.Group("/api/students")

	api.Get("/", handler.GetAll)       // Tüm öğrencileri listeler
	api.Get("/:id", handler.GetByID)   // id ile tek öğrenci getiriir : sabit bir değer olmadığı anlamına gelir istmeci dinamik bir değer yazar
	api.Post("/", handler.Create)      // Yeni öğrenci ekler
	api.Put("/:id", handler.Update)    // olan öğrenciyi günceller
	api.Delete("/:id", handler.Delete) // id ye göre öğrenciyi siler
}

//yönlendirme tablosu gibi
