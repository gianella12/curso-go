package main

import "fmt"

func main() {

	nombre := "Giane"
	edad := 20

	if nombre == "Giane" {
		fmt.Println("Hola Giane")
	} else {
		fmt.Println("Hola desconocido..")
	}

	if edad >= 18 {
		fmt.Println("Ya puedes votar")
	}

	if 8%2 == 0 {
		fmt.Println("el 8 es par")
	} else {
		fmt.Println("el 8 no es par")
	}

	if numero := 9; numero < 0 {
		fmt.Println(numero, " es negativo")
	} else if numero < 10 {
		fmt.Println(numero, " es un digito")
	} else {
		fmt.Println(numero, " es un numero mayor a 9")
	}
}
