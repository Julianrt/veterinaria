package main

import (
	"log"

	"github.com/Julianrt/veterinaria/models"
)

func main() {
	models.InitDB()

	cliente1 := models.NewCliente(111, "Panfilooo", "Panfilito", "6622240000", "perro", 2.4, 2, "parvovirus")
	cliente2 := models.NewCliente(111, "Panfilo Dominguez", "Panfilito", "6622240000", "perro", 2.4, 2, "parvovirus")
	cliente3 := models.NewCliente(111, "Panfilo", "Panfilito", "6622240000", "perro", 2.4, 2, "parvovirus")
	cliente4 := models.NewCliente(111, "Panfiloooo123", "Panfilito", "6622240000", "perro", 2.4, 2, "parvovirus")

	log.Println(cliente2)
	cliente1.Save()
	cliente2.Save()
	cliente3.Save()
	cliente4.Save()
	log.Println(cliente2)

	log.Println("--------------------------------------")
	clientes, err := models.GetClientes()
	if err != nil {
		log.Println("Error al obtener todos los clientes -> " + err.Error())
		return
	}
	log.Println(clientes)
	log.Println("--------------------------------------")

	cliente3.Edad = 10
	cliente3.Save()
	log.Println(cliente3)

	log.Println("--------------------------------------")

	cliente1.Delete()

	clientes, _ = models.GetClientes()
	log.Println(clientes)
}
