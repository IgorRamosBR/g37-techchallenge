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

Subir dependências
```bash
  docker-compose up -d
```

Rodar aplicação
```bash
  docker run -e ENVIRONMENT=dev --network="host" g37-lanches:latest
```
## Documentação
[Documentation](https://github.com/IgorRamosBR/g37-techchallenge/tree/master/api)


## Arquitetura
Arquitetura hexagonal com a estrutura de pastas baseada no [Standard Go Project Layout](https://github.com/golang-standards/project-layout#go-directories) 

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