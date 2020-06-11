package routes

import (
	"github.com/Julianrt/veterinaria/handlers"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

//StartHandleRoutes f
func StartHandleRoutes(app *fiber.App) {
	app.Settings.Templates = html.New("./templates", ".html")

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Jalando")
	})
	empleadoRoutes(app)
	webApp(app)
}

func empleadoRoutes(app *fiber.App) {
	app.Post("/empleados/", handlers.CreateEmpleado)
	app.Get("/empleados/", handlers.GetEmpleados)
	app.Get("/empleados/:id/", handlers.GetEmpleado)
	app.Put("/empleados/:id/", handlers.UpdateEmpleado)
	app.Delete("/empleados/:id/", handlers.DeleteEmpleado)
}

func webApp(app *fiber.App) {
	app.Get("/historial/", handlers.Historial)
	app.Get("/agendar/", handlers.AgendarCita)
	app.Post("/agendar/", handlers.AgendarCita)
	app.Get("/agenda/", handlers.Agenda)
	app.Get("/consulta/", handlers.Consulta)
	app.Get("/registrar/", handlers.Registrar)
}
