package models

import (
	"time"
)

//Agenda estructura para manejar el modelo de agenda
type Agenda struct {
	IDAgenda   int       `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	IDDueno    int       `gorm:""`
	IDServicio int       `gorm:""`
	Fecha      time.Time `gorm:""`
}
