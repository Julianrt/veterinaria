package models

//DetalleVenta estructura para manejar el modelo de los detalles de venta
type DetalleVenta struct {
	IDServicio int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	IDDetalle  int `gorm:""`
	IDVenta    int `gorm:""`
}
