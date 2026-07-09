# 📝 TODO - Goman API (Fase 1)

## [ ] Setup Inicial & Infraestrutura (`internal/infra`)
- [ ] Configurar o roteador HTTP (Chi, Fiber ou net/http puro) no `internal/infra`.
- [ ] Criar um handler básico de `/healthz` ou `/ping` para testar o servidor.
- [ ] Instalar e configurar o parser de YAML (`gopkg.in/yaml.v3`).

## [ ] Core & Modelagem (`internal/core`)
- [ ] Definir a struct `Template` baseada no formato do arquivo YAML.
- [ ] Definir a struct `Alias` para o mapeamento dos pacotes (ex: `lipgloss` -> URL completa).
- [ ] Criar as interfaces (Ports) para o repositório/armazenamento que a infra vai implementar depois.

## [ ] Endpoints da API (`cmd/api`)
- [ ] Inicializar o servidor HTTP no `main.go`.
- [ ] Criar a rota mockada `GET /v1/aliases` para simular o retorno de um pacote curto.
- [ ] Criar a rota mockada `GET /v1/templates` para simular a listagem dos templates em YAML.

## [ ] Automação & Teste
- [ ] Configurar o `Makefile` com comandos básicos: `make run`, `make build` e `make test`.
- [ ] Rodar o servidor localmente e fazer um `curl` para garantir que as rotas respondem em JSON.
