package handlers

import (
	"net/http"
	"strconv"

	"github.com/Julianrt/veterinaria/models"

	"github.com/gofiber/fiber"
)

//CreateEmpleado handler para guardar un empleado en la bd
func CreateEmpleado(c *fiber.Ctx) {
	e := new(models.Empleado)

	if err := c.BodyParser(e); err != nil {
		c.Status(http.StatusBadRequest).Send(err.Error())
		return
	}
	if len(e.NombreEmpleado) < 1 {
		c.Status(http.StatusBadRequest).Send("Necesario ingresar el nombre")
		return
	}
	if err := e.Save(); err != nil {
		c.Status(http.StatusInternalServerError).Send(err.Error())
		return
	}
	c.JSON(e)
}

//GetEmpleados handler para mostrar todos los empleados
func GetEmpleados(c *fiber.Ctx) {
	empleados, err := models.GetEmpleados()
	if err != nil {
		c.Status(http.StatusInternalServerError).Send(err.Error())
		return
	}
	if err := c.JSON(empleados); err != nil {
		c.Status(http.StatusInternalServerError).Send(err.Error())
	}
}

//GetEmpleado handler para mostrar un empleado en especifico
func GetEmpleado(c *fiber.Ctx) {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.Status(http.StatusBadRequest).Send("el parametro id debe de ser numerico entero mayor a 0")
		return
	}

	empleado, err := models.GetEmpleadoByID(id)
	if err != nil || empleado == nil {
		c.Status(http.StatusBadRequest).Send("no se obtuvo empleado: " + err.Error())
		return
	}

	if err = c.JSON(empleado); err != nil {
		c.Status(http.StatusBadRequest).Send(err.Error())
	}
}

//UpdateEmpleado handler para actualizar un empleado
func UpdateEmpleado(c *fiber.Ctx) {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.Status(http.StatusBadRequest).Send("el parametro id debe de ser numerico entero mayor a 0")
		return
	}

	empleado := new(models.Empleado)
	if err := c.BodyParser(&empleado); err != nil {
		c.Status(http.StatusBadRequest).Send(err.Error())
		return
	}
	if len(empleado.NombreEmpleado) < 1 {
		c.Status(http.StatusBadRequest).Send("Necesario ingresar el nombre")
		return
	}
	empleado.IDEmpleado = id
	if err := empleado.Save(); err != nil {
		c.Status(http.StatusInternalServerError).Send(err.Error())
	}
}

//DeleteEmpleado handler para eliminar un empleado
func DeleteEmpleado(c *fiber.Ctx) {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.Status(http.StatusBadRequest).Send("el parametro id debe de ser numerico entero mayor a 0")
		return
	}

	empleado, err := models.GetEmpleadoByID(id)
	if err != nil || empleado == nil {
		c.Status(http.StatusBadRequest).Send("no se obtuvo empleado: " + err.Error())
		return
	}

	if err = empleado.Delete(); err != nil {
		c.Status(http.StatusInternalServerError).Send(err.Error())
	}
}
