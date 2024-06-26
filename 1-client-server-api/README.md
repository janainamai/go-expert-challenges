# Desafio 1 - Client Server API

## Descrição do desafio
Entregar dois sistemas em Go:
- `client.go`
- `server.go`

### Requisitos do client:
- O `client.go` deverá realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.
- Utilizando o pacote context, o `client.go` terá um timeout máximo de 300ms para receber o resultado do `server.go`.
- O `client.go` precisará receber do `server.go` apenas o valor atual do câmbio (campo `bid` do JSON). 
- O `client.go` terá que salvar a cotação atual em um arquivo `quotation.txt` no formato: `Dólar: {valor}`

### Requisitos do server:
- O `server.go` deverá consumir a API contendo o câmbio de Dólar e Real no endereço:  `https://economia.awesomeapi.com.br/json/last/USD-BRL` e em seguida deverá retornar no formato JSON o resultado para o cliente.
- O timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
- Usando o package `context`, o `server.go` deverá registrar no banco de dados SQLite cada cotação recebida.
- O timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
- Se ocorrer erro de timeout, conforme configurado no context, deverá ser registrado nos logs.
- O endpoint necessário gerado pelo `server.go` para este desafio deverá ser: `/cotacao` e a porta a ser utilizada pelo servidor HTTP deverá ser a `8080`.

## Arquitetura do sistema Client

### Divisão de camadas principais e suas responsabilidades
O sistema foi dividido em três camadas principais: handler, core e infra.
- Handler: camada que processa a entrada de dados que atualmente são requisições de protocolo http.
- Core: camada responsável pela lógica de negócio, recebe dados do handler e interage com a camada de infraestrutura através de contratos de interface definidos no package `ports`.
- Infra: camada que orquestra a interação do sistema com infrastruturas externas, atualmente interage com um database e uma aplicação através de protocolo http.
  - Database SQLITE: registra cotações que foram obtidas pelo sistema.
  - Aplicação `economia.awesomeapi.com.br`: retorna a cotação atual do Dólar.

## Executando o sistema localmente
1. Executar a aplicação `server.go`
   1. No terminal, na pasta `server`, executar o comando `go run app/main.go`.
2. Executar a aplicação `client.go`
   1. No terminal, na pasta `client`, executar o comando `go run app/main.go`.

### Validando funcionamento do `client`
1. Abrir o arquivo `quotation.txt` e verificar se consta o valor da cotação do Dolar.
   1. Exemplo válido: `Dólar: 5.4232

### Validando funcionamento do `server`
Acessar url de auditoria para obter os registros atuais que existem no `database` do `server`. É possível escolher o formato de saída desejado para visualizar os registros salvos:
- Json Http: executar url `http://localhost:8080/cotacao/audit`
- HTML: executar url `http://localhost:8080/cotacao/audit?format=html`
- Console: executar url `http://localhost:8080/cotacao/audit?format=console`
