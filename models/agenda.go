package models

import (
	"errors"
	"time"
)

//Agenda estructura para manejar el modelo de agenda
type Agenda struct {
	IDAgenda   int       `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	IDDueno    int       `gorm:""`
	IDServicio int       `gorm:""`
	Fecha      time.Time `gorm:""`
}

//Agendas lista de agenda
type Agendas []Agenda

//NewAgenda Crea y retorna una agenda
func NewAgenda(idDueno, idServicio int, fecha time.Time) *Agenda {
	agenda := &Agenda{
		IDDueno:    idDueno,
		IDServicio: idServicio,
		Fecha:      fecha,
	}
	return agenda
}

//GetAgendas f
func GetAgendas() (*Agendas, error) {
	var agendas Agendas
	err := Find(&agendas)
	return &agendas, err
}

//GetAgendaByID f
func GetAgendaByID(id int) (*Agenda, error) {
	var agenda Agenda
	err := First(&agenda, id)
	return &agenda, err
}

//Save f
func (a *Agenda) Save() error {
	if a.IDAgenda == 0 {
		return a.create()
	}
	return a.update()
}

func (a *Agenda) create() error {
	return Create(a)
}

func (a *Agenda) update() error {
	return Save(a)
}

//Delete f
func (a *Agenda) Delete() error {
	if a.IDAgenda == 0 {
		return errors.New("No se puede eliminar este modelo agenda porque tiene el id = 0")
	}
	return Delete(a)
}
