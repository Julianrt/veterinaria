package webserver

import (
	"os"

	"github.com/Julianrt/veterinaria/models"
	"github.com/Julianrt/veterinaria/routes"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

//StartServer put server running
func StartServer() {
	models.InitDB()
	defer models.CloseConnection()

	app := fiber.New()

	app.Use(cors.New())

	routes.StartHandleRoutes(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
