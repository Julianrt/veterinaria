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
	cliente, _ := models.GetClienteByTelefono("api")
	mascota, _ := models.ValidateMascotaOwner(cliente.IDDueno, "api")

	var cita models.CitaReservada
	cita.IDDueno = cliente.IDDueno
	cita.IDMascota = mascota.IDMascota
	cita.Fecha = date
	err = cita.Save()
	if err != nil {
		c.Status(http.StatusInternalServerError).Send("Error: " + err.Error())
		return
	}
	c.JSON(cita)
}
