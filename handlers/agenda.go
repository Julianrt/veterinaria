package handlers

import (
	"net/http"

	"github.com/Julianrt/veterinaria/models"
	"github.com/Julianrt/veterinaria/utils"

	"github.com/gofiber/fiber"
)

//AgendarConsulta handler
func AgendarConsulta(c *fiber.Ctx) {
	fecha := new(models.FechaCita)

	if err := c.BodyParser(fecha); err != nil {
		c.Status(http.StatusBadRequest).Send("Error: " + err.Error())
		return
	}

	date, err := utils.FillDate(fecha.Fecha, fecha.Hora)
	if err != nil {
		c.Status(http.StatusBadRequest).Send("Error: " + err.Error())
		return
	}

	if len(fecha.ClienteCorreo) < 1 {
		c.Status(http.StatusBadRequest).Send("Error: necesario escribir un correo")
		return
	}
	cliente, err := models.GetClienteByCorreo(fecha.ClienteCorreo)
	if err != nil {
		if err.Error() == "record not found" {
			cliente.NombreDueno = fecha.ClienteNombre
			cliente.Correo = fecha.ClienteCorreo
			if err := cliente.Save(); err != nil {
				c.Status(http.StatusInternalServerError).Send("Error: " + err.Error())
				return
			}
		} else {
			c.Status(http.StatusInternalServerError).Send("Error: " + err.Error())
			return
		}
	}
	if cliente.NombreDueno != fecha.ClienteNombre {
		c.Status(http.StatusBadRequest).Send("Error: Este correo está registrado con otro cliente")
		return
	}

	//mascota, _ := models.ValidateMascotaOwner(cliente.IDDueno, "api")

	var cita models.CitaReservada
	cita.IDDueno = cliente.IDDueno
	cita.IDMascota = 1
	cita.Fecha = date
	err = cita.Save()
	if err != nil {
		c.Status(http.StatusInternalServerError).Send("Error: " + err.Error())
		return
	}
	c.Status(http.StatusCreated)
}
