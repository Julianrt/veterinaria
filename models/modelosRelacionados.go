package models

import "time"

//CitaShow modelo para mostrar informacion de las citas
//con una llamada relacionando tablas en la base de datos
type CitaShow struct {
	IDCita        int
	NombreDueno   string
	NombreMascota string
	Telefono      string
	Correo        string
	Fecha         string
}

//CitasShow una lista de citaShow
type CitasShow []CitaShow

//GetCitasShowOrderByFecha obtiene las citas relacionando las tablas
//y las ordena por la fecha
func GetCitasShowOrderByFecha() (*CitasShow, error) {
	var citas CitasShow
	tabla := "cita_reservadas"
	columns := "cita_reservadas.id_cita, clientes.nombre_dueno, mascota.nombre_mascota, "
	columns += "clientes.telefono, clientes.correo, cita_reservadas.fecha"
	inner1 := "INNER JOIN clientes ON cita_reservadas.id_dueno = clientes.id_dueno"
	inner2 := "INNER JOIN mascota ON cita_reservadas.id_mascota = mascota.id_mascota"
	conditions := "ORDER BY cita_reservadas.fecha"

	err := FindJoins(tabla, columns, inner1+" "+inner2+" "+conditions, &citas)
	return &citas, err
}

//ConsultaShow modelo para mostrar informacion de las consultas
//con una llamada relacionando tablas en la base de datos
type ConsultaShow struct {
	NombreDueno   string
	NombreMascota string
	TipoAnimal    string
	Edad          int
	Peso          float32
	Vacunas       string
	Prescripcion  string
	FechaConsulta time.Time
}

//ConsultasShow una lista de consultaShow
type ConsultasShow []ConsultaShow

//GetConsultasShowOrderDesc obtiene las consultas relacionando las tablas
func GetConsultasShowOrderDesc() (ConsultasShow, error) {
	var consultas ConsultasShow
	tabla := "historial_consulta"

	columns := "clientes.nombre_dueno, mascota.nombre_mascota, mascota.tipo_animal, "
	columns += "mascota.edad, mascota.peso, mascota.vacunas, "
	columns += "historial_consulta.prescripcion, historial_consulta.fecha_consulta"

	join := "INNER JOIN clientes ON historial_consulta.id_dueno = clientes.id_dueno "
	join += "INNER JOIN mascota ON historial_consulta.id_mascota = mascota.id_mascota "

	conditions := "ORDER BY historial_consulta.fecha_consulta DESC"

	err := FindJoins(tabla, columns, join+conditions, &consultas)
	return consultas, err
}
