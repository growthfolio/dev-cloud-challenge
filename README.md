# ☁️ Dev Cloud Challenge - API REST em Go

## 🎯 Objetivo de Aprendizado
API RESTful desenvolvida em Go como desafio de **Desenvolvimento Web e Cloud**. Implementa **CRUD completo** para gestão de dados estudantis com **PostgreSQL**, **Docker** e **CI/CD**, aplicando boas práticas de desenvolvimento e deploy em nuvem.

## 🛠️ Tecnologias Utilizadas
- **Linguagem:** Go
- **Banco de dados:** PostgreSQL
- **Containerização:** Docker, Docker Compose
- **Administração:** pgAdmin
- **CI/CD:** GitHub Actions
- **Deploy:** Heroku
- **Documentação:** Swagger UI
- **Qualidade:** SonarCloud

## 🚀 Demonstração
```bash
# Endpoints principais
GET    /students           # Listar todos os estudantes
POST   /students           # Criar novo estudante
PUT    /students/{id}      # Atualizar estudante
DELETE /students/{id}      # Remover estudante

# Swagger Documentation
https://dev-cloud-challenge-b3f5485f2dcf.herokuapp.com/swagger/index.html
```

## 📁 Estrutura do Projeto
```
dev-cloud-challenge/
├── cmd/
│   └── main.go                # Entry point da aplicação
├── internal/
│   ├── handlers/              # HTTP handlers
│   ├── models/                # Data models
│   ├── repository/            # Data access layer
│   └── services/              # Business logic
├── docs/                      # Swagger documentation
├── bin/                       # Compiled binaries
├── docker-compose.yml         # Orquestração de serviços
├── Dockerfile                 # Container configuration
└── .github/workflows/         # CI/CD pipelines
```

## 💡 Principais Aprendizados

### 🌐 API RESTful Design
- **HTTP methods:** GET, POST, PUT, DELETE apropriados
- **Status codes:** Códigos de resposta semânticos
- **JSON handling:** Serialização e deserialização
- **Error handling:** Tratamento consistente de erros
- **Middleware:** Logging, CORS, authentication

### 🐘 PostgreSQL Integration
- **Database migrations:** Versionamento de schema
- **Connection pooling:** Gerenciamento eficiente de conexões
- **Query optimization:** Consultas performáticas
- **Transaction management:** Consistência de dados
- **Environment configuration:** Configuração flexível

### 🐳 Containerização e Deploy
- **Docker multi-stage:** Builds otimizados
- **Docker Compose:** Orquestração local
- **Environment variables:** Configuração externa
- **Health checks:** Monitoramento de saúde
- **Cloud deployment:** Deploy automatizado

## 🧠 Conceitos Técnicos Estudados

### 1. **Clean API Architecture**
```go
// Handler layer
func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    result, err := h.service.CreateStudent(student)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
```

### 2. **Database Layer**
```go
// Repository pattern
type StudentRepository interface {
    Create(student *Student) error
    GetByID(id int) (*Student, error)
    GetAll() ([]*Student, error)
    Update(student *Student) error
    Delete(id int) error
}

type postgresStudentRepo struct {
    db *sql.DB
}
```

### 3. **Docker Configuration**
```dockerfile
# Multi-stage build
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

## 🚧 Desafios Enfrentados
1. **Database connectivity:** Configuração de conexão PostgreSQL
2. **Environment management:** Variáveis de ambiente em diferentes ambientes
3. **CORS handling:** Configuração para frontend integration
4. **Error consistency:** Padronização de respostas de erro
5. **Performance optimization:** Otimização de queries e conexões

## 📚 Recursos Utilizados
- [Go Web Programming](https://www.manning.com/books/go-web-programming)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker Best Practices](https://docs.docker.com/develop/best-practices/)
- [REST API Design Guidelines](https://restfulapi.net/)

## 📈 Próximos Passos
- [ ] Implementar autenticação JWT
- [ ] Adicionar rate limiting
- [ ] Implementar caching com Redis
- [ ] Adicionar métricas e monitoring
- [ ] Implementar testes de integração
- [ ] Adicionar validação de dados avançada

## 🔗 Projetos Relacionados
- [Go PriceGuard API](../go-priceguard-api/) - API Go com Clean Architecture
- [Go Antifraud MS](../go-antifraud-ms/) - Microserviço Go avançado
- [Spring Blog Platform](../spring-blog-platform/) - API similar em Java

---

**Desenvolvido por:** Felipe Macedo  
**Contato:** contato.dev.macedo@gmail.com  
**GitHub:** [FelipeMacedo](https://github.com/felipemacedo1)  
**LinkedIn:** [felipemacedo1](https://linkedin.com/in/felipemacedo1)

> 💡 **Reflexão:** Este projeto consolidou conhecimentos em desenvolvimento de APIs REST e deploy em nuvem. A experiência com Docker e CI/CD foi fundamental para compreender o ciclo completo de desenvolvimento moderno.