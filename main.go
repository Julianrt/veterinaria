package main

import (
	"log"

	"github.com/Julianrt/veterinaria/models"
)

func main() {
	models.InitDB()

	var err error

	err = models.Create(&models.Empleado{Nombre: "nombre1", Direccion: "Direccion1", Telefono: "1234567890"})
	err = models.Create(&models.Empleado{Nombre: "nombre2", Direccion: "Direccion2", Telefono: "1234567890"})
	err = models.Create(&models.Empleado{Nombre: "nombre3", Direccion: "Direccion3", Telefono: "1234567890"})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Hola1")
	}

	var empleados []models.Empleado
	if err = models.Find(&empleados); err != nil {
		log.Println(err)
	} else {
		log.Println("Hola2")
	}
	log.Println(empleados)

	log.Println("-------------------------------------")

	var empleado models.Empleado
	models.First(&empleado, 2)
	log.Println(empleado)

	empleado.Nombre = "Panfilo"
	empleado.IDEmpleado = 0
	models.Save(&empleado)

	models.First(&empleado, 2)
	log.Println(empleado)

	models.Find(&empleados)
	log.Println(empleados)

	log.Println("--------------------------------------")

	models.Delete(empleado)
	models.Find(&empleados)
	log.Println(empleados)

}
