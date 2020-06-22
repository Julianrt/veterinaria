package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Julianrt/veterinaria/models"
)

//GetCurrentDate f
func GetCurrentDate() time.Time {
	return time.Now()
}

//ValidateDate f
func ValidateDate(fechaCita time.Time) bool {
	now := time.Now()
	fechaAhora := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), 0, 0, time.UTC)

	if fechaAhora.Equal(fechaCita) {
		return false
	}
	if fechaAhora.After(fechaCita) {
		return false
	}
	return true
}

//FillDate f
func FillDate(fecha, tiempo string) (time.Time, error) {
	year, month, day, err := splitDate(fecha)
	if err != nil {
		return time.Time{}, err
	}
	hour, minute, err := splitTime(tiempo)
	if err != nil {
		return time.Time{}, err
	}
	date := time.Date(year, getMonth(month), day, hour, minute, 0, 0, time.UTC)
	return date, err
}

func splitDate(date string) (int, int, int, error) {
	split := strings.Split(date, "-")
	if len(split) != 3 {
		return 0, 0, 0, errors.New("formato de fecha incorrecto")
	}
	yearStr := split[0]
	monthStr := split[1]
	dayStr := split[2]

	var year, month, day int
	var err error

	if year, err = strconv.Atoi(yearStr); err != nil {
		return 0, 0, 0, err
	}
	if month, err = strconv.Atoi(monthStr); err != nil {
		return 0, 0, 0, err
	}
	if day, err = strconv.Atoi(dayStr); err != nil {
		return 0, 0, 0, err
	}

	return year, month, day, err
}

func splitTime(time string) (int, int, error) {
	split := strings.Split(time, ":")
	if len(split) != 2 {
		return 0, 0, errors.New("formato de hora incorrecto")
	}
	hourStr := split[0]
	minuteStr := split[1]

	var hour, minute int
	var err error

	if hour, err = strconv.Atoi(hourStr); err != nil {
		return 0, 0, err
	}
	if minute, err = strconv.Atoi(minuteStr); err != nil {
		return 0, 0, err
	}

	return hour, minute, err
}

func getMonth(month int) time.Month {
	if month == 1 {
		return time.January
	}
	if month == 2 {
		return time.February
	}
	if month == 3 {
		return time.March
	}
	if month == 4 {
		return time.April
	}
	if month == 5 {
		return time.May
	}
	if month == 6 {
		return time.June
	}
	if month == 7 {
		return time.July
	}
	if month == 8 {
		return time.August
	}
	if month == 9 {
		return time.September
	}
	if month == 10 {
		return time.October
	}
	if month == 11 {
		return time.November
	}
	if month == 12 {
		return time.December
	}
	return 0
}

//GetFechasOcupadas f
func GetFechasOcupadas() ([]models.Fechas, error) {
	var fechas []models.Fechas
	citas, err := models.GetCitasShowOrderByFecha()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(citas); i++ {

		tiene, indice := tieneLaFecha(fechas, citas[i])
		if !tiene {
			horas := []string{getTimeByObTime(citas[i].Fecha)}
			fechas = append(fechas, models.Fechas{
				Fecha: getDateByObTime(citas[i].Fecha),
				Horas: horas,
			})
		} else {
			fechas[indice].Horas = append(fechas[indice].Horas, getTimeByObTime(citas[i].Fecha))
		}

	}

	return fechas, nil
}

func tieneLaFecha(fechas []models.Fechas, cita models.CitaShow) (bool, int) {
	indice := 0
	for i := 0; i < len(fechas); i++ {
		if fechas[i].Fecha == getDateByObTime(cita.Fecha) {
			indice = i
			return true, indice
		}
	}
	return false, 0
}

func getDateByObTime(fecha time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", fecha.Year(), fecha.Month(), fecha.Day())
}

func getTimeByObTime(hora time.Time) string {
	return fmt.Sprintf("%02d:%02d", hora.Hour(), hora.Minute())
}

//GetFechasDisponibles f
func GetFechasDisponibles() []models.Fechas {
	fechasOcupadas, _ := GetFechasOcupadas()
	fechasDisponibles := loadTodasLasFechas()

	for i := 0; i < len(fechasOcupadas); i++ {
		tieneLasFechas(fechasDisponibles, fechasOcupadas[i])
	}
	return fechasDisponibles
}

func tieneLasFechas(fechas []models.Fechas, fechaOcupada models.Fechas) {
	for i := 0; i < len(fechas); i++ {
		if fechas[i].Fecha == fechaOcupada.Fecha {
			for j := 0; j < len(fechas[i].Horas); j++ {
				for m := 0; m < len(fechaOcupada.Horas); m++ {
					if fechas[i].Horas[j] == fechaOcupada.Horas[m] {
						fechas[i].Horas = removeIndex(fechas[i].Horas, j)
					}
				}
			}
		}
	}
}

func loadTodasLasFechas() []models.Fechas {
	var dia int
	var fechas []models.Fechas

	now := time.Now()
	mes, _ := strconv.Atoi(fmt.Sprintf("%d", now.Month()))

	for i := mes; i <= 12; i++ {
		if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
			if dia == 0 {
				fechas = append(fechas, fillDate(i, now.Day(), 31)...)
				dia++
			} else {
				fechas = append(fechas, fillDate(i, 0, 31)...)
			}
		} else if i == 2 {
			if dia == 0 {
				fechas = append(fechas, fillDate(i, now.Day(), 28)...)
				dia++
			} else {
				fechas = append(fechas, fillDate(i, 0, 28)...)
			}
		} else if i == 4 || i == 6 || i == 9 || i == 11 {
			if dia == 0 {
				fechas = append(fechas, fillDate(i, now.Day(), 30)...)
				dia++
			} else {
				fechas = append(fechas, fillDate(i, 0, 30)...)
			}
		}
	}
	return fechas
}

func fillDate(mes, dia, dias int) []models.Fechas {
	var fechas []models.Fechas
	for i := dia; i < dias; i++ {
		fechas = append(fechas, models.Fechas{fmt.Sprintf("2020-%02d-%02d", mes, i+1), hoursWorked()})
	}
	return fechas
}

func hoursWorked() []string {
	horas := []string{"08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00"}
	return horas
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
