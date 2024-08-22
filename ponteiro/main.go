package main

import "fmt"

func main() {
	var p *int

	i := 10
	p = &i
	fmt.Println(p)  // mostrar endereço aonde está alocado na memoria
	fmt.Println(*p) // mostrar o value do nosso ponteiro
}
