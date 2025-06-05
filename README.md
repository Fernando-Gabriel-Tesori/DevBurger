📘 README.md
markdown
Copiar
Editar
# 🍔 DevBurger API

DevBurger é uma API REST escrita em Go com o objetivo de simular o backend de uma hamburgueria moderna. Este projeto serve como estudo de boas práticas em desenvolvimento web com Go, utilizando arquitetura MVC, autenticação JWT, Docker, PostgreSQL e MongoDB.

---

## 🚀 Tecnologias

- **Go 1.24**
- **PostgreSQL** – armazenamento relacional (usuários, produtos, pedidos)
- **MongoDB** – armazenamento não relacional (ex: logs, análises)
- **JWT** – autenticação
- **Docker & Docker Compose** – containerização
- **HTTPie** – para testes de API

---

## 📁 Estrutura de Pastas

.
├── cmd/api/ # Entry point da aplicação
├── controllers/ # Lógica dos handlers HTTP
├── models/ # Modelos de dados
├── middlewares/ # Autenticação, validações, etc.
├── routes/ # Definição de rotas públicas e protegidas
├── database/ # Conexões com PostgreSQL e MongoDB
├── utils/ # Funções auxiliares (hash, JWT, etc.)
├── Dockerfile # Build final para produção
├── Dockerfile.debug # Ambiente de debug
├── docker-compose.yml # Orquestração dos serviços
├── .env # Variáveis de ambiente (não versionar)
├── .env.example # Exemplo de configuração
└── go.mod / go.sum # Dependências do Go

yaml
Copiar
Editar

---

## ⚙️ Configuração

### Pré-requisitos

- Docker
- Docker Compose

### Instalação

1. Clone o repositório:

```bash
git clone 
cd backend
Copie o arquivo de variáveis de ambiente:

bash
Copiar
Editar
cp .env.example .env
Suba os containers com Docker Compose:

bash
Copiar
Editar
docker-compose up --build
A API estará disponível em http://localhost:3000.

🔐 Autenticação
A autenticação é baseada em JWT.

POST /signup – Cria um novo usuário

POST /login – Retorna um token JWT

Rotas protegidas requerem o header:

http
Copiar
Editar
Authorization: Bearer <seu-token-jwt>
📬 Endpoints Principais
Rota	Método	Protegida	Descrição
/signup	POST	❌	Registro de usuário
/login	POST	❌	Login e geração de token JWT
/users	GET	✅	Listar usuários
/products	CRUD	✅	Gerenciar produtos
/orders	CRUD	✅	Gerenciar pedidos

Para testes, recomenda-se usar HTTPie. Exemplo:

bash
Copiar
Editar
http POST :3000/signup name=Joao email=joao@mail.com password=123456
🧪 Testes
Para testar as rotas protegidas, utilize o token JWT obtido no login:

bash
Copiar
Editar
export TOKEN="seu_token_aqui"
http GET :3000/users Authorization:"Bearer $TOKEN"
🐳 Debug com Docker
bash
Copiar
Editar
docker build -f Dockerfile.debug -t devburger-debug .
docker run -it devburger-debug
📌 Observações
O projeto utiliza arquitetura MVC para manter separação de responsabilidades.

O PostgreSQL é utilizado como banco relacional principal.

O MongoDB pode ser usado para armazenar logs, sessões ou histórico de pedidos.

O projeto foi idealizado como laboratório de estudo.

📖 Licença
Projeto educacional sem fins lucrativos. Uso livre para estudos e aprendizagem.

