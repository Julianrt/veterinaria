package models

//Historial estructura para manejar el modelo de historial
type Historial struct {
	IDHistorial  int    `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	IDDueno      int    `json:"" `
	Enfermedades string `json:"" `
	Medicamentos string `json:"" `
	Prescripcion string `json:"" `
}
