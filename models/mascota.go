package models

//Mascota struct para manejar el modelo de mascota
type Mascota struct {
	IDMascota     int     `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id_mascota"`
	IDDueno       int     `gorm:"" json:"id_dueno"`
	NombreMascota string  `gorm:"" json:"nombre_mascota"`
	TipoAnimal    string  `gorm:"" json:"tipo_animal"`
	Edad          int     `gorm:"" json:"edad"`
	Peso          float32 `gorm:"" json:"peso"`
	Vacunas       string  `gorm:"" json:"vacunas"`
}
