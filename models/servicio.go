package models

//Servicio struct para manejar el modelo de servicios
type Servicio struct {
	IDServicio     int     `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id_servicio"`
	NombreServicio string  `gorm:"" json:"nombre_servicio"`
	PrecioServicio float32 `gorm:"" json:"precio_servicio"`
}
