package handlers

import (
	"log"

	"github.com/Julianrt/veterinaria/models"
	"github.com/Julianrt/veterinaria/utils"

	"github.com/gofiber/fiber"
)

//Historial handler que renderiza un template
func Historial(c *fiber.Ctx) {
	c.Render("historial", fiber.Map{
		"Title": "Historial",
	})
}

//AgendarCita handler que renderiza un template
func AgendarCita(c *fiber.Ctx) {

	if c.Method() == "GET" {
		c.Render("agendar_cita", fiber.Map{
			"Title": "Agendar cita",
		})
	} else if c.Method() == "POST" {
		fecha := c.FormValue("fecha")
		hora := c.FormValue("hora")
		nombreDueno := c.FormValue("nombre_dueno")
		nombreMascota := c.FormValue("nombre_mascota")
		telefono := c.FormValue("telefono")
		correo := c.FormValue("correo")

		if fecha == "" || hora == "" || nombreDueno == "" || nombreMascota == "" ||
			telefono == "" || correo == "" {

			log.Println("Tienes que llenar todos los campos")
			c.Redirect("/agendar/")
			return
		}

		cliente, err := models.GetClienteByTelefono(telefono)
		if cliente.IDDueno == 0 || err != nil {
			log.Println("ERROR: " + err.Error())
			cliente = models.NewCliente(nombreDueno, telefono, correo)
			if err = cliente.Save(); err != nil {
				log.Println("No se pudo guardar el cliente -> " + err.Error())
			}
		}

		mascota, err := models.ValidateMascotaOwner(cliente.IDDueno, nombreMascota)
		if err != nil || mascota.IDMascota == 0 {
			log.Println("No se encontro es mascota")
			mascota.IDDueno = cliente.IDDueno
			mascota.NombreMascota = nombreMascota
			err = mascota.Save()
			if err != nil {
				log.Println("No se pudo guardar la mascota -> " + err.Error())
			}
		}

		date, err := utils.FillDate(fecha, hora)
		if err != nil {
			log.Println(err.Error())
		}
		cita := models.NewCita(cliente.IDDueno, mascota.IDMascota, date)
		if err := cita.Save(); err != nil {
			log.Println("No se pude guardar la cita -> " + err.Error())
		}

		c.Redirect("/agendar/")
	}

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
