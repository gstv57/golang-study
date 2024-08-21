package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Começando jogo..... \n")
	var number_random int = rand.Intn(5)

	fmt.Print("Número aleátorio gerado com sucesso!! Tente adivinhar \n")
	var input_usuario int

	for {
		_, err := fmt.Scan(&input_usuario)

		if err != nil {
			fmt.Print("Aconteceu alguma coisa de errado: ", err)
			return
		}

		if input_usuario < number_random {
			fmt.Print("O número aleátorio é maior.")
		}
		if input_usuario > number_random {
			fmt.Print("O número aleátorio é menor.")
		}
		if input_usuario == number_random {
			fmt.Printf("Parabens!!! Você acertou, o número era %d!", number_random)
		}

	}
}
