# g37-lanches

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

Serviço de controle de pedidos desenvolvido em Golang.




## Tech Stack

**API:** Go

**Infra:** PostgreSQL, SQS


## Requisitos

- go 1.20
- docker

## Como executar?

Build image

```bash
  docker build -t g37-lanches:latest .
```

Rodar aplicação
```bash
  docker-compose up -d
```
## Documentação
Após rodar a aplicação

[Documentation](https://linktodocumentation)


## Arquitetura
Estrutura de pastas baseada no [Standard Go Project Layout](https://github.com/golang-standards/project-layout#go-directories) 

```bash
├── cmd
└── internal
    |── application
    ├── core
    │   ├── domain
    │   ├── ports
    │   │   ├── in.go
    │   │   └── out.go
    │   └── services
    ├── infra
```