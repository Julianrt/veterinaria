package models

import (
	"errors"
	"time"
)

//CitaReservada estructura para manejar el modelo de citas
type CitaReservada struct {
	IDCita    int       `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id_cita"`
	IDDueno   int       `gorm:"" json:"id_dueno"`
	IDMascota int       `gorm:"" json:"id_mascota"`
	Fecha     time.Time `gorm:"" json:"fecha"`
}

//Citas lista de cita
type Citas []CitaReservada

//NewCita crea y retorna un modelo cita
func NewCita(idDueno, idMascota int, fecha time.Time) *CitaReservada {
	cita := &CitaReservada{
		IDDueno:   idDueno,
		IDMascota: idMascota,
		Fecha:     fecha,
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
func GetCitaByID(id int) (*CitaReservada, error) {
	var cita CitaReservada
	err := First(&cita, id)
	return &cita, err
}

//Save f
func (c *CitaReservada) Save() error {
	if c.IDCita == 0 {
		return c.create()
	}
	return c.update()
}

func (c *CitaReservada) create() error {
	return Create(c)
}

func (c *CitaReservada) update() error {
	return Save(c)
}

//Delete f
func (c *CitaReservada) Delete() error {
	if c.IDCita == 0 {
		return errors.New("No se puede eliminar este modelo cita porque tiene el id = 0")
	}
	return Delete(c)
}
