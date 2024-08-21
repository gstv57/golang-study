package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digite um número para saber se ele é primo: ")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print("Aconteceu algo de errado: ", err)
		return
	}

	isPrimo(number)
}

func isPrimo(number int) {
	if number <= 1 {
		fmt.Println("Números menores ou iguais a 1 não são primos.")
		return
	}

	if number == 2 || number == 3 {
		fmt.Printf("%d é um número primo.\n", number)
		return
	}

	if number%2 == 0 {
		fmt.Printf("%d não é um número primo.\n", number)
		return
	}

	raizQuadrada := int(math.Sqrt(float64(number)))

	for i := 3; i <= raizQuadrada; i += 2 {
		if number%i == 0 {
			fmt.Printf("%d não é um número primo.\n", number)
			return
		}
	}

	fmt.Printf("%d é um número primo!\n", number)
}
