package main

import "fmt"

func main() {
	fmt.Println("Digite um número para converter: ")
	var temperatura float64
	_, err := fmt.Scan(&temperatura)
	if err != nil {
		fmt.Println("Aconteceu algo de errado: ", err)
		return
	}

	to_fahrenheit(temperatura)
	to_kelvin(temperatura)

}

func to_fahrenheit(number float64) {
	// Multiplicamos a temperatura em ºC por 1,8 e somamos 32 ao resultado.
	result := number*1.8 + 32
	fmt.Printf("A temperatura convertida de ºC para Fahrenheit é: %.2f\n", result)
}

func to_kelvin(number float64) {
	var base float64 = 273.15
	result := number + base
	fmt.Printf("A temperatura convertida de ºC para Kelvin é: %.2f\n", result)
}
