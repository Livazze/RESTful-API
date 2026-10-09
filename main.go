package main

import (
	"context"
	"log"

	"db-api/database"
	"db-api/handlers"
	"db-api/repository"
	"db-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// veritabnı bağlantıısnı başlatır
	db := database.ConnectDB()
	defer db.Close()

	// Katmanları birbirine bağlanır (Dependency Injection
	repo := repository.NewStudentRepository(context.Background(), db) //repository db ile etkileşime girer
	handler := handlers.NewStudentHandler(repo)                       //handler repoyu içine alır

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:5174, http://127.0.0.1:5173, http://127.0.0.1:5174",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Routesları ekler
	routes.SetupRoutes(app, handler)

	log.Println("The server is running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}

/*hiçbiri kendi içind esistemi yaratmıyor birbirinden paramete olarak alıyor
Dependency Injection bir bileşenin kendi bağımlılıklarını içerde oluşturmak yerine
 bunları dış kaynaklardan almasını sağlayan bir yazılım tasarım desenidir. new ile bağımlılıklar dışarıdan verirldi*/
