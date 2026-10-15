# Go sintaxy

Antes de irmos para as variáveis, vale a pena dar uma olhada rápida
na Sintaxe do Go (Go syntax), que é a base da estrutura do código.

O Go foi desenhado para ser uma linguagem limpa,
minimalista e previsível. Aqui estão os pontos principais
sobre a sintaxe que você precisa saber logo de cara:

- Link [da Cheatsheet Go: Sintaxe Básica](https://golang.com.br/cheatsheet/sintaxe-basica/)

## Cola rápida de sintaxe Go

Os fragmentos desta tabela pressupõem que as variáveis e funções usadas
já existem; os exemplos completos aparecem nas seções correspondentes.

| Quero… | Sintaxe para consultar | Referência |
|---|---|---|
| Declarar uma variável dentro de uma função | `nome := "Ana"` | Variáveis |
| Declarar com tipo e valor zero | `var total int` | Variáveis |
| Definir uma constante | `const porta = 8080` | Constantes |
| Converter um número explicitamente | `float64(total)` | Conversões |
| Comparar e combinar condições | `total > 0 && ativo` | Operadores |
| Tratar um erro | `if err != nil { return err }` | Controle de fluxo |
| Repetir com contador | `for i := 0; i < 3; i++ { fmt.Println(i) }` | For |
| Percorrer um slice | `for i, v := range valores { fmt.Println(i, v) }` | For |
| Declarar uma função | `func soma(a, b int) int { return a + b }` | Funções |
| Acessar um valor por ponteiro | `p := &total; *p = 10` | Ponteiros |
| Adiar o fechamento de um arquivo | `defer f.Close()` | Defer |


1. Sem Ponto e Vírgula (quase sempre)
Em Go, você não usa ponto e vírgula (;) no final das linhas.
O compilador coloca o ponto e vírgula automaticamente para você nos
bastidores (através de uma regra chamada automatic semicolon insertion). 
Se você colocar, o compilador pode até reclamar.

2. Chaves Obrigatórias ({ })
Diferente de linguagens como Python onde a indentação importa para blocos de
código, o Go exige o uso de chaves {}. Além disso, a chave de abertura precisa
ficar na mesma linha da declaração (por exemplo, na função main ou num if). 
Se você **pular a chave para a linha de baixo, o compilador dá erro**.

Certo:

```Go

func main() {
    // código aqui
}

```
3. O Formatador Oficial: `go fmt`
Essa é uma das ferramentas mais amadas da comunidade Go. 
Não existe discussão sobre estilo de código em Go (onde colocar chaves, espaços, tabulações).
A própria linguagem tem um formatador oficial embutido.

Quando você roda `go fmt ./...` no terminal, ele reorganiza e formata todo o seu código
automaticamente para o padrão exato exigido pela comunidade. No VS Code, isso geralmente
acontece sozinho toda vez que você salva o arquivo (Ctrl + S).

4. Visibilidade (Letras Maiúsculas e Minúsculas)

O Go não usa palavras-chave como public, private ou protected para controlar o acesso a variáveis,
funções ou estruturas. Ele usa a regra da letra maiúscula:

- Iniciada com letra minúscula (ex: soma, idade): É privada (visível apenas dentro do mesmo pacote/pasta).

- Iniciada com letra maiúscula (ex: Soma, Println): É pública (exportada, ou seja, pode ser usada por outros pacotes ou arquivos externos).

Com essa base de sintaxe na cabeça, o código que você escreve em Go fica limpo, padronizado e fácil de ler de qualquer lugar.