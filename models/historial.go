package models

import "errors"

//Historial estructura para manejar el modelo de historial
type Historial struct {
	IDHistorial  int    `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id_historial"`
	IDDueno      int    `gorm:"" json:"id_dueno"`
	Enfermedades string `gorm:"" json:"enfermedades"`
	Medicamentos string `gorm:"" json:"medicamentos"`
	Prescripcion string `gorm:"" json:"prescripcion"`
}

//Historials lista de historial
type Historials []Historial

//NewHistorial f
func NewHistorial(idDueno int, enfermadades, medicamentos, prescripcion string) *Historial {
	h := &Historial{
		IDDueno:      idDueno,
		Enfermedades: enfermadades,
		Medicamentos: medicamentos,
		Prescripcion: prescripcion,
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
func GetHistorialByID(id int) (*Historial, error) {
	var h Historial
	err := First(&h, id)
	return &h, err
}

//Save f
func (h *Historial) Save() error {
	if h.IDHistorial == 0 {
		return h.create()
	}
	return h.update()
}

func (h *Historial) create() error {
	return Create(h)
}

func (h *Historial) update() error {
	return Save(h)
}

//Delete f
func (h *Historial) Delete() error {
	if h.IDHistorial == 0 {
		return errors.New("No se pudo eliminar el modelo ya que su id tiene valor 0")
	}
	return Delete(h)
}
