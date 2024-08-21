package main

import "fmt"

func main() {

	fmt.Print("Qual o ano atual?")
	var ano_atual int
	_, err := fmt.Scan(&ano_atual)
	if err != nil {
		fmt.Print("Aconteceu algo de errado: ", err)
		return
	}

	fmt.Printf("Qual seu ano de nascimento?")
	var ano_nascimento int

	_, err = fmt.Scan(&ano_nascimento)
	if err != nil {
		fmt.Print("Aconteceu algo de errado: ", err)
		return
	}

	calc_idade(ano_atual, ano_nascimento)
}

func calc_idade(ano_atual int, ano_nascimento int) {
	// apenas com base no ano atual...
	dif := ano_atual - ano_nascimento
	fmt.Println("Sua idade é: ", dif)
}
