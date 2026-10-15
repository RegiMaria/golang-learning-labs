# WSL -Ubuntu

## Como instalar passo a passo no WSL:

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
