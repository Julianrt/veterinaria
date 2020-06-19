package handlers

import (
	"log"
	"strconv"

	"github.com/Julianrt/veterinaria/models"
	"github.com/Julianrt/veterinaria/utils"

	"github.com/gofiber/fiber"
)

//Historial handler que renderiza un template
func Historial(c *fiber.Ctx) {
	consultas, err := models.GetConsultasShowOrderDesc()
	if err != nil {
		log.Println(err.Error())
	}

	data := struct {
		Consultas models.ConsultasShow
	}{
		consultas,
	}

	c.Render("historial", data)
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

		date, err := utils.FillDate(fecha, hora)
		if err != nil {
			log.Println(err.Error())
		}
		if !utils.ValidateDate(date) {
			log.Println("Fecha incorrecta para guardar cita")
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
		} else {
			if cliente.NombreDueno != nombreDueno || cliente.Correo != correo {
				log.Println("Ese numero telefonico está registrado pero con otro cliente o correo")
				c.Redirect("/agendar/")
				return
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

		cita := models.NewCita(cliente.IDDueno, mascota.IDMascota, date)
		if err := cita.Save(); err != nil {
			log.Println("No se pude guardar la cita -> " + err.Error())
		}

		c.Redirect("/agendar/")
	}

}

//Agenda handler que renderiza un template
func Agenda(c *fiber.Ctx) {
	citas, err := models.GetCitasShowOrderByFecha()
	if err != nil {
		log.Println(err.Error())
	}

	data := struct {
		Citas models.CitasShow
	}{
		citas,
	}

	if err := c.Render("agenda", data); err != nil {
		log.Println(err.Error())
	}
}

//Consulta handler que renderiza un template
func Consulta(c *fiber.Ctx) {

	if c.Method() == "GET" {
		idCita, err := strconv.Atoi(c.Query("cita"))
		if err != nil {
			log.Println(err.Error())
		}

		cita, err := models.GetCitaByID(idCita)
		if err != nil || cita.IDCita == 0 {
			log.Println("No se encontré la cita: " + err.Error())
		}

		cliente, err := models.GetClienteByCita(cita.IDCita)
		if err != nil || cliente.IDDueno == 0 {
			log.Println("No se encontró el cliente " + err.Error())
		}

		mascota, err := models.GetMascotaByCita(cita.IDCita)
		if err != nil || mascota.IDMascota == 0 {
			log.Println("No se encontró la mascota " + err.Error())
		}

		data := struct {
			Cita    models.CitaReservada
			Cliente models.Cliente
			Mascota models.Mascota
		}{
			*cita,
			*cliente,
			*mascota,
		}
		c.Render("consulta", data)

	} else if c.Method() == "POST" {

		idCliente, err := strconv.Atoi(c.Query("cliente"))
		if err != nil {
			log.Println(err.Error())
		}
		cliente, err := models.GetClienteByID(idCliente)
		if err != nil || cliente.IDDueno == 0 {
			log.Println(err.Error())
		}

		idMascota, err := strconv.Atoi(c.Query("mascota"))
		if err != nil {
			log.Println(err.Error())
		}
		mascota, err := models.GetMascotaByID(idMascota)
		if err != nil || mascota.IDMascota == 0 {
			log.Println(err.Error())
		}
		if mascota.IDDueno != cliente.IDDueno {
			log.Println("el cliente no es el dueño de la mascota")
		}

		edadMascota, _ := strconv.Atoi(c.FormValue("edad"))
		pesoMascota, _ := strconv.ParseFloat(c.FormValue("peso"), 32)
		vacunas := c.FormValue("vacunas")
		tipoAnimal := c.FormValue("tipo_animal")
		prescripcion := c.FormValue("prescripcion")

		mascota.Edad = edadMascota
		mascota.Peso = float32(pesoMascota)
		mascota.Vacunas = vacunas
		mascota.TipoAnimal = tipoAnimal

		consulta := models.NewHistorial(cliente.IDDueno, mascota.IDMascota, prescripcion, utils.GetCurrentDate())
		if err := consulta.Save(); err != nil {
			log.Println("No se pudo guardar la consulta -> " + err.Error())
		}

		if err := mascota.Save(); err != nil {
			log.Println("No se pudo actualizar a la mascota -> " + err.Error())
		}
		c.Redirect("/agenda/")
	}
}

//Registrar handler que renderiza un template
func Registrar(c *fiber.Ctx) {
	c.Render("registrar", fiber.Map{
		"Title": "Registrar",
	})
}

//FechasOcupadas handler
func FechasOcupadas(c *fiber.Ctx) {
	fechas, err := utils.GetFechasOcupadas()
	if err != nil {
		log.Println(err)
	}

	c.JSON(fechas)
}
