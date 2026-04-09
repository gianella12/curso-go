package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Println("escribe ", i, " como")

	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Toca trabajar!!")
	}

	t := time.Now()
	fmt.Println("La hora es: ", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("debes decir buenos dias")
	default:
		fmt.Println("debes decir buenas tardes")
	}

	dia := time.Now().Weekday()
	fmt.Println("Hoy es:", dia)

	switch dia {
	case time.Monday:
		fmt.Println("Hoy es Lunes")
	case time.Tuesday:
		fmt.Println("Hoy es Martes")
	case time.Wednesday:
		fmt.Println("Hoy es Miercoles")
	case time.Thursday:
		fmt.Println("Hoy es Jueves")
	case time.Friday:
		fmt.Println("Hoy es Viernes")
	case time.Saturday:
		fmt.Println("Hoy es Sabado")
	case time.Sunday:
		fmt.Println("Hoy es Domingo")
	}
}
