package main

import (
	"fmt"
	"log"
)

func multiply(input int, initial int, limit int) {
	for i := initial; i <= limit; i++ {
		log.Printf("%d x %d é igual a: %d\n", input, i, input*i)
	}
}

func main() {
	fmt.Print("Digite um número: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Aconteceu algo de errado: ", err)
		return
	}

	fmt.Print("Tabuada de: ")
	var initial int
	_, err = fmt.Scan(&initial)
	if err != nil {
		log.Print("Aconteceu algo de errado: ", err)
	}

	fmt.Print("Tabuada até: ")
	var limit int
	_, err = fmt.Scan(&limit)
	if err != nil {
		fmt.Print("Aconteceu algo de errado: ", err)
		return
	}

	multiply(input, initial, limit)
}
