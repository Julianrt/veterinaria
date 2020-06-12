package models

import (
	"log"

	"github.com/jinzhu/gorm"

	//sqlite is used to handle sqlite dialect with gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	//EQUAL e
	EQUAL = " = "
	//DIFFERENT e
	DIFFERENT = " <> "
	//IN e
	IN = " IN "
	//LIKE e
	LIKE = " LIKE "
	//AND e
	AND = " AND "
	//GREATER e
	GREATER = " > "
	//LESS e
	LESS = " < "
	//QUESTION e
	QUESTION = "?"
	//INQUESTION e
	INQUESTION = "(?)"
)

var db *gorm.DB

//InitDB connect to db and creates or automigrates tables
func InitDB() {
	CreateDBConn()
	autoMigrateTables()
}

//CreateDBConn inicia una conexion con la base de datos
func CreateDBConn() {
	var err error

	if getConnection() != nil {
		log.Println("a db connections already exists")
		return
	}

	db, err = gorm.Open("sqlite3", "veterinaria.db")
	if err != nil || db == nil {
		panic("it couldn't connect -> " + err.Error())
	} else {
		log.Println("db connection OK")
	}
}

//Hace una automigracion de los modelos para crear o
//modificar las tablas si hubo algun cambio en el modelo
func autoMigrateTables() {
	// Migrate the schema
	db.AutoMigrate(&Agenda{})
	db.AutoMigrate(&CitaReservada{})
	db.AutoMigrate(&Cliente{})
	db.AutoMigrate(&DetalleVenta{})
	db.AutoMigrate(&Empleado{})
	db.AutoMigrate(&HistorialConsulta{})
	db.AutoMigrate(&Mascota{})
	db.AutoMigrate(&Servicio{})
	db.AutoMigrate(&Venta{})
}

func getConnection() *gorm.DB {
	return db
}

//CloseConnection closes connection with db
func CloseConnection() {
	if err := db.Close(); err != nil {
		log.Println(err.Error())
	}
}

//Create creates a new record on a table
func Create(value interface{}) error {
	result := db.Create(value)
	return result.Error
}

//First gets a record by a condition
func First(out interface{}, where ...interface{}) error {
	result := db.First(out, where...)
	return result.Error
}

//FirstWithCondition gets a record by a condition with a specific column
func FirstWithCondition(out, query interface{}, value ...interface{}) error {
	result := db.Where(query, value...).First(out)
	return result.Error
}

//Find gets records
func Find(out interface{}, where ...interface{}) error {
	result := db.Find(out, where...)
	return result.Error
}

//FindWithCondition gets records by conditions
func FindWithCondition(out, query interface{}, value ...interface{}) error {
	result := db.Where(query, value...).Find(out)
	return result.Error
}

//Save changes on db
func Save(value interface{}) error {
	result := db.Save(value)
	return result.Error
}

//Where looks for records
func Where(query interface{}, args ...interface{}) error {
	result := db.Where(query, args...)
	return result.Error
}

//Delete deletes a record
func Delete(value interface{}, where ...interface{}) error {
	result := db.Delete(value, where...)
	return result.Error
}
