package main

import (
	"fmt"
	"strconv"
)

func main() {
	numero1 := 10
	numero2 := 20

	fmt.Println(numero1, numero2)
	numeroDoble := 20.5
	resultado := numero2 + int(numeroDoble)
	fmt.Println(resultado)

	nombre := "Gianella"
	apellido := "Lastra"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	edad := 21
	edadYnombre := nombreCompleto + " su edad: " + strconv.Itoa(edad)
	fmt.Println(edadYnombre)

}
