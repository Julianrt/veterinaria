package models

import (
	"errors"
)

//Cliente estructura para manejar el modelo cliente
type Cliente struct {
	IDDueno        int     `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	IDHistorial    int     `gorm:""`
	NombreDueno    string  `gorm:""`
	NombrePaciente string  `gorm:""`
	Telefono       string  `gorm:""`
	TipoAnimal     string  `gorm:""`
	Peso           float32 `gorm:""`
	Edad           int     `gorm:""`
	Vacunas        string  `gorm:""`
}

//Clientes lista de cliente
type Clientes []Cliente

//NewCliente crea un objeto cliente y lo retorna
func NewCliente(idHistorial int, nombreDueno, nombrePaciente, telefono, tipoAnimal string,
	peso float32, edad int, vacunas string) *Cliente {

	cliente := &Cliente{
		IDHistorial:    idHistorial,
		NombreDueno:    nombreDueno,
		NombrePaciente: nombrePaciente,
		Telefono:       telefono,
		TipoAnimal:     tipoAnimal,
		Peso:           peso,
		Edad:           edad,
		Vacunas:        vacunas,
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
