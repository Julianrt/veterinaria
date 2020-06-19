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
	return fmt.Sprintf("%d-%d-%d", fecha.Year(), fecha.Month(), fecha.Day())
}

func getTimeByObTime(hora time.Time) string {
	return fmt.Sprintf("%02d:%02d", hora.Hour(), hora.Minute())
}
