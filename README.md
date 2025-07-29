# Natsa's Place - Backend API

## Descrição

Backend robusto para uma plataforma de aluguel de casas e hospedagem, voltada para uma única empresa (modelo proprietário). A plataforma permite que apenas a empresa gerencie os imóveis, enquanto os usuários podem visualizar, reservar e avaliar as propriedades.

## Funcionalidades Principais

### ✅ Sistema de Autenticação

- [x] Registro de usuários
- [x] Login e logout
- [x] Verificação de email
- [x] Recuperação de senha
- [x] Refresh token
- [x] Middleware de autenticação JWT
- [x] Controle de acesso baseado em roles (USER/ADMIN)

### 🏠 Gestão de Propriedades (Em Desenvolvimento)

- [ ] Listagem de casas disponíveis
- [ ] Filtros avançados (preço, localização, datas, tipo)
- [ ] Detalhes da propriedade
- [ ] Sistema de amenidades
- [ ] Gestão de imagens
- [ ] Controle de disponibilidade

### 📅 Sistema de Reservas (Em Desenvolvimento)

- [ ] Seleção de datas
- [ ] Verificação de disponibilidade
- [ ] Cálculo de preços
- [ ] Confirmação de reservas
- [ ] Gestão de check-in/check-out

### ⭐ Sistema de Avaliações (Em Desenvolvimento)

- [ ] Avaliações pós-estadia
- [ ] Comentários públicos e privados
- [ ] Sistema de notas por categoria
- [ ] Moderação de avaliações

### 📧 Sistema de Notificações (Em Desenvolvimento)

- [ ] Notificações por email
- [ ] Notificações por SMS
- [ ] Lembretes automáticos
- [ ] Preferências de notificação

### 💳 Sistema de Pagamentos (Em Desenvolvimento)

- [ ] Integração com Stripe
- [ ] Processamento de pagamentos
- [ ] Gestão de reembolsos
- [ ] Histórico de transações

## Arquitetura

O projeto segue uma arquitetura limpa e organizada:

```
back-end-natsas-place/
├── config/           # Configurações da aplicação
├── controllers/      # Controladores HTTP
├── database/         # Configuração e migrações do banco
├── middleware/       # Middlewares personalizados
├── models/          # Modelos de dados
├── repository/      # Camada de acesso aos dados
├── routes/          # Definição das rotas
├── services/        # Lógica de negócio
└── utils/           # Utilitários e helpers
```

## Tecnologias Utilizadas

- **Go 1.24**: Linguagem principal
- **Fiber v2**: Framework web
- **GORM**: ORM para Go
- **PostgreSQL**: Banco de dados
- **JWT**: Autenticação stateless
- **bcrypt**: Hash de senhas
- **Validator**: Validação de dados
- **Mailgun**: Envio de emails
- **Stripe**: Processamento de pagamentos

## Instalação

1. Clone o repositório:

```bash
git clone <repository-url>
cd back-end-natsas-place
```

2. Instale as dependências:

```bash
go mod download
```

3. Configure as variáveis de ambiente:

```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

4. Execute as migrações:

```bash
go run main.go
```

## Variáveis de Ambiente

```env
# Database Configuration
DATABASE_URL=postgresql://username:password@localhost:5432/natsas_place?sslmode=disable

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here

# Server Configuration
PORT=3001
ENVIRONMENT=development
FRONTEND_URL=http://localhost:3000
BACKEND_URL=http://localhost:3001

# Email Configuration
EMAIL_SENDER=noreply@natsasplace.com
EMAIL_PASSWORD=your-email-password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587

# Mailgun Configuration
MAILGUN_DOMAIN=your-domain.com
MAILGUN_API_KEY=your-mailgun-api-key

# Stripe Configuration
STRIPE_SECRET_KEY=sk_test_your-stripe-secret-key
STRIPE_PUBLIC_KEY=pk_test_your-stripe-public-key
```

## Estrutura da API

### Endpoints de Autenticação

```
POST /api/v1/auth/register       # Registrar usuário
POST /api/v1/auth/login          # Login
POST /api/v1/auth/refresh        # Renovar token
POST /api/v1/auth/logout         # Logout
POST /api/v1/auth/forgot-password # Recuperar senha
POST /api/v1/auth/reset-password  # Redefinir senha
GET  /api/v1/auth/verify-email   # Verificar email
POST /api/v1/auth/resend-verification # Reenviar verificação
```

### Endpoints de Perfil (Protegidos)

```
GET  /api/v1/profile            # Obter perfil
PUT  /api/v1/profile            # Atualizar perfil
POST /api/v1/profile/change-password # Alterar senha
GET  /api/v1/me                 # Informações básicas do usuário
```

### Endpoints de Administração (Admin Only)

```
GET    /api/v1/admin/users      # Listar usuários
GET    /api/v1/admin/users/:id  # Obter usuário
PUT    /api/v1/admin/users/:id  # Atualizar usuário
DELETE /api/v1/admin/users/:id  # Deletar usuário
```

## Modelos de Dados

### User

- Informações básicas do usuário
- Controle de acesso (roles)
- Verificação de email
- Histórico de login

### Property

- Detalhes da propriedade
- Preços e políticas
- Amenidades
- Status de disponibilidade

### Booking

- Reservas dos usuários
- Datas de check-in/out
- Cálculo de preços
- Status da reserva

### Review

- Avaliações por categoria
- Comentários públicos/privados
- Sistema de moderação

### Payment

- Transações
- Histórico de pagamentos
- Reembolsos

## Middleware

### AuthMiddleware

- Valida tokens JWT
- Extrai informações do usuário
- Protege rotas privadas

### AdminMiddleware

- Verifica privilégios de administrador
- Usado em rotas administrativas

### RateLimitMiddleware

- Limita requisições por IP
- Protege contra ataques

## Segurança

- **JWT**: Autenticação stateless
- **bcrypt**: Hash seguro de senhas
- **CORS**: Controle de origem
- **Rate Limiting**: Proteção contra spam
- **Validação**: Entrada de dados
- **Middleware**: Controle de acesso

## Desenvolvimento

### Executar em modo desenvolvimento:

```bash
go run main.go
```

### Executar testes:

```bash
go test ./...
```

### Compilar para produção:

```bash
go build -o tmp/main .
```

### Estrutura de logs:

```
[2024-01-01 10:00:00] 200 - GET /api/v1/health - 1.234ms
```

## Roadmap

### Fase 1 - Autenticação ✅

- [x] Sistema de autenticação completo
- [x] Middleware de segurança
- [x] Controle de acesso

### Fase 2 - Propriedades (Em Desenvolvimento)

- [ ] CRUD de propriedades
- [ ] Sistema de amenidades
- [ ] Upload de imagens
- [ ] Filtros e busca

### Fase 3 - Reservas (Em Desenvolvimento)

- [ ] Sistema de booking
- [ ] Controle de disponibilidade
- [ ] Cálculo de preços
- [ ] Gestão de reservas

### Fase 4 - Avaliações (Em Desenvolvimento)

- [ ] Sistema de reviews
- [ ] Moderação
- [ ] Estatísticas

### Fase 5 - Notificações (Em Desenvolvimento)

- [ ] Email/SMS
- [ ] Lembretes automáticos
- [ ] Preferências

### Fase 6 - Pagamentos (Em Desenvolvimento)

- [ ] Integração Stripe
- [ ] Processamento
- [ ] Reembolsos

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## Contato

Para dúvidas ou sugestões, entre em contato através do email: dev@natsasplace.com
