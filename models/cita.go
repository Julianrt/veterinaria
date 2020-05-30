package models

import (
	"errors"
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

//Citas lista de cita
type Citas []Cita

//NewCita crea y retorna un modelo cita
func NewCita(idServicio int, precio float32, fecha time.Time, descripcion string) *Cita {
	cita := &Cita{
		IDServicio:  idServicio,
		Precio:      precio,
		Fecha:       fecha,
		Descripcion: descripcion,
	}
	return cita
}

//GetCitas f
func GetCitas() (*Citas, error) {
	var citas Citas
	err := Find(&citas)
	return &citas, err
}

//GetCitaByID f
func GetCitaByID(id int) (*Cita, error) {
	var cita Cita
	err := First(&cita, id)
	return &cita, err
}

//Save f
func (c *Cita) Save() error {
	if c.IDCita == 0 {
		return c.create()
	}
	return c.update()
}

func (c *Cita) create() error {
	return Create(c)
}

func (c *Cita) update() error {
	return Save(c)
}

//Delete f
func (c *Cita) Delete() error {
	if c.IDCita == 0 {
		return errors.New("No se puede eliminar este modelo cita porque tiene el id = 0")
	}
	return Delete(c)
}
