package models

import "errors"

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

//Mascotas lista de mascota
type Mascotas []Mascota

//NewMascota crea un objeto mascota y retorna un apuntador de ese objeto
func NewMascota(idDueno int, nombreMascota, tipoAnimal string, edad int, peso float32, vacunas string) *Mascota {
	mascota := &Mascota{
		IDDueno:       idDueno,
		NombreMascota: nombreMascota,
		TipoAnimal:    tipoAnimal,
		Edad:          edad,
		Peso:          peso,
		Vacunas:       vacunas,
	}
	return mascota
}

//GetMascotas f
func GetMascotas() (*Mascotas, error) {
	var mascotas Mascotas
	err := Find(&mascotas)
	return &mascotas, err
}

//GetMascotasByIDDueno f
func GetMascotasByIDDueno(idDueno int) (*Mascotas, error) {
	var mascotas Mascotas
	err := FindWithCondition(&mascotas, "id_dueno = ?", idDueno)
	return &mascotas, err
}

//GetMascotaByID f
func GetMascotaByID(id int) (*Mascota, error) {
	var mascota Mascota
	err := First(&mascota, id)
	return &mascota, err
}

//GetMascotaByCita f
func GetMascotaByCita(idCita int) (*Mascota, error) {
	var cita CitaReservada
	if err := First(&cita, idCita); err != nil {
		return nil, err
	}
	mascota, err := GetMascotaByID(cita.IDMascota)
	return mascota, err
}

//ValidateMascotaOwner checa si el cliente ya tiene una mascota registrada con ese nombre
func ValidateMascotaOwner(idDueno int, nombreMascota string) (*Mascota, error) {
	var mascota Mascota
	err := FirstWithCondition(&mascota, "id_dueno = ? AND nombre_mascota = ?", idDueno, nombreMascota)
	return &mascota, err
}

//Save guarda o actualiza el modelo de mascota
func (m *Mascota) Save() error {
	if m.IDMascota == 0 {
		return m.create()
	}
	return m.update()
}

func (m *Mascota) create() error {
	return Create(m)
}

func (m *Mascota) update() error {
	return Save(m)
}

//Delete elimina el modelo de la base de datos
func (m *Mascota) Delete() error {
	if m.IDMascota == 0 {
		return errors.New("No se puede eliminar este modelo de la bd porque su id es 0")
	}
	return Delete(m)
}
