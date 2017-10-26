package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Selection Sort")

	// Se declara el arreglo del números.

	arreglo := []int{4, 5, 2343, 3, 7, 6, 1}

	// Se imprime el arreglo de números desordenados.

	fmt.Println(arreglo)

	// Se realiza el bucle a partir del primer elemento hacia la derecha.

	for i := 0; i < len(arreglo); i++ {

		fmt.Println("Se evalua el número: " + strconv.Itoa(arreglo[i]))

		elemento := arreglo[i]

		minValor := elemento
		minPosicion := i

		for j := i + 1; j < len(arreglo); j++ {
			if minValor > arreglo[j] {
				minValor = arreglo[j]
				minPosicion = j
			}
		}

		arreglo[minPosicion] = elemento
		arreglo[i] = minValor

		fmt.Println("# Se encontró el mínimo " + strconv.Itoa(minValor) + " [" + strconv.Itoa(minPosicion) + "]")
		fmt.Println("# " + strconv.Itoa(elemento) + " cambia de lugar con " + strconv.Itoa(minValor))

		// Se imprime el arreglo ordenado.

		fmt.Println(arreglo)
	}
}
