package handlers

import (
	"log"

	"github.com/gofiber/fiber"
)

//GetServicios f
func GetServicios(c *fiber.Ctx) {
	data := struct {
		Clave     int    `json:"clave"`
		Nombre    string `json:"nombre"`
		Precio    int    `json:"precio"`
		URLImagen string `json:"url_imagen"`
	}{
		Clave:     1,
		Nombre:    "Consulta",
		Precio:    50,
		URLImagen: "https://consulta-veterinaria.herokuapp.com/public/images/consulta.jpg",
	}

	if err := c.JSON(data); err != nil {
		log.Println(err.Error())
	}
}
