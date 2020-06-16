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
		URLImagen: "https://lh3.googleusercontent.com/proxy/kBASuvkCQaHpuULIr7PRnaAAoYTqny4xeNyCQMdC14D7JZu73fDwmJBJrXKrgvctjgh7lH771xPm6PNtEaurKNaIQgD0FX4lOkOX2a_aFlBjUAZFAPvz5ygL3IltO3APeuWs",
	}

	if err := c.JSON(data); err != nil {
		log.Println(err.Error())
	}
}
