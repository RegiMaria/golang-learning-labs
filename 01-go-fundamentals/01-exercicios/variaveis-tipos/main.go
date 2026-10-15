package main

import (
	"fmt"
)
func main() {
	// 1. Declaração de variáveis:
	// - Nome
	// - Idade
	// - Altura
	// - Peso
	// - Ativo (booleano)

	// Exemplo

	nome := "João"
	idade := 30
	altura := 1.75
	peso := 70.5
	ativo := true

	// 2. Impressão das variáveis na tela usando fmt.Println ou fmt.Printf
	fmt.Println("-- Informações do usuário --")
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Peso:", peso)
	fmt.Println("Ativo:", ativo)

	// // 3. Desafio extra: Declare uma constante com o ano atual (2026) 
	// e calcule o seu ano de nascimento aproximado.
	const anoAtual = 2026
	anoNascimento := anoAtual - idade
	fmt.Println("Ano de nascimento aproximado:", anoNascimento)
}

// Sobre const
// Use var ou := para dados que vão mudar ao longo do tempo (como a pontuação de um usuário, o texto digitado num input, a hora atual).
// Use const para valores que são fixos e universais para a regra do seu programa (como o valor de Pi, configurações padrão, limites máximos).