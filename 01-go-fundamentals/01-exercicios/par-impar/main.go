package main

import "fmt"

func main() {
	fmt.Println("-- Verificador de par ou ímpar --")

	// Loop de 1 a 20
	for i := 1; i <=20; i++ {
		// Verifica se o resto da divisão por 2 é igual a 0
		if i % 2 == 0 {
			fmt.Printf("%d é par\n", i)
		} else {
			fmt.Printf("%d é ímpar\n", i)
		}
	}







}