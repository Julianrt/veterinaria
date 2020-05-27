package models

import (
	"time"
)

//Venta estructura para manejar el modelo de las ventas
type Venta struct {
	IDVenta    int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	IDEmpleado int       `gorm:""`
	Total      float32   `gorm:""`
	fecha      time.Time `gorm:""`
}
