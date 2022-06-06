# solutiondelivery

O serviço criado irá realizar a persistênica dos dados do arquivo base_teste.txt disponibilzado e levantará uma API.

## Sumário

- [Requisitos](#requisitos)
- [Execusão](#execusão)
- [Credenciais da aplicação](#credenciais-da-aplicação)
- [Estrutura do Banco de Dados](#estrutura-do-banco-de-dados)
- [To-Do](#to-do)

## Requisitos

- [Docker](https://docs.docker.com/);
- [Docker compose](https://docs.docker.com/compose/);
- [GUI de banco de dados];

## Execusão

1. Realizar build da aplicação:
```bash
docker compose build 
```
2. Executar o projeto:
```bash
docker compose up
 ```
3. Executar o projeto: em segundo plano:
````
docker compose up -d
````
## Credenciais da aplicação

  - Credenciais do banco de dados:
    - **DB name** = solutiondelivery;
    - **DB user** = postgres;
    - **DB password** = localdev;
    - **DB host** = admin;
    - **DB port** = 5432;
    - **DB local access url** =  http://localhost:5432.

## Estrutura do Banco de Dados
 
Column               |     Type      | Description
-------------------- | ------------- | -------------------------------------
id                   | INTEGER       | Identificador (chave primária).
cpf                  | VARCHAR       | Documento do cliente.
private              | BOOLEAN       | Verifica se o cliente é privado.
incompleto           | BOOLEAN       | Verifica se é incompleto.
data_ultima_compra   | DATE          | Data da última compra.  
ticket_medio         | VARCHAR       | valor médio gato.
ticket_ultima_compra | VARCHAR       | Valor da última compra.
loja_mais_frequente  | VARCHAR       | CNPJ da loja mais frequentada.
loja_ultima_compra   | VARCHAR       | CNPJ da loja que foi realizado a última compra.
cpf_valido           | BOOLEAN          | Validação se o CPF é válido.
cnpj_valido          | BOOLEAN          | Validação se o CNPJ da loja mais frequente é válido.
cnpj_ult_compra_valido | BOOLEAN        | Validação se o CNPJ da loja da última cimpra é válido.

Método    | URL   
--------- | ------
GET     | http://localhost:8080/clientes
GET     | http://localhost:8080/cliente/:id 
POST    | http://localhost:8080/cliente - {"cpf": "000.000.000-00","private": "1","incompleto": "0","data_ultima_compra": "2012-03-06","ticket_medio":"1564.00","ticket_ultima_compra":"","loja_mais_frequente":"","loja_ultima_compra":""} (json exemplo para salvar)
PUT     | http://localhost:8080/cliente/:id
DELETE  | http://localhost:8080/cliente/:id

## To-Do

1. Adicionar integração contínua na plataforma do repositório com testes;
2. Melhorar o tempo de processamento de carregamento e persistênica dos dados.
