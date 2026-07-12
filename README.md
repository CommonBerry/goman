# Goman — Gerenciador de Templates e Aliases

Goman é uma API REST leve escrita em Go (Fiber) para gerenciar templates e aliases usados pela CLI `gg`. Permite armazenar templates reutilizáveis de projetos Go e criar aliases de pacotes, separando dados da lógica da CLI.

## Sumário

- [Visão Geral](#visão-geral)
- [Recursos](#recursos)
- [Pré-requisitos](#pré-requisitos)
- [Início Rápido](#início-rápido)
- [Variáveis de Ambiente](#variáveis-de-ambiente)
- [Endpoints da API](#endpoints-da-api)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Desenvolvimento](#desenvolvimento)
- [Docker](#docker)
- [Notas de Segurança](#notas-de-segurança)
- [Licença](#licença)
- [Contribuindo](#contribuindo)
- [Suporte](#suporte)

---

## Visão Geral

Goman atua como repositório central para templates de projetos e aliases consumidos pela CLI `gg`. Isso permite atualizar templates e aliases sem aumentar o binário ou forçar reinstalações da CLI.

Principais casos de uso:

- Templates de projeto (estruturas repetíveis)
- Aliases para pacotes/boilerplate
- Resolução rápida pela CLI via consultas HTTP

---

## Recursos

- API REST construída com Fiber
- Persistência em PostgreSQL
- Endpoint de health check
- Autenticação por chave (header `X-API-Key`) para operações protegidas
- Resposta em JSON
- Suporte a Docker / Podman
- Modo desenvolvimento com hot reload (Air)

---

## Pré-requisitos

- Go 1.26.4+
- PostgreSQL 12+
- Make
- (Opcional) Docker & Docker Compose ou Podman

---

## Início Rápido

1. Clone o repositório:

```bash
git clone https://github.com/CommonBerry/goman.git
cd goman
```

2. Copie e ajuste as variáveis de ambiente:

```bash
cp .env.example .env
# Edite .env conforme necessário
```

Exemplo mínimo de `.env`:

```env
PORT=3000
GOMAN_ADMIN_KEY=minha-chave-secreta-super-segura
DATABASE_URL=postgres://goman:goman@localhost:5432/goman?sslmode=disable
POSTGRES_USER=goman
POSTGRES_PASSWORD=goman
POSTGRES_DB=goman
```

3. Iniciar PostgreSQL (Makefile):

```bash
make postgres-up
```

4. Executar o servidor:

```bash
make run
```

A API estará disponível em: http://localhost:3000

5. Verificar health:

```bash
curl http://localhost:3000/health
```

Resposta esperada:

```json
{
  "status": "UP",
  "timestamp": "2026-07-12T15:03:22.210Z",
  "checks": { "postgres": "healthy" }
}
```

---

## Variáveis de Ambiente

| Variável | Padrão | Descrição |
|---|---:|---|
| `PORT` | `3000` | Porta do servidor |
| `GOMAN_ADMIN_KEY` | — | Chave API para endpoints protegidos (obrigatória) |
| `DATABASE_URL` | — | String de conexão PostgreSQL (obrigatória) |
| `POSTGRES_USER` | `goman` | Usuário do banco (usado no Docker Compose) |
| `POSTGRES_PASSWORD` | `goman` | Senha do banco (usado no Docker Compose) |
| `POSTGRES_DB` | `goman` | Nome do banco (usado no Docker Compose) |

---

## Endpoints da API

Todos os endpoints retornam JSON. Endpoints protegidos exigem o header `X-API-Key`.

### Health

- GET /health — verifica disponibilidade e dependências.

### Templates

- GET /templates — lista templates
- GET /templates/:name — busca template por nome
- POST /templates — cria template (protegido)
- PUT /templates/:uuid — atualiza template (protegido)
- DELETE /templates/:uuid — remove template (protegido)

Exemplos:

Listar templates:

```bash
curl http://localhost:3000/templates
```

Criar template (protegido):

```bash
curl -X POST http://localhost:3000/templates \
  -H "X-API-Key: sua-chave-admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "rest_api_simples",
    "content": "{\"estrutura\": \"template de API REST básica\"}"
  }'
```

Resposta de exemplo (201):

```json
{
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "name": "rest_api_simples",
  "content": "{\"estrutura\": \"template de API REST básica\"}",
  "created_at": "2026-07-12T15:00:00Z",
  "updated_at": "2026-07-12T15:00:00Z"
}
```

### Aliases

- GET /aliases — lista aliases
- GET /aliases/:name — busca alias por nome
- POST /aliases — cria alias (protegido)
- PUT /aliases/:uuid — atualiza alias (protegido)
- DELETE /aliases/:uuid — remove alias (protegido)

Exemplo de criação:

```bash
curl -X POST http://localhost:3000/aliases \
  -H "X-API-Key: sua-chave-admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "api_simples",
    "template_name": "rest_api_simples"
  }'
```

Resposta de exemplo (201):

```json
{
  "uuid": "660e8400-e29b-41d4-a716-446655440001",
  "name": "api_simples",
  "template_name": "rest_api_simples",
  "created_at": "2026-07-12T15:00:00Z",
  "updated_at": "2026-07-12T15:00:00Z"
}
```

---

## Estrutura do Projeto

```text
goman/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── routes/
│       ├── routes.go
│       ├── auth.go
│       └── health.go
├── internal/
│   ├── core/
│   │   ├── types.go
│   │   └── database.go
│   └── infra/
│       └── postgres.go
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── Dockerfile.dev
├── Makefile
├── go.mod
└── README.md
```

Componentes principais:

- `cmd/api/main.go` — inicialização da aplicação e conexão com DB
- `cmd/routes/*` — definição de rotas e middlewares
- `internal/core` — tipos de domínio e interface do repositório
- `internal/infra` — implementação PostgreSQL

---

## Desenvolvimento

Comandos úteis (Makefile):

```bash
make build       # build do binário
make run         # executar servidor de desenvolvimento
make watch       # hot reload (Air)
make air-install # instala Air
make test        # roda testes
make fmt         # go fmt
make tidy        # go mod tidy
make clean       # limpa artefatos
```

Executando localmente:

```bash
make postgres-up      # levanta Postgres
make run              # executa a API
# ou
make watch             # desenvolvimento com hot reload
```

---

## Docker

Pré-requisitos: Docker & Docker Compose ou Podman.

```bash
cp .env.example .env
make docker-up
# logs
make docker-logs
# parar
make docker-down
```

Para usar Podman:

```bash
COMPOSE="podman compose" make docker-up
```

Modo watch com Docker:

```bash
make docker-watch-up
make docker-watch-logs
make docker-watch-down
```

---

## Notas de Segurança

- Proteja os endpoints sensíveis com uma chave forte (`GOMAN_ADMIN_KEY`).
- Nunca versionar o arquivo `.env` (inclua no `.gitignore`).
- Use HTTPS e rate limiting em produção.

---

## Licença

Este repositório faz parte da iniciativa CommonBerry e é distribuído sob a [licença](LICENSE) GPLv3.

---

## Contribuindo

Contribuições são bem-vindas. Ao enviar PRs, por favor:

- Formate o código: `make fmt`
- Verifique possíveis issues com: `go vet ./...`
- Adicione testes quando pertinente

---

## Suporte

Abra uma issue no GitHub para dúvidas, bugs ou sugestões.
