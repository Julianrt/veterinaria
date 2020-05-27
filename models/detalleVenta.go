package models

//DetalleVenta estructura para manejar el modelo de los detalles de venta
type DetalleVenta struct {
	IDDetalleVenta int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	IDServicio     int `gorm:""`
	IDVenta        int `gorm:""`
}
