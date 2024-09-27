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

### Como executar o projeto:
- Certifique-se de ter o docker executando em sua máquina.
- Execute via terminal o arquivo docker: `docker-compose up -d`
- Executar o projeto via terminal: `go run cmd/app/main.go`
- Testar a aplicação utilizando as 3 formas:
    - Rest:
    - gRPC:
    - GraphQL: