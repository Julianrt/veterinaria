package routes

import (
	"github.com/Julianrt/veterinaria/handlers"

	"github.com/gofiber/fiber"
)

//StartHandleRoutes f
func StartHandleRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Jalando")
	})
	empleadoRoutes(app)
}

func empleadoRoutes(app *fiber.App) {
	app.Post("/empleados/", handlers.CreateEmpleado)
	app.Get("/empleados/", handlers.GetEmpleados)
	app.Get("/empleados/:id/", handlers.GetEmpleado)
	app.Put("/empleados/:id/", handlers.UpdateEmpleado)
	app.Delete("/empleados/:id/", handlers.DeleteEmpleado)
}
