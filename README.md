# 🚀 CRUD de Cadastro de Pessoas

Este projeto é um CRUD simples para cadastro de pessoas utilizando Golang, Gin e SQLite. O projeto inclui documentação com Swagger.

## 🛠️ Tecnologias Utilizadas

- **Golang** - Linguagem principal
- **Gin** - Framework web
- **SQLite** - Banco de dados
- **Swagger** - Documentação da API

## 📂 Estrutura do Projeto

```
project/
├── cmd/
│   └── main.go
├── docs/
│   └── swagger.go
├── internal/
│   ├── handlers/
│   │   └── pessoa_handler.go
│   ├── models/
│   │   └── pessoa.go
│   └── db/
│       └── database.go
├── pkg/
│   └── router/
│       └── router.go
└── go.mod
```

## 🚧 Funcionalidades

- **Listar Pessoas** ➡️ `GET /pessoas`
- **Criar Pessoa** ➡️ `POST /pessoas`
- **Atualizar Pessoa** ➡️ `PUT /pessoas/{id}`
- **Deletar Pessoa** ➡️ `DELETE /pessoas/{id}`

## 🔧 Configuração do Ambiente

1. Clone o repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   ```
2. Acesse o diretório do projeto:
   ```bash
   cd cadastroPessoa
   ```
3. Instale as dependências:
   ```bash
   go mod tidy
   ```
4. Execute o projeto:
   ```bash
   go run cmd/main.go
   ```

## 📑 Documentação com Swagger

- Acesse a documentação no navegador após rodar o projeto:
  ```
  http://localhost:8080/swagger/index.html
  ```

## 🏗️ Endpoints da API

- **GET** `/pessoas` - Lista todas as pessoas
- **POST** `/pessoas` - Cria uma nova pessoa
- **PUT** `/pessoas/{id}` - Atualiza uma pessoa pelo ID
- **DELETE** `/pessoas/{id}` - Deleta uma pessoa pelo ID
