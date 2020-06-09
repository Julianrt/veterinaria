package models

import (
	"errors"
	"time"
)

//Venta estructura para manejar el modelo de las ventas
type Venta struct {
	IDVenta      int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id_venta"`
	IDServicio   int       `gorm:"" json:"id_servicio"`
	IDEmpleado   int       `gorm:"" json:"id_empleado"`
	FechaVenta   time.Time `gorm:"" json:"fecha_venta"`
	PagoServicio float32   `gorm:"" json:"pago_servicio"`
}

//Ventas es una lista de venta
type Ventas []Venta

//NewVenta f
func NewVenta(idServicio, idEmpleado int, fechaVenta time.Time, pagoServicio float32) *Venta {
	v := &Venta{
		IDServicio:   idServicio,
		IDEmpleado:   idEmpleado,
		FechaVenta:   fechaVenta,
		PagoServicio: pagoServicio,
	}
	return v
}

//GetVentas f
func GetVentas() (*Ventas, error) {
	var v Ventas
	err := Find(&v)
	return &v, err
}

//GetVentaByID f
func GetVentaByID(id int) (*Venta, error) {
	var v Venta
	err := First(&v, id)
	return &v, err
}

//Save f
func (v *Venta) Save() error {
	if v.IDVenta == 0 {
		return v.create()
	}
	return v.update()
}

func (v *Venta) create() error {
	return Create(v)
}

func (v *Venta) update() error {
	return Save(v)
}

//Delete f
func (v *Venta) Delete() error {
	if v.IDVenta == 0 {
		return errors.New("No se pudo eliminar venta porque el valor de su id es 0")
	}
	return Delete(v)
}
