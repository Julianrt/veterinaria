package webserver

import (
	"github.com/Julianrt/veterinaria/models"
	"github.com/Julianrt/veterinaria/routes"

	"github.com/gofiber/fiber"
)

//StartServer put server running
func StartServer() {
	models.InitDB()
	defer models.CloseConnection()

	app := fiber.New()

	routes.StartHandleRoutes(app)

	app.Listen(3000)
}
