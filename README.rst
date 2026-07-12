Goman - Gerenciador de Templates e Aliases
==========================================

**Goman** é uma API REST leve construída com Go e Fiber que gerencia templates e aliases para a CLI **gg**. Armazene templates reutilizáveis de projetos Go e crie aliases para pacotes, tudo centralizado em um único repositório.

---

Sumário
-------

* `Visão Geral <#visão-geral>`_
* `Recursos <#recursos>`_
* `Pré-requisitos <#pré-requisitos>`_
* `Início Rápido <#início-rápido>`_
* `Variáveis de Ambiente <#variáveis-de-ambiente>`_
* `Endpoints da API <#endpoints-da-api>`_
* `Estrutura do Projeto <#estrutura-do-projeto>`_
* `Desenvolvimento <#desenvolvimento>`_
* `Docker <#docker>`_

---

Visão Geral
-----------

Goman é o repositório centralizado de templates e aliases para a **CLI gg** - um wrapper poderoso da CLI do Go que adiciona funcionalidades avançadas:

* **Templates de Projetos**: Armazene estruturas de projetos Go reutilizáveis com configurações pré-definidas
* **Aliases de Pacotes**: Crie atalhos para pacotes Go frequentemente utilizados
* **Lookups Rápidos**: A CLI gg consulta Goman para resolver templates e aliases instantaneamente

Por que Goman?
~~~~~~~~~~~~~~

A CLI gg precisava de uma forma centralizada e escalável para gerenciar templates e aliases sem aumentar o tamanho do binário. Goman oferece:
* Gerenciamento remoto de templates e aliases
* Atualizações em tempo real sem reinstalar a CLI
* Separação clara entre lógica da CLI (gg) e dados (Goman)

---

Recursos
--------

* Compatível com API REST usando framework Fiber
* Compatível com persistência em banco de dados PostgreSQL
* Compatível com endpoint de verificação de saúde
* Compatível com autenticação por chave API (header ``X-API-Key``)
* Compatível com request/response em JSON
* Compatível com suporte Docker (compatível com Podman)
* Compatível com modo desenvolvimento com hot reload via Air

---

Pré-requisitos
--------------

Antes de executar Goman, certifique-se de ter:

* **Go** 1.26.4 ou superior
* **PostgreSQL** 12 ou superior
* **Make** (para comandos convenientes)
* **Docker & Docker Compose** (ou **Podman** - defina ``COMPOSE=podman compose`` no Makefile)

---

Início Rápido
-------------

#. Clonar o Repositório
~~~~~~~~~~~~~~~~~~~~~~~

`````bash
git clone https://github.com/CommonBerry/goman.git
cd goman
`````

#. Configurar Variáveis de Ambiente
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Copie o arquivo de exemplo:

`````bash
cp .env.example .env
`````

Edite ``.env`` e defina seus valores:

`````env
PORT=3000
GOMAN*ADMIN*KEY=sua-chave-secreta-forte
DATABASE*URL=postgres://goman:goman@localhost:5432/goman?sslmode=disable
POSTGRES*USER=goman
POSTGRES*PASSWORD=goman
POSTGRES*DB=goman
`````

#. Iniciar PostgreSQL
~~~~~~~~~~~~~~~~~~~~~

`````bash
make postgres-up
`````

#. Executar o Servidor
~~~~~~~~~~~~~~~~~~~~~~

`````bash
make run
`````

A API estará disponível em ``http://localhost:3000``

#. Verificar Saúde
~~~~~~~~~~~~~~~~~~

`````bash
curl http://localhost:3000/health
`````

Resposta esperada:
`````json
{
  "status": "UP",
  "timestamp": "2026-07-12T15:03:22.210Z",
  "checks": {
    "postgres": "healthy"
  }
}
`````

---

Variáveis de Ambiente
---------------------

| Variável | Padrão | Descrição |
|----------|--------|-----------|
| ``PORT`` | ``3000`` | Porta do servidor |
| ``GOMAN*ADMIN*KEY`` | - | Chave API para endpoints protegidos (obrigatória) |
| ``DATABASE*URL`` | - | String de conexão PostgreSQL (obrigatória) |
| ``POSTGRES*USER`` | ``goman`` | Usuário do banco (para docker-compose) |
| ``POSTGRES*PASSWORD`` | ``goman`` | Senha do banco (para docker-compose) |
| ``POSTGRES*DB`` | ``goman`` | Nome do banco (para docker-compose) |

**Exemplo de arquivo ``.env``:**

`````env
PORT=3000
GOMAN*ADMIN*KEY=minha-chave-secreta-super-segura
DATABASE*URL=postgres://goman:goman@localhost:5432/goman?sslmode=disable
POSTGRES*USER=goman
POSTGRES*PASSWORD=goman
POSTGRES*DB=goman
`````

---

Endpoints da API
----------------

Todos os endpoints retornam responses em JSON. Endpoints protegidos requerem o header ``X-API-Key``.

Verificação de Saúde
~~~~~~~~~~~~~~~~~~~~

#### ``GET /health``

Verifica a saúde da API e do banco de dados.

**Response (200):**
`````json
{
  "status": "UP",
  "timestamp": "2026-07-12T15:03:22.210Z",
  "checks": {
    "postgres": "healthy"
  }
}
`````

---

Templates
~~~~~~~~~

#### ``GET /templates``

Lista todos os templates disponíveis.

**Exemplo:**
`````bash
curl http://localhost:3000/templates
`````

**Response (200):**
`````json
[
  {
    "uuid": "550e8400-e29b-41d4-a716-446655440000",
    "name": "rest*api*simples",
    "content": "{\"estrutura\": \"template de API REST básica\"}",
    "created*at": "2026-07-12T15:00:00Z",
    "updated*at": "2026-07-12T15:00:00Z"
  }
]
`````

---

#### ``GET /templates/:name``

Busca um template pelo nome.

**Exemplo:**
`````bash
curl http://localhost:3000/templates/rest*api*simples
`````

**Response (200):**
`````json
{
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "name": "rest*api*simples",
  "content": "{\"estrutura\": \"template de API REST básica\"}",
  "created*at": "2026-07-12T15:00:00Z",
  "updated*at": "2026-07-12T15:00:00Z"
}
`````

**Response (404):** Template não encontrado

---

#### ``POST /templates`` ⚠️ Protegido

Cria um novo template.

**Headers:**
`````
X-API-Key: sua-chave-admin
Content-Type: application/json
`````

**Corpo do Request:**
`````json
{
  "name": "rest*api*simples",
  "content": "{\"estrutura\": \"template de API REST básica\"}"
}
`````

**Exemplo:**
`````bash
curl -X POST http://localhost:3000/templates \
  -H "X-API-Key: sua-chave-admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "rest*api*simples",
    "content": "{\"estrutura\": \"template de API REST básica\"}"
  }'
`````

**Response (201):**
`````json
{
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "name": "rest*api*simples",
  "content": "{\"estrutura\": \"template de API REST básica\"}",
  "created*at": "2026-07-12T15:00:00Z",
  "updated*at": "2026-07-12T15:00:00Z"
}
`````

---

#### ``PUT /templates/:uuid`` ⚠️ Protegido

Atualiza um template existente.

**Headers:**
`````
X-API-Key: sua-chave-admin
Content-Type: application/json
`````

**Corpo do Request:**
`````json
{
  "name": "rest*api*simples",
  "content": "{\"estrutura\": \"template atualizado\"}"
}
`````

**Exemplo:**
`````bash
curl -X PUT http://localhost:3000/templates/550e8400-e29b-41d4-a716-446655440000 \
  -H "X-API-Key: sua-chave-admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "rest*api*simples",
    "content": "{\"estrutura\": \"template atualizado\"}"
  }'
`````

**Response (200):**
`````json
{
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "name": "rest*api*simples",
  "content": "{\"estrutura\": \"template atualizado\"}",
  "created*at": "2026-07-12T15:00:00Z",
  "updated*at": "2026-07-12T15:03:22Z"
}
`````

---

#### ``DELETE /templates/:uuid`` ⚠️ Protegido

Deleta um template.

**Headers:**
`````
X-API-Key: sua-chave-admin
`````

**Exemplo:**
`````bash
curl -X DELETE http://localhost:3000/templates/550e8400-e29b-41d4-a716-446655440000 \
  -H "X-API-Key: sua-chave-admin"
`````

**Response (204):** Sem conteúdo (sucesso)

---

Aliases
~~~~~~~

#### ``GET /aliases``

Lista todos os aliases disponíveis.

**Exemplo:**
`````bash
curl http://localhost:3000/aliases
`````

**Response (200):**
`````json
[
  {
    "uuid": "660e8400-e29b-41d4-a716-446655440001",
    "name": "api*simples",
    "template*name": "rest*api*simples",
    "created*at": "2026-07-12T15:00:00Z",
    "updated*at": "2026-07-12T15:00:00Z"
  }
]
`````

---

#### ``GET /aliases/:name``

Busca um alias pelo nome.

**Exemplo:**
`````bash
curl http://localhost:3000/aliases/api*simples
`````

**Response (200):**
`````json
{
  "uuid": "660e8400-e29b-41d4-a716-446655440001",
  "name": "api*simples",
  "template*name": "rest*api*simples",
  "created*at": "2026-07-12T15:00:00Z",
  "updated*at": "2026-07-12T15:00:00Z"
}
`````

**Response (404):** Alias não encontrado

---

#### ``POST /aliases`` ⚠️ Protegido

Cria um novo alias.

**Headers:**
`````
X-API-Key: sua-chave-admin
Content-Type: application/json
`````

**Corpo do Request:**
`````json
{
  "name": "api*simples",
  "template*name": "rest*api*simples"
}
`````

**Exemplo:**
`````bash
curl -X POST http://localhost:3000/aliases \
  -H "X-API-Key: sua-chave-admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "api*simples",
    "template*name": "rest*api*simples"
  }'
`````

**Response (201):**
`````json
{
  "uuid": "660e8400-e29b-41d4-a716-446655440001",
  "name": "api*simples",
  "template*name": "rest*api*simples",
  "created*at": "2026-07-12T15:00:00Z",
  "updated*at": "2026-07-12T15:00:00Z"
}
`````

---

#### ``PUT /aliases/:uuid`` ⚠️ Protegido

Atualiza um alias existente.

**Headers:**
`````
X-API-Key: sua-chave-admin
Content-Type: application/json
`````

**Corpo do Request:**
`````json
{
  "name": "api*simples",
  "template*name": "rest*api*simples"
}
`````

**Exemplo:**
`````bash
curl -X PUT http://localhost:3000/aliases/660e8400-e29b-41d4-a716-446655440001 \
  -H "X-API-Key: sua-chave-admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "api*simples",
    "template*name": "rest*api*simples"
  }'
`````

**Response (200):**
`````json
{
  "uuid": "660e8400-e29b-41d4-a716-446655440001",
  "name": "api*simples",
  "template*name": "rest*api*simples",
  "created*at": "2026-07-12T15:00:00Z",
  "updated_at": "2026-07-12T15:03:22Z"
}
`````

---

#### ``DELETE /aliases/:uuid`` ⚠️ Protegido

Deleta um alias.

**Headers:**
`````
X-API-Key: sua-chave-admin
`````

**Exemplo:**
`````bash
curl -X DELETE http://localhost:3000/aliases/660e8400-e29b-41d4-a716-446655440001 \
  -H "X-API-Key: sua-chave-admin"
`````

**Response (204):** Sem conteúdo (sucesso)

---

Estrutura do Projeto
--------------------

`````
goman/
├── cmd/
│   ├── api/
│   │   └── main.go              # Ponto de entrada da aplicação
│   └── routes/
│       ├── routes.go            # Definições das rotas da API
│       ├── auth.go              # Middleware de autenticação
│       └── health.go            # Tipos do health check
├── internal/
│   ├── core/
│   │   ├── types.go             # Tipos do domínio (Template, Alias)
│   │   └── database.go          # Interface do banco de dados
│   └── infra/
│       └── postgres.go          # Implementação PostgreSQL
├── .env.example                 # Template de variáveis de ambiente
├── docker-compose.yml           # Serviços Docker (API + PostgreSQL)
├── Dockerfile                   # Imagem de produção
├── Dockerfile.dev               # Imagem de desenvolvimento
├── Makefile                     # Comandos de build e execução
├── go.mod                       # Definição do módulo Go
└── README.md                    # Este arquivo
`````

Componentes Principais
~~~~~~~~~~~~~~~~~~~~~~

* **``cmd/api/main.go``**: Inicializa a app Fiber, carrega variáveis de ambiente, conecta ao PostgreSQL
* **``cmd/routes/routes.go``**: Define todos os endpoints da API e seus handlers
* **``cmd/routes/auth.go``**: Middleware ``Protected()`` valida o header ``X-API-Key``
* **``internal/core/types.go``**: Modelos do domínio (Template, Alias)
* **``internal/core/database.go``**: Interface do banco de dados definindo operações
* **``internal/infra/postgres.go``**: Implementação PostgreSQL com connection pool

---

Desenvolvimento
---------------

Pré-requisitos
~~~~~~~~~~~~~~

* Go 1.26.4+
* PostgreSQL 12+

Comandos Make Disponíveis
~~~~~~~~~~~~~~~~~~~~~~~~~

`````bash
make build              # Build do binário para bin/goman
make run                # Executa servidor de desenvolvimento
make watch              # Executa com hot reload (requer Air)
make air-install        # Instala ferramenta Air para hot reload
make test               # Executa todos os testes
make fmt                # Formata código com go fmt
make tidy               # Arruma dependências Go
make clean              # Remove artefatos de build
`````

Executando Localmente
~~~~~~~~~~~~~~~~~~~~~

#. **Configurar PostgreSQL:**
   `````bash
   make postgres-up
   `````

#. **Iniciar servidor de desenvolvimento com hot reload:**
   `````bash
   make air-install  # Somente primeira vez
   make watch
   `````

#. **Ou executar diretamente:**
   `````bash
   make run
   `````

#. **Testar endpoints:**
   `````bash
   curl http://localhost:3000/health
   `````

#. **Parar PostgreSQL:**
   `````bash
   make postgres-down
   `````

---

Docker
------

Pré-requisitos
~~~~~~~~~~~~~~

* Docker & Docker Compose instalados, ou
* Podman & Podman Compose instalados

Executando com Docker
~~~~~~~~~~~~~~~~~~~~~

#. **Configurar ambiente:**
   `````bash
   cp .env.example .env
   `````

#. **Build e inicie todos os serviços:**
   `````bash
   make docker-up
   `````

   Isso inicia:
   - PostgreSQL (porta 5432)
   - Goman API (porta 3000)

#. **Verificar:**
   `````bash
   curl http://localhost:3000/health
   `````

#. **Ver logs:**
   `````bash
   make docker-logs
   `````

#. **Verificar containers executando:**
   `````bash
   make docker-ps
   `````

#. **Parar serviços:**
   `````bash
   make docker-down
   `````

Usando Podman
~~~~~~~~~~~~~

Por padrão, o Makefile usa ``docker compose``. Para usar Podman:

`````bash
COMPOSE=podman compose make docker-up
`````

Ou exporte a variável:
`````bash
export COMPOSE="podman compose"
make docker-up
`````

Desenvolvimento com Docker (Hot Reload)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Modo watch com hot reload em Docker:

`````bash
make docker-watch-up
make docker-watch-logs
make docker-watch-down
`````

---

Notas de Segurança
------------------

* **Chave API**: O header ``X-API-Key`` protege operações POST, PUT e DELETE
* Sempre use uma chave forte e aleatória em produção
* Nunca commit ``.env`` no controle de versão (está em ``.gitignore``)
* Use ``https://`` em produção e adicione rate limiting para deployments públicos

---

Licença
-------

Este projeto faz parte da iniciativa CommonBerry e é distribuido sobre a licensa `GPLv3 <LICENSE>`_

---

Contribuindo
------------

Contribuições são bem-vindas! Por favor garanta:

#. Código formatado: ``make fmt``
#. Sem problemas de linting: ``go vet ./...``

---

Suporte
-------

Para problemas, dúvidas ou sugestões, abra uma issue no GitHub.
