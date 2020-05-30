package models

import "errors"

//Empleado estructura para manejar el modelo de empleados
type Empleado struct {
	IDEmpleado int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nombre     string `gorm:""`
	Direccion  string `gorm:""`
	Telefono   string `gorm:""`
}

//Empleados lista de Empleado
type Empleados []Empleado

//NewEmpleado f
func NewEmpleado(nombre, direccion, telefono string) *Empleado {
	empleado := &Empleado{
		Nombre:    nombre,
		Direccion: direccion,
		Telefono:  telefono,
	}
	return empleado
}

//GetEmpleados f
func GetEmpleados() (*Empleados, error) {
	var empleados Empleados
	err := Find(&empleados)
	return &empleados, err
}

//GetEmpleadoByID f
func GetEmpleadoByID(id int) (*Empleado, error) {
	var empleado Empleado
	err := First(&empleado, id)
	return &empleado, err
}

//Save f
func (e *Empleado) Save() error {
	if e.IDEmpleado == 0 {
		return e.create()
	}
	return e.update()
}

func (e *Empleado) create() error {
	return Create(e)
}

func (e *Empleado) update() error {
	return Save(e)
}

//Delete f
func (e *Empleado) Delete() error {
	if e.IDEmpleado == 0 {
		return errors.New("No se puede eliminar este modelo porque tiene un id = 0")
	}
	return Delete(e)
}
