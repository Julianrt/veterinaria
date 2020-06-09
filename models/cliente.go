package models

import (
	"errors"
)

//Cliente estructura para manejar el modelo cliente
type Cliente struct {
	IDDueno     int    `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id_dueno"`
	IDMascota   int    `gorm:"" json:"id_mascota"`
	NombreDueno string `gorm:"" json:"nombre_dueno"`
	Telefono    string `gorm:"" json:"telefono"`
	Correo      string `gorm:"" json:"correo"`
}

//Clientes lista de cliente
type Clientes []Cliente

//NewCliente crea un objeto cliente y lo retorna
func NewCliente(idMastoca int, nombreDueno, telefono, correo string) *Cliente {

	cliente := &Cliente{
		IDMascota:   idMastoca,
		NombreDueno: nombreDueno,
		Telefono:    telefono,
		Correo:      correo,
	}
	return cliente
}

//GetClientes f
func GetClientes() (*Clientes, error) {
	var clientes Clientes
	err := Find(&clientes)
	return &clientes, err
}

//GetClienteByID f
func GetClienteByID(id int) (*Cliente, error) {
	var cliente Cliente
	err := First(&cliente, id)
	return &cliente, err
}

//Save si el IDDueno es 0, se registra el modelo en la bd,
// y si es diferente a 0, se actualiza en bd
func (c *Cliente) Save() error {
	if c.IDDueno == 0 {
		return c.create()
	}
	return c.update()
}

func (c *Cliente) create() error {
	return Create(c)
}

func (c *Cliente) update() error {
	return Save(c)
}

//Delete f
func (c *Cliente) Delete() error {
	if c.IDDueno == 0 {
		return errors.New("No se puede eliminar este cliente porque tiene el id = 0")
	}
	return Delete(c)
}
