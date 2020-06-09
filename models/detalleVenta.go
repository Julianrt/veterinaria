package models

import (
	"errors"
)

//DetalleVenta estructura para manejar el modelo de los detalles de venta
type DetalleVenta struct {
	IDDetalleVenta int `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id_detalle_venta"`
	IDServicio     int `gorm:"" json:"id_servicio"`
	IDVenta        int `gorm:"" json:"id_venta"`
	IDEmpleado     int `gorm:"" json:"id_empleado"`
}

//DetalleVentas lista de detalle venta
type DetalleVentas []DetalleVenta

//NewDetalleVenta f
func NewDetalleVenta(idServicio, idVenta, idEmpleado int) *DetalleVenta {
	detalleVenta := &DetalleVenta{
		IDServicio: idServicio,
		IDVenta:    idVenta,
		IDEmpleado: idEmpleado,
	}
	return detalleVenta
}

//GetDetalleVentas f
func GetDetalleVentas() (*DetalleVentas, error) {
	var detalleVentas DetalleVentas
	err := Find(&detalleVentas)
	return &detalleVentas, err
}

//GetDetalleVentaByID f
func GetDetalleVentaByID(id int) (*DetalleVenta, error) {
	var detalleVenta DetalleVenta
	err := First(&detalleVenta, id)
	return &detalleVenta, err
}

//Save f
func (dv *DetalleVenta) Save() error {
	if dv.IDDetalleVenta == 0 {
		return dv.create()
	}
	return dv.update()
}

func (dv *DetalleVenta) create() error {
	return Create(dv)
}

func (dv *DetalleVenta) update() error {
	return Save(dv)
}

//Delete f
func (dv *DetalleVenta) Delete() error {
	if dv.IDDetalleVenta == 0 {
		return errors.New("No se puedo eliminar el modelo de DetalleVenta porque tiene su id = 0")
	}
	return Delete(dv)
}
