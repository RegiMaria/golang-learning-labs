<div align="center">
<img src="https://github.com/user-attachments/assets/41d928e3-f591-41dc-927f-160ce7c22aaa" width="800" alt="AWS Logo"/>

# Golang

**Reúne estudos, exercícios, experimentos e pequenos projetos em Go**
</div>

- Material de estudos Go/Br: [Go para Iniciantes: Primeiros Passos em 2026](https://golang.com.br/aprenda/go-para-iniciantes/)

## Objetivo

Minha jornada de aprendizado em Go (Golang): 
estudos para iniciantes, laboratórios práticos, exercícios de programação 
e pequenos projetos para explorar a linguagem e desenvolver habilidades na prática.

##  Estrutura
```
golang-learning-labs/   Apresentação e progresso
README.md
01-fundamentals/        Variáveis, tipos e funções
02-control-flow/        Condições e loops
03-data-structures/     Arrays, slices e maps
04-functions-and-methods/
05-structs-and-interfaces/
06-error-handling/
07-concurrency/          Goroutines e channels
08-mini-projects/        Projetos práticos
```
--- 


<div align="center">
<img src="https://github.com/user-attachments/assets/7f808db5-f81b-4298-882a-ba3c2d494c10" width="800" alt="AWS Logo"/>

#  Go (Golang) - Learning Roadmap

</div>

My learning journey with Go, documenting the fundamentals, practical exercises, experiments, and mini-projects I complete along the way.

> Learning Go, one commit at a time. 

## 💗 Progresso

<div align="center">

<!-- 🎀 Para atualizar: troque o número depois de /progress/ pela sua % (tópicos concluídos ÷ 10 × 100) -->
<img src="https://geps.dev/progress/5?dangerColor=93c5fd&warningColor=38bdf8&successColor=0284c7" width="420"/>

<br/><br/>
</div>




## Learning Roadmap

### 1. Go Fundamentals

| Conceito | O que aprender | Para que serve |
|---|---|---|
| Go syntax | Estrutura do código, palavras-chave e convenções | Escrever código Go simples e legível |
| Variables and types | Variáveis, constantes e tipos de dados | Armazenar e manipular informações |
| Functions | Declaração, parâmetros e retornos | Organizar e reutilizar código |
| Control flow | `if`, `switch`, `for` | Controlar a execução do programa |
| Arrays, slices and maps | Coleções de dados | Armazenar e acessar conjuntos de informações |
| Structs and methods | Estruturas e métodos | Modelar entidades e organizar comportamentos |
| Interfaces | Contratos de comportamento | Criar componentes flexíveis e desacoplados |
| Error handling | Tratamento explícito de erros | Identificar e tratar falhas de execução |
| Packages and modules | Organização e dependências | Estruturar projetos Go |

### 2. Computer Science Fundamentals

O vídeo destaca que aprender a sintaxe não é suficiente. É importante entender como o programa funciona por baixo dos panos.

| Conceito | O que aprender | Onde é utilizado |
|---|---|---|
| CPU and cores | Processador, núcleos e capacidade de execução | Entender os limites do processamento |
| Processes and threads | Processos, threads e seus ciclos de execução | Aplicações que executam múltiplas tarefas |
| Concurrency | Gerenciar várias tarefas em andamento | APIs, serviços e processamento de requisições |
| Parallelism | Executar tarefas literalmente ao mesmo tempo em diferentes recursos de processamento | Processamento intensivo e operações independentes |
| Memory management | Stack, heap e alocação de memória | Aplicações eficientes e análise de consumo |
| Garbage Collector (GC) | Gerenciamento automático de memória | Reduzir a necessidade de liberar memória manualmente |
| Race conditions | Acessos concorrentes que podem produzir resultados inconsistentes | Sistemas com tarefas compartilhando dados |
| Performance and latency | Tempo de resposta, consumo de recursos e gargalos | Sistemas de alta demanda |

### 3. Goroutines

**O que são?**

Goroutines são unidades leves de execução gerenciadas pelo runtime do Go.

**Para que servem?**

Permitem iniciar tarefas concorrentes sem precisar criar uma thread do sistema operacional para cada tarefa.

**Onde são utilizadas?**

- Processamento de várias requisições de API.
- Execução de tarefas em segundo plano.
- Consumo de mensagens de filas.
- Processamento de múltiplos arquivos.
- Chamadas simultâneas a serviços externos.

**O que estudar:**

- [ ] Como criar uma goroutine com `go`.
- [ ] Diferença entre execução sequencial e concorrente.
- [ ] Ciclo de vida das goroutines.
- [ ] Como evitar que a aplicação termine antes das tarefas.
- [ ] Como coordenar e aguardar a conclusão das tarefas.
- [ ] Como identificar vazamentos de goroutines (*goroutine leaks*).

### 4. Channels

**O que são?**

Channels são mecanismos de comunicação tipados que permitem que goroutines enviem e recebam valores.

**Para que servem?**

- Trocar dados entre goroutines.
- Sincronizar a execução de tarefas.
- Sinalizar que uma tarefa terminou.
- Distribuir trabalhos entre vários workers.
- Coordenar o processamento concorrente.

**Onde são utilizados?**

- Worker pools.
- Pipelines de processamento.
- Processamento de mensagens.
- Coleta de resultados de várias tarefas.
- Coordenação de operações concorrentes.

**O que estudar:**

- [ ] Criar channels com `make`.
- [ ] Enviar e receber valores com `<-`.
- [ ] Entender channels buffered e unbuffered.
- [ ] Fechar channels com `close`.
- [ ] Iterar sobre channels com `range`.
- [ ] Utilizar `select` para lidar com múltiplas operações de comunicação.
- [ ] Entender bloqueios e deadlocks.
- [ ] Saber quando utilizar channels e quando utilizar mutexes.

### 5. Synchronization and Race Conditions

**O que aprender:**

- [ ] Condições de corrida (*race conditions*).
- [ ] Compartilhamento de memória entre goroutines.
- [ ] `sync.WaitGroup` para aguardar tarefas.
- [ ] `sync.Mutex` e `sync.RWMutex` para proteger dados compartilhados.
- [ ] `sync.Once` para executar uma inicialização uma única vez.
- [ ] Uso de `go run -race` para detectar possíveis condições de corrida.
- [ ] Diferenças entre comunicação por channels e sincronização por locks.

**Por que isso importa?**

Duas ou mais goroutines podem acessar o mesmo dado ao mesmo tempo.
Sem a coordenação correta, o resultado pode ser inconsistente, causar
falhas ou comprometer a integridade dos dados.

### 6. Performance and Scalability

**O que aprender:**

- [ ] Benchmarking com o pacote `testing`.
- [ ] Profiling com `pprof`.
- [ ] Análise de CPU e consumo de memória.
- [ ] Alocação de memória e garbage collection.
- [ ] Identificação de gargalos.
- [ ] Medição de latência e throughput.
- [ ] Limitação de concorrência.
- [ ] Testes de carga e comportamento sob pressão.

**Onde é utilizado?**

APIs de alta volumetria, microsserviços, sistemas de pagamento, pipelines de 
dados e aplicações que precisam atender muitas requisições.

### 7. Build and Deployment

**O que aprender:**

- [ ] Compilar aplicações com `go build`.
- [ ] Executar testes com `go test`.
- [ ] Gerenciar dependências com Go Modules.
- [ ] Gerar binários para diferentes sistemas e arquiteturas.
- [ ] Criar imagens Docker para aplicações Go.
- [ ] Configurar variáveis de ambiente.
- [ ] Integrar testes e builds a pipelines de CI/CD.

**Por que isso importa?**

Go permite compilar aplicações em binários executáveis, simplificando muitos
cenários de distribuição e implantação. Ainda é necessário considerar dependências externas,
sistema operacional, arquitetura e requisitos de runtime da aplicação.

## Hands-on Labs

| Lab | Objetivo | Status |
|---|---|---|
| Hello World | Criar e executar o primeiro programa Go | ⬜ Not started |
| Variables and Types | Praticar variáveis e tipos de dados | ⬜ Not started |
| Functions and Structs | Criar funções e modelar dados | ⬜ Not started |
| Error Handling | Tratar erros explicitamente | ⬜ Not started |
| First Goroutine | Executar uma tarefa concorrente | ⬜ Not started |
| Channels Basics | Enviar e receber dados entre goroutines | ⬜ Not started |
| WaitGroup Lab | Aguardar a conclusão de múltiplas tarefas | ⬜ Not started |
| Race Condition Lab | Reproduzir e investigar uma condição de corrida | ⬜ Not started |
| Worker Pool | Distribuir tarefas entre workers | ⬜ Not started |
| HTTP API | Criar uma API HTTP simples | ⬜ Not started |
| Benchmark Lab | Medir o desempenho de uma função | ⬜ Not started |
| Dockerized Go App | Empacotar uma aplicação Go com Docker | ⬜ Not started |

## 🎯 Learning Goals

- [ ] Compreender os fundamentos da linguagem Go.
- [ ] Entender concorrência e paralelismo.
- [ ] Criar e coordenar goroutines.
- [ ] Utilizar channels para comunicação entre tarefas.
- [ ] Evitar race conditions e deadlocks.
- [ ] Desenvolver APIs e pequenos serviços.
- [ ] Investigar problemas de performance.
- [ ] Testar, compilar e empacotar aplicações Go.

---

*This repository documents my learning process. Each lab is an opportunity to practice, experiment, make mistakes, and improve my understanding of Go.*

<div align="center">
<img src="https://github.com/user-attachments/assets/498cfc54-5590-4cbb-b845-d030ca3554f4" width="200" alt="AWS Logo"/>

Se gostou do repositório deixa uma ⭐ ✨

</div>


