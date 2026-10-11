package main

import (
	"context"
	"db-api/database"
	"db-api/handlers"
	"db-api/repository"
	"db-api/routes"
	"log"
	"os"

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
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
		AllowCredentials: false,
	}))

	// Routesları ekler
	routes.SetupRoutes(app, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("The server is running on port :%s", port)
	log.Fatal(app.Listen(":" + port))
}

/*hiçbiri kendi içind esistemi yaratmıyor birbirinden paramete olarak alıyor
Dependency Injection bir bileşenin kendi bağımlılıklarını içerde oluşturmak yerine
 bunları dış kaynaklardan almasını sağlayan bir yazılım tasarım desenidir. new ile bağımlılıklar dışarıdan verirldi*/
