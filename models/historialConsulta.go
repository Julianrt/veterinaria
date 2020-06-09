package models

import (
	"errors"
	"time"
)

//HistorialConsulta estructura para manejar el modelo de historial
type HistorialConsulta struct {
	IDConsulta    int       `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id_consulta"`
	IDDueno       int       `gorm:"" json:"id_dueno"`
	IDMascota     int       `gorm:"" json:"id_mascota"`
	Prescripcion  string    `gorm:"" json:"prescripcion"`
	FechaConsulta time.Time `gorm:"" json:"fecha_consulta"`
}

//Historials lista de historial
type Historials []HistorialConsulta

//NewHistorial f
func NewHistorial(idDueno, idMascota int, prescripcion string, fecha time.Time) *HistorialConsulta {
	h := &HistorialConsulta{
		IDDueno:       idDueno,
		IDMascota:     idMascota,
		Prescripcion:  prescripcion,
		FechaConsulta: fecha,
	}
	return h
}

//GetHistorials f
func GetHistorials() (*Historials, error) {
	var h Historials
	err := Find(&h)
	return &h, err
}

//GetHistorialByID f
func GetHistorialByID(id int) (*HistorialConsulta, error) {
	var h HistorialConsulta
	err := First(&h, id)
	return &h, err
}

//Save f
func (h *HistorialConsulta) Save() error {
	if h.IDConsulta == 0 {
		return h.create()
	}
	return h.update()
}

func (h *HistorialConsulta) create() error {
	return Create(h)
}

func (h *HistorialConsulta) update() error {
	return Save(h)
}

//Delete f
func (h *HistorialConsulta) Delete() error {
	if h.IDConsulta == 0 {
		return errors.New("No se pudo eliminar el modelo ya que su id tiene valor 0")
	}
	return Delete(h)
}
