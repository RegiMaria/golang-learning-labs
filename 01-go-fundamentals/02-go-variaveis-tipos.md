## 1. Resumo Prático: Variáveis e Tipos

Em Go, você tem três formas principais de declarar valores:

1. Inferência de tipo (mais comum dentro de funções): O Go descobre o tipo pelo valor atribuído

```Go
nome := "Maria" // string
idade := 25     // int

```

2. Declaração explícita: Você diz exatamente o tipo da variável

```Go

var preco float64 = 19.99
var ativo bool = true
```

3. Constantes (const): Valores que nunca mudam após definidos.

```Go
    const anoAtual = 2026
	anoNascimento := anoAtual - idade
	fmt.Println("Ano de nascimento aproximado:", anoNascimento)

```
Sobre `const`

1. O const (abreviação de constant, ou constante) serve para declarar um
valor fixo que nunca poderá mudar durante a execução do programa.

2. Se você tentar alterar o valor de uma constante em qualquer lugar do código
depois que ela foi criada, o compilador do Go vai recusar e dar um erro, impedindo 
que o programa seja compilado.

Exemplo prático:
```Go
package main

import "fmt"

func main() {
    // Definindo uma constante
    const taxaJuros = 0.05
    
    // Definindo uma variável normal (que pode mudar)
    saldo := 1000.0

    fmt.Println("Saldo inicial:", saldo)
    fmt.Println("Taxa fixa:", taxaJuros)

    // Se você tentar fazer isso:
    // taxaJuros = 0.10 -> O compilador dá ERRO: "cannot assign to taxaJuros"

    // Mas a variável saldo pode mudar:
    saldo = saldo + (saldo * taxaJuros)
    fmt.Println("Saldo com juros:", saldo)
}
```
Em resumo:

- Use `var` ou `:=` para dados que vão mudar ao longo do tempo (como a pontuação de um usuário, o texto digitado num input, a hora atual).

- Use `const` para valores que são fixos e universais para a regra do seu programa (como o valor de Pi, configurações padrão, limites máximos).

### Tipos que você vai usar o tempo todo

| Tipo | Uso | Exemplo |
|---|---|---|
| `string` | Texto (UTF-8) | `"golang.com.br"` |
| `int` | Inteiro dependente da arquitetura | `42` |
| `int64` | Inteiro de tamanho fixo | `1_000_000` |
| `float64` | Número decimal | `3.14` |
| `bool` | Verdadeiro/falso | `true` |
| `error` | Resultado de operação que pode falhar | `err != nil` |


### Variáveis

| **Forma** | **Sintaxe** | **Onde funciona** |
| :--- | :--- | :--- |
| Declaração explícita | `var nome string = "Ana"` | Qualquer escopo |
| Tipo inferido | `var idade = 30` | Qualquer escopo |
| Declaração curta | `nome := "Ana"` | Apenas dentro de funções |
| Múltiplas | `var a, b int = 1, 2` | Qualquer escopo |
| Bloco `var` | `var ( ... )` | Qualquer escopo |
| Zero value | `var x int` → `0` | Qualquer escopo |