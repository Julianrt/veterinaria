package models

//Empleado estructura para manejar el modelo de empleados
type Empleado struct {
	IDEmpleado int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nombre     string `gorm:""`
	Direccion  string `gorm:""`
	Telefono   string `gorm:""`
}
