package models

import (
	"time"
)

//Cita estructura para manejar el modelo de citas
type Cita struct {
	IDCita      int       `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	IDServicio  int       `gorm:""`
	Precio      float32   `gorm:""`
	Fecha       time.Time `gorm:""`
	Descripcion string    `gorm:""`
}
