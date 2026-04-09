package main

import "fmt"

func main() {
	var arreglo [5]int
	fmt.Println("Arreglo completo: ", arreglo)

	arreglo[4] = 100
	fmt.Println("Arreglo completo: ", arreglo)
	fmt.Println("Arreglo en la posicion 4: ", arreglo[4])

	fmt.Println("Tamaño del arreglo: ", len(arreglo))

	lista := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original: ", lista)

	// los ... indican a go que es el quien declare el tamaño del arreglo no es el programador ni el usuario
	lista2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Inferido: ", lista2)
	fmt.Println("Tamaño del arreglo inferido por go: ", len(lista2))
}
