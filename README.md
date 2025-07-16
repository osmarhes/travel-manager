# Travel Manager API

API de gerenciamento de pedidos de viagem corporativa, construída com **Go** e **Echo**, utilizando autenticação JWT, testes automatizados e Docker.

---

## 🚀 Funcionalidades

- Registro e login de usuários
- CRUD de pedidos de viagem
- Alteração de status (aprovado/cancelado)
- Filtros por status, datas e destino
- Middleware JWT
- Notificação simulada por log
- Testes com `testify`

---

## 🐳 Como rodar com Docker

```bash
docker-compose up --build
```

A API estará disponível em: [http://localhost:8080](http://localhost:8080)

---

## ⚙️ Variáveis de ambiente (.env)

```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=travel_manager
JWT_SECRET=super-secret-key
PORT=8080
```

---

## 🧪 Rodando os testes

Localmente (fora do Docker):

```bash
go test ./... -v
```
go run main.go
---

## 📂 Estrutura de diretórios

```
cmd/server         # Inicialização do servidor
internal/auth      # Registro/Login de usuários
internal/user      # Modelo de usuário
internal/travel    # Pedido de viagem
internal/authmiddleware # Middleware JWT
pkg/database       # Conexão e migrations
```

---

## 📬 Endpoints principais

| Método | Rota               | Protegido | Descrição                  |
|--------|--------------------|-----------|----------------------------|
| POST   | /auth/register     | ❌        | Cria um novo usuário       |
| POST   | /auth/login        | ❌        | Autentica e retorna token  |
| GET    | /travels           | ✅        | Lista pedidos do usuário   |
| POST   | /travels           | ✅        | Cria novo pedido           |
| GET    | /travels/:id       | ✅        | Consulta por ID            |
| PUT    | /travels/:id/status| ✅        | Atualiza status            |

---

## 🛠 Requisitos

- Go 1.23+
- Docker e Docker Compose
- PostgreSQL (via container)

---

## ✅ Testado e funcionando com

- Go 1.23.0
- Echo v4
- PostgreSQL 15
- Docker Compose v2+

---

## 📄 Licença

Projeto técnico demonstrativo.