# backend_test

## Clone o projeto
```
$ git clone https://github.com/joaoghabriell1/backend-test.git
```

## Setup


Crie a base de dados:

```
$ CREATE DATABASE backend_test;
```

Crie um arquivo: `.env` com os seguintes parâmetros :

```
[API]
PORT = 3000

[DATABASE]
DB_HOST = localhost
DB_USER = postgres
DB_PASSWORD = postgres
DB_NAME = backend_test
DB_PORT = 5432
```

### Rodando

No diretório raiz do projeto clonado:
```
$ go run main.go 
```
