package models

import (
	"errors"
	"time"
)

//Venta estructura para manejar el modelo de las ventas
type Venta struct {
	IDVenta    int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id_venta"`
	IDEmpleado int       `gorm:"" json:"id_empleado"`
	Total      float32   `gorm:"" json:"total"`
	Fecha      time.Time `gorm:"" json:"fecha"`
}

//Ventas es una lista de venta
type Ventas []Venta

//NewVenta f
func NewVenta(idEmpleado int, total float32, fecha time.Time) *Venta {
	v := &Venta{
		IDEmpleado: idEmpleado,
		Total:      total,
		Fecha:      fecha,
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
