package models

//Cliente estructura para manejar el modelo cliente
type Cliente struct {
	IDDueno        int     `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	IDHistorial    int     `gorm:""`
	NombreDueno    string  `gorm:""`
	NombrePaciente string  `gorm:""`
	Telefono       int     `gorm:""`
	TipoAnimal     string  `gorm:""`
	Peso           float32 `gorm:""`
	Edad           int     `gorm:""`
	Vacunas        string  `gorm:""`
}
