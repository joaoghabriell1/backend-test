# backend_test

## Instalação
```
$ git clone https://github.com/joaoghabriell1/backend-test.git
```

## How to run

### Ferramentas Necessárias

- PostgreSQL

### Ready

Crie a base de dados:

```
$ CREATE DATABASE backend_test;
```

### Dotenv

Modifique o arquivo: `./.env`

```
[API]
PORT = 3000

[Database]
DB_HOST = localhost
DB_USER = postgres
DB_PASSWORD = postgres
DB_NAME = backend_test
DB_PORT = 5432
...
```

### Rodando
```
$ go run main.go 
```


## Tecnologias

- Gin
- Gorm
