package handlers

import (
	"github.com/gofiber/fiber"
)

//Historial handler que renderiza un template
func Historial(c *fiber.Ctx) {
	c.Render("historial", fiber.Map{
		"Title": "Historial",
	})
}

//Agenda handler que renderiza un template
func Agenda(c *fiber.Ctx) {
	c.Render("agenda", fiber.Map{
		"Title": "Agenda",
	})
}

//Consulta handler que renderiza un template
func Consulta(c *fiber.Ctx) {
	c.Render("consulta", fiber.Map{
		"Title": "Consulta",
	})
}

//Registrar handler que renderiza un template
func Registrar(c *fiber.Ctx) {
	c.Render("registrar", fiber.Map{
		"Title": "Registrar",
	})
}
