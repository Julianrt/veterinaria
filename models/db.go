package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//InitDB inicia una conexion con la base de datos y
//Hace una automigracion de los modelos para crear o
//modificar las tablas si hubo algun cambio en el modelo
func InitDB() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	} else {
		log.Println("db connection OK")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Agenda{})
	db.AutoMigrate(&Cita{})
	db.AutoMigrate(&Cliente{})
	db.AutoMigrate(&DetalleVenta{})
	db.AutoMigrate(&Empleado{})
	db.AutoMigrate(&Historial{})
	db.AutoMigrate(&Venta{})

	db.Create(&Empleado{Nombre: "nombre1", Direccion: "Direccion1", Telefono: "1234567890"})
	db.Create(&Empleado{Nombre: "nombre3", Direccion: "Direccion3", Telefono: "1234567890"})
	db.Create(&Empleado{Nombre: "nombre2", Direccion: "Direccion2", Telefono: "1234567890"})

	var empleados []Empleado
	db.Find(&empleados)
	log.Println(empleados)
}
