 # Desafio 4 - Sistema de clima com cloud run

## Descrição do desafio
Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

## Requisitos

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- Deverá ser realizado o deploy no Google Cloud Run.
- O sistema deve responder adequadamente nos seguintes cenários:
    - Em caso de sucesso:
        - Código HTTP: 200
        - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
    - Em caso de falha, caso o CEP não seja válido (com formato correto):
        - Código HTTP: 422
        - Mensagem: invalid zipcode
    - ​​​Em caso de falha, caso o CEP não seja encontrado:
        - Código HTTP: 404
        - Mensagem: can not find zipcode

## Requisitos de entrega
- O código-fonte completo da implementação.
- Testes automatizados demonstrando o funcionamento.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

## Dicas para o desenvolvimento
- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
    - Sendo F = Fahrenheit
    - Sendo C = Celsius
    - Sendo K = Kelvin

## Como executar o projeto
Certifique-se de ter o docker executando em sua máquina.

Execute via terminal o arquivo docker: `docker-compose up -d`

Execute o projeto via terminal: `go run cmd/app/main.go`

Servidores disponíveis
- HTTP: 8000

## Como executar as requisições

### Rest http
Executar as requisições utilizando o arquivo `rest.http` que se encontra na pasta `test`.

## Como testar o serviço no cloud run
https://labs-temperature-534419467934.us-central1.run.app/temperature?zipcode=06233903