 # Desafio 3 - Clean architecture

## Descrição do desafio
Criar uma aplicação aplicando conceitos de clean architecture, contendo uma api que cria Order e uma api que lista Orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

## Requisitos
- Criar as migrações necessárias para o database.
- Criar arquivo api.http com a request para criar e listar as orders.
- Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
- Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Como executar o projeto
Certifique-se de ter o docker executando em sua máquina.

Execute via terminal o arquivo docker: `docker-compose up -d`

Execute o projeto via terminal: `go run cmd/app/main.go`

Os 3 servidores estarão disponíveis nas respectivas portas:
- HTTP: 8000
- gRPC: 50051
- GraphQL: 8080

## Como acessar os servidores

### Rest http
Executar as requisições utilizando o arquivo `rest.http` que se encontra na pasta `test`.

### gRPC
Utilizar Evans para interagir com as APIs
- Instalar evans via terminal: `brew install evans`
- Executar os comandos abaixo para iniciar o servidor e acessar as APIs:
    - `evans -r repl`
    - `package pb`
    - `service OrderService`
    - `call CreateOrder` ou `call ListOrders`
    - `ctrl + d` para sair
    
### GraphQL
Acessar GraphQL Playgroud no navegador `http://localhost:8080/`
- Colar o conteúdo do arquivo `test/graphql.text`
- Criar e listar orders


## Instruções para o desenvolvedor

Para gerar código grpc, executar no terminal: `make grpc-gen`

Para gerar código graphql, executar no terminal: `make graph-gen`