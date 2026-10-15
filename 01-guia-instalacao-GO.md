# WSL -Ubuntu

## 1. Como instalar passo a passo no WSL:

Abra o seu terminal do WSL (Ubuntu, Debian, etc.) e execute os comandos abaixo:

1. Baixe o arquivo correto:

```Bash
curl -LO https://golang.org/dl/go1.27.2.linux-amd64.tar.gz
```
2. Remova qualquer versão antiga (se houver) e extraia o novo arquivo para /usr/local:

```Bash
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.27.2.linux-amd64.tar.gz
```
3. Adicione o Go ao seu PATH (para que o terminal saiba onde encontrar o comando go):
```Bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
4. Confirme se deu tudo certo:

```Bash
go version
```
Documentação [Go/BR - GO para iniciantes](https://golang.com.br/aprenda/go-para-iniciantes/)

---

## 2. Crie o módulo e o Hello World

1. Crie a pasta do projeto e entre nela
```Bash
mkdir hello-go && cd hello-go
```

2. Inicialize o módulo Go
O `go mod init` cria o arquivo `go.mod`, que gerencia as dependências do seu projeto:

```Bash
go mod init exemplo.com/hello
```

3. Crie `main.go`:

```Bash

package main

import "fmt"

func main() {
	fmt.Println("Olá, Go!")
}
```
Deve ser criada na mesma pasta onde está o arquivo `go.mod`
(ou seja, dentro da pasta hello-go)

A raiz do seu projeto Go é justamente a pasta que contém o `go.mod`.
O compilador do Go usa a presença do go.mod para saber onde o módulo
começa e onde buscar os arquivos `.go`.

Portanto, a estrutura da sua pasta hello-go deve ficar assim:

```
hello-go/
├── go.mod
└── main.go
```

4. Execute
`go run . `
# Olá, Go!

5.Ou compile um binário:
```Bash

go build -o hello .
./hello

```
Quando você executAR `go build -o hello .`, o Go compilará
seu código e criará um arquivo executável chamado `hello.`
Depois, ao executar `./hello`, você rodará esse programa.


---

## 3. O que cada linha significa

| **Trecho** | **Papel** |
| :--- | :--- |
| **`package main`** | Marca um executável (não uma biblioteca) |
| **`import "fmt"`** | Pacote da stdlib para formatar e imprimir |
| **`func main()`** | Ponto de entrada do programa |
| **`fmt.Println(...)`** | Escreve uma linha no stdout |