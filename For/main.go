package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // se puede hacer esto o usar i += 1 o i = i + 1
		/*
			comentarios de multiples lineas
		*/
	}
	fmt.Println("...............")
	for numero := 0; numero < 3; numero++ {
		fmt.Println(numero)
	}

	fmt.Println("...............")
	for rango := range 3 {
		fmt.Println("rango: ", rango)
	}

	fmt.Println("...............")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("...............")
	for valor := range 6 {
		if valor%2 == 0 {
			continue
		}
		fmt.Println(valor, ": Valores inpar")
	}
}
