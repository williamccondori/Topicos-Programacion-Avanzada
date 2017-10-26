package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Insertion Sort")

	// Se declara el arreglo del números.

	arreglo := []int{4, 5, 2343, 3, 7, 6, 1}

	// Se imprime el arreglo de números desordenados.

	fmt.Println(arreglo)

	// Se realiza el bucle a partir del segundo elemento hacia la derecha.

	for i := 1; i < len(arreglo); i++ {

		fmt.Println("Se evalua el número: " + strconv.Itoa(arreglo[i]))

		// Se realiza el bucle a partir del elemento actual del primer bucle y se recorre hacia la derecha.

		for j := i; j > 0; j-- {

			// Se evalúa la condición.

			if arreglo[j] < arreglo[j-1] {

				// Se realiza el cambio de posición.

				valorMenor := arreglo[j]
				valorMayor := arreglo[j-1]
				fmt.Println("# " + strconv.Itoa(valorMenor) + " es menor a " + strconv.Itoa(valorMayor))
				arreglo[j-1] = valorMenor
				arreglo[j] = valorMayor
			}
		}

		// Se imprime el arreglo ordenado.

		fmt.Println(arreglo)
	}
}
