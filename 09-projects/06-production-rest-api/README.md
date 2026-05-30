# Production-Grade REST API with Kafka & MongoDB

## 📋 Project Overview

A real-world, production-grade REST API built with Go featuring:
- **REST API**: Chi router with comprehensive endpoints
- **Database**: MongoDB with CRUD operations
- **Event Streaming**: Kafka producer & consumer
- **Authentication**: JWT token-based auth
- **Error Handling**: Comprehensive error responses
- **Logging**: Structured logging with context
- **Testing**: Unit and integration tests
- **Docker**: Containerized deployment
- **CI/CD**: GitHub Actions pipeline

**Difficulty**: ⭐⭐⭐ Advanced  
**Estimated Duration**: 20-25 hours  
**Prerequisites**: Modules 01-08 completed

---

## 🎯 Learning Objectives

By completing this project, you will learn:

- [ ] Building scalable REST APIs with Chi
- [ ] MongoDB integration and CRUD operations
- [ ] Apache Kafka producer/consumer patterns
- [ ] JWT authentication and middleware
- [ ] Error handling and validation
- [ ] Structured logging and monitoring
- [ ] Unit and integration testing
- [ ] Docker and docker-compose
- [ ] Environment configuration management
- [ ] Graceful shutdown and resource cleanup

---

## 📁 Project Structure

```
06-production-rest-api/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── models/
│   │   ├── user.go                # User model
│   │   ├── product.go             # Product model
│   │   └── event.go               # Kafka event model
│   ├── database/
│   │   ├── mongodb.go             # MongoDB connection & setup
│   │   ├── user_repository.go     # User CRUD operations
│   │   └── product_repository.go  # Product CRUD operations
│   ├── handlers/
│   │   ├── user_handler.go        # User HTTP handlers
│   │   ├── product_handler.go     # Product HTTP handlers
│   │   └── health_handler.go      # Health check endpoint
│   ├── kafka/
│   │   ├── producer.go            # Kafka producer
│   │   └── consumer.go            # Kafka consumer
│   ├── middleware/
│   │   ├── auth.go                # JWT authentication
│   │   ├── logging.go             # Request logging
│   │   └── error.go               # Error handling
│   └── router.go                   # Route definitions
├── migrations/
│   └── indexes.js                 # MongoDB indexes
├── scripts/
│   ├── docker-compose.yml         # Docker services
│   └── seed.go                    # Database seeding
├── tests/
│   ├── handlers_test.go           # Handler tests
│   ├── repository_test.go         # Repository tests
│   └── integration_test.go        # Integration tests
├── .env.example                    # Environment template
├── go.mod                          # Go modules
├── go.sum                          # Module checksums
├── Dockerfile                      # API container
├── docker-compose.yml              # All services
└── README.md                       # This file
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- MongoDB 6.0+
- Apache Kafka 3.5+
- Docker & Docker Compose
- Git

### Quick Start with Docker

```bash
# Clone and navigate
cd 09-projects/06-production-rest-api

# Start all services
docker-compose up -d

# Wait for services to be ready (30 seconds)
sleep 30

# Run the server
make run

# API is available at http://localhost:8080
```

### Local Development Setup

```bash
# 1. Install dependencies
go mod download
go mod tidy

# 2. Set up environment
cp .env.example .env
# Edit .env with your local settings

# 3. Start dependencies (MongoDB, Kafka)
docker-compose up -d mongo kafka

# 4. Run migrations
make migrate

# 5. Start the server
make run

# 6. Server runs on http://localhost:8080
```

---

## 📊 Visual Architecture

### Request Lifecycle Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant MW as Middleware Stack
    participant R as Chi Router
    participant H as Handler
    participant S as Service/Repo
    participant DB as MongoDB
    participant K as Kafka

    C->>MW: HTTP Request
    MW->>MW: Logger Middleware
    MW->>MW: Auth Middleware (JWT)
    MW->>MW: CORS Middleware
    MW->>MW: RequestID Middleware
    MW->>R: Route Matched
    
    R->>H: Call Handler
    H->>H: Validate Request Body
    
    alt Create/Update Operation
        H->>S: userRepo.Create()
        S->>DB: InsertOne()
        DB-->>S: Inserted Result
        S-->>H: User Object
        H->>K: Publish Event
        K-->>H: Ack
    else Read Operation
        H->>S: userRepo.GetByID()
        S->>DB: FindOne()
        DB-->>S: User Document
        S-->>H: User Object
    else Delete Operation
        H->>S: userRepo.Delete()
        S->>DB: DeleteOne()
        DB-->>S: Deleted Count
        S-->>H: Success
        H->>K: Publish Delete Event
    end
    
    H-->>C: JSON Response
```

### Kafka Event Streaming Architecture

```mermaid
flowchart TB
    subgraph Services[Application Services]
        API[REST API Handler]
        REPO[Repository Layer]
    end
    
    subgraph Kafka[Apache Kafka]
        PROD[Producer]
        BROKER[Kafka Broker]
        CONS[Consumer]
        TOPIC_USER[Topic: user-events]
        TOPIC_PRODUCT[Topic: product-events]
    end
    
    subgraph Storage[Data Storage]
        MONGO[(MongoDB)]
    end
    
    subgraph External[External Systems]
        CACHE[Cache Layer]
        LOG[Log Aggregator]
        EMAIL[Email Service]
        ANALYTICS[Analytics Platform]
    end
    
    API -->|HTTP Request| REPO
    REPO -->|CRUD Operations| MONGO
    MONGO -->|Response| REPO
    REPO -->|Event| PROD
    
    PROD -->|Publish| BROKER
    BROKER --> TOPIC_USER
    BROKER --> TOPIC_PRODUCT
    BROKER -->|Deliver| CONS
    
    CONS -->|Process Event| CACHE
    CONS -->|Process Event| LOG
    CONS -->|Process Event| EMAIL
    CONS -->|Process Event| ANALYTICS
    
    style API fill:#4CAF50,color:white
    style REPO fill:#2196F3,color:white
    style MONGO fill:#47A248,color:white
    style BROKER fill:#231F20,color:white
    style PROD fill:#FF6F00,color:white
    style CONS fill:#FF6F00,color:white
```

### MongoDB CRUD Sequence

```mermaid
sequenceDiagram
    participant H as HTTP Handler
    participant R as UserRepository
    participant DB as MongoDB
    participant K as Kafka Producer

    Note over H,R: CREATE User
    H->>R: userRepo.Create(ctx, user)
    R->>DB: collection.InsertOne(ctx, user)
    DB-->>R: InsertOneResult
    R-->>H: &User{ID: insertID}
    H->>K: Publish "user.created" event

    Note over H,R: READ User
    H->>R: userRepo.GetByID(ctx, id)
    R->>R: objectID := primitive.ObjectIDFromHex(id)
    R->>DB: collection.FindOne(ctx, bson.M{"_id": objectID})
    DB-->>R: SingleResult
    R-->>H: &User{...}

    Note over H,R: UPDATE User
    H->>R: userRepo.Update(ctx, id, update)
    R->>DB: collection.UpdateOne(ctx, filter, update)
    DB-->>R: UpdateResult
    R-->>H: Success/Error

    Note over H,R: DELETE User
    H->>R: userRepo.Delete(ctx, id)
    R->>DB: collection.DeleteOne(ctx, filter)
    DB-->>R: DeleteResult
    R-->>H: Success/Error
```

### Middleware Chain

```mermaid
flowchart LR
    REQ[HTTP Request] --> L[Logger]
    L --> RI[RequestID]
    RI --> AUTH[JWT Auth]
    AUTH --> CORS[CORS]
    CORS --> RL[Rate Limiter]
    RL --> TIMEOUT[Timeout]
    TIMEOUT --> HNDLR[Handler]
    HNDLR --> RESP[JSON Response]
    
    subgraph Error[Error Handling]
        TIMEOUT -.-> TIMEOUT_ERR[504 Timeout]
        AUTH -.-> AUTH_ERR[401 Unauthorized]
        RL -.-> RL_ERR[429 Too Many Requests]
    end
    
    style REQ fill:#4CAF50,color:white
    style RESP fill:#4CAF50,color:white
    style HNDLR fill:#2196F3,color:white
    style TIMEOUT_ERR fill:#f44336,color:white
    style AUTH_ERR fill:#f44336,color:white
    style RL_ERR fill:#f44336,color:white
```

### Producer-Consumer Event Flow

```mermaid
flowchart TB
    subgraph Producer[Producer - Publish Event]
        EVENT[Create Event Object] --> SERIALIZE[Serialize to JSON]
        SERIALIZE --> WRITE[kafka.WriteMessages]
        WRITE --> TOPIC{Topic Selection}
        TOPIC -->|user-events| U_TOPIC[(user-events<br/>3 partitions)]
        TOPIC -->|product-events| P_TOPIC[(product-events<br/>3 partitions)]
    end
    
    subgraph Consumer[Consumer - Process Event]
        U_TOPIC --> FETCH{kafka.FetchMessage}
        P_TOPIC --> FETCH
        FETCH --> DESERIALIZE[Unmarshal JSON]
        DESERIALIZE --> ROUTE{Event Type}
        ROUTE -->|user.created| UC[Handle Create]
        ROUTE -->|user.updated| UU[Handle Update]
        ROUTE -->|user.deleted| UD[Handle Delete]
        ROUTE -->|product.*| PH[Handle Product Event]
        UC --> COMMIT[Commit Offset]
        UU --> COMMIT
        UD --> COMMIT
        PH --> COMMIT
    end
    
    WRITE -->|Error| RETRY{Retry}
    RETRY -->|3 attempts| WRITE
    RETRY -->|Failed| DEAD[Dead Letter Queue]
    
    style EVENT fill:#FF6F00,color:white
    style WRITE fill:#FF6F00,color:white
    style U_TOPIC fill:#231F20,color:white
    style P_TOPIC fill:#231F20,color:white
    style FETCH fill:#FF6F00,color:white
    style COMMIT fill:#4CAF50,color:white
    style DEAD fill:#f44336,color:white
```

### Graceful Shutdown Flow

```mermaid
sequenceDiagram
    participant S as Server
    participant OS as OS Signal
    participant HC as HTTP Server
    participant DB as MongoDB
    participant K as Kafka
    
    Note over S: Running normally
    OS->>S: SIGINT/SIGTERM
    S->>S: Create shutdown context (30s timeout)
    
    S->>HC: server.Shutdown(ctx)
    HC->>HC: Stop accepting new connections
    HC->>HC: Drain active connections
    
    S->>K: consumer.Close()
    K-->>S: Consumer stopped
    
    S->>K: producer.Close()
    K-->>S: Producer flushed & closed
    
    S->>DB: mongo.Close(ctx)
    DB-->>S: Disconnected
    
    HC-->>S: Server shut down gracefully
    S-->>OS: All resources cleaned up
```

---

## 📚 Key Concepts

### 1. REST API Design

```go
// User endpoints
GET    /api/v1/users              # List all users
GET    /api/v1/users/{id}         # Get user by ID
POST   /api/v1/users              # Create new user
PUT    /api/v1/users/{id}         # Update user
DELETE /api/v1/users/{id}         # Delete user

// Product endpoints
GET    /api/v1/products           # List all products
GET    /api/v1/products/{id}      # Get product by ID
POST   /api/v1/products           # Create new product
PUT    /api/v1/products/{id}      # Update product
DELETE /api/v1/products/{id}      # Delete product

// Health check
GET    /health                    # Health status
```

### 2. MongoDB CRUD Operations

```go
// Create
func (r *UserRepository) Create(ctx context.Context, user *User) (*User, error)

// Read
func (r *UserRepository) GetByID(ctx context.Context, id string) (*User, error)
func (r *UserRepository) List(ctx context.Context, filter interface{}) ([]*User, error)

// Update
func (r *UserRepository) Update(ctx context.Context, id string, update *User) error

// Delete
func (r *UserRepository) Delete(ctx context.Context, id string) error
```

### 3. Kafka Integration

**Producer Pattern** - Publish events when data changes:
```go
// Publish user created event
event := &Event{
    Type: "user.created",
    Data: user,
}
producer.Publish(ctx, "user-events", event)
```

**Consumer Pattern** - React to events:
```go
// Listen for product updates
consumer.Subscribe("product-events", func(event *Event) error {
    // Handle product event
    return nil
})
```

### 4. Authentication Middleware

```go
// JWT Token in request header
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...

// Middleware validates token
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Validate token
        // Add user to context
        next.ServeHTTP(w, r)
    })
}
```

### 5. Error Handling

```go
// Structured error responses
{
    "error": {
        "code": "VALIDATION_ERROR",
        "message": "Invalid input",
        "details": {
            "field": "email",
            "issue": "email already exists"
        }
    }
}
```

---

## 🔗 Essential Links & Resources

### Go & REST APIs
- [Chi Router Documentation](https://github.com/go-chi/chi)
- [Building Web Services with Go](https://golang.org/doc/tutorial/web-service-gin)
- [RESTful API Best Practices](https://restfulapi.net/)

### MongoDB
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [MongoDB Documentation](https://docs.mongodb.com/)
- [CRUD Operations Guide](https://docs.mongodb.com/drivers/go/current/crud/)

### Kafka
- [Apache Kafka Documentation](https://kafka.apache.org/documentation/)
- [Confluent Kafka Go Client](https://github.com/confluentinc/confluent-kafka-go)
- [Kafka Design Patterns](https://kafka.apache.org/intro)

### Authentication
- [JWT (JSON Web Tokens)](https://jwt.io/)
- [Go JWT Libraries](https://github.com/golang-jwt/jwt)
- [OAuth 2.0 Basics](https://www.oauth.com/)

### Docker & DevOps
- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [12 Factor App](https://12factor.net/)

### Testing
- [Go Testing Package](https://golang.org/pkg/testing/)
- [Table-Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Testify Framework](https://github.com/stretchr/testify)

---

## 💻 Implementation Details

### 1. Main Application Entry Point

See `cmd/server/main.go`:
- Initialize logger
- Load configuration
- Connect to MongoDB
- Setup Kafka client
- Start HTTP server
- Handle graceful shutdown

### 2. MongoDB Repository Pattern

See `internal/database/user_repository.go`:
- Connection pooling
- CRUD operations
- Query optimization
- Error handling
- Bulk operations

### 3. Kafka Producer

See `internal/kafka/producer.go`:
- Publish events to topics
- Handle publication errors
- Async confirmation
- Message serialization

### 4. Kafka Consumer

See `internal/kafka/consumer.go`:
- Subscribe to topics
- Process messages
- Handle errors gracefully
- Offset management

### 5. HTTP Handlers

See `internal/handlers/user_handler.go`:
- Request validation
- Error responses
- Context usage
- Response formatting

### 6. Middleware Stack

See `internal/middleware/`:
- Authentication
- Logging
- Error handling
- CORS
- Rate limiting

---

## 🧪 Testing

### Run All Tests
```bash
make test
```

### Unit Tests
```bash
# Test handlers
go test ./internal/handlers -v

# Test repositories
go test ./internal/database -v

# Test Kafka logic
go test ./internal/kafka -v
```

### Integration Tests
```bash
# Full integration test (requires MongoDB & Kafka)
go test ./tests -v -tags=integration
```

### Coverage Report
```bash
make coverage
```

---

## 🐳 Docker Commands

### Build & Run
```bash
# Build Docker image
docker build -t go-api:latest .

# Run with Docker Compose
docker-compose up

# View logs
docker-compose logs -f api

# Stop services
docker-compose down
```

### Service Management
```bash
# Check service status
docker-compose ps

# View MongoDB logs
docker-compose logs mongo

# MongoDB shell
docker-compose exec mongo mongosh

# Kafka topics
docker-compose exec kafka kafka-topics --bootstrap-server localhost:9092 --list
```

---

## 📊 API Examples

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepass123"
  }'
```

### Get All Users
```bash
curl -X GET http://localhost:8080/api/v1/users
```

### Update Product
```bash
curl -X PUT http://localhost:8080/api/v1/products/123 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Product",
    "price": 99.99
  }'
```

### Health Check
```bash
curl -X GET http://localhost:8080/health
```

---

## 🔐 Environment Configuration

See `.env.example`:

```env
# Server
PORT=8080
ENV=production

# MongoDB
MONGO_URI=mongodb://mongo:27017
MONGO_DB=go_api_db

# Kafka
KAFKA_BROKERS=kafka:9092
KAFKA_USER_TOPIC=user-events
KAFKA_PRODUCT_TOPIC=product-events

# JWT
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRY=24h

# Logging
LOG_LEVEL=info
```

---

## 🚀 Deployment

### Using Docker Compose (Development/Staging)
```bash
docker-compose -f docker-compose.yml up -d
```

### Using Kubernetes (Production)
See `k8s/` directory for Kubernetes manifests

### Cloud Deployment
- AWS ECS: ECR + ECS with RDS MongoDB
- GCP: Cloud Run + Cloud MongoDB
- Azure: Container Instances + Azure Cosmos DB

---

## 📈 Performance Considerations

### Optimization Tips

1. **Database Indexing**
   - Index frequently queried fields
   - Use compound indexes for common queries
   - Monitor index usage

2. **Caching**
   - Implement Redis for hot data
   - Cache API responses
   - Use ETags for conditional requests

3. **Connection Pooling**
   - Configure MongoDB connection pool
   - Kafka connection pool
   - HTTP client pooling

4. **Pagination**
   - Implement cursor-based pagination
   - Limit result sets
   - Use projection to reduce data transfer

5. **Concurrency**
   - Use goroutines for I/O operations
   - Buffered channels for message queues
   - Context timeouts for requests

---

## 🐛 Troubleshooting

### MongoDB Connection Issues
```bash
# Check MongoDB logs
docker-compose logs mongo

# Verify connection
docker-compose exec mongo mongosh --eval "db.adminCommand('ping')"
```

### Kafka Issues
```bash
# Check Kafka logs
docker-compose logs kafka

# List topics
docker-compose exec kafka kafka-topics --bootstrap-server localhost:9092 --list

# Check consumer groups
docker-compose exec kafka kafka-consumer-groups --bootstrap-server localhost:9092 --list
```

### API Issues
```bash
# Check API logs
docker-compose logs api

# Test API health
curl http://localhost:8080/health

# Check environment variables
docker-compose exec api env | grep GO_API
```

---

## 📝 Common Tasks

### Add New API Endpoint

1. Create model in `internal/models/`
2. Create handler in `internal/handlers/`
3. Create repository in `internal/database/`
4. Add route in `internal/router.go`
5. Write tests

### Add Kafka Topic

1. Update `docker-compose.yml`
2. Add topic to consumer in `internal/kafka/consumer.go`
3. Update event types in `internal/models/event.go`

### Create Database Migration

```bash
# Create migration
go run migrations/main.go create_user_index

# Run migration
make migrate
```

---

## 🎯 Learning Checklist

- [ ] Understand REST API design principles
- [ ] Implement CRUD operations with MongoDB
- [ ] Publish events to Kafka
- [ ] Consume and process Kafka events
- [ ] Implement JWT authentication
- [ ] Write comprehensive tests
- [ ] Handle errors gracefully
- [ ] Use structured logging
- [ ] Deploy with Docker
- [ ] Monitor and optimize performance

---

## 🔗 Additional Resources

### Code Quality
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Best Practices](https://golang.org/doc/code)

### System Design
- [System Design Interview](https://www.youtube.com/watch?v=ZgdS0EUiiHY)
- [Microservices Patterns](https://microservices.io/patterns/)
- [Design Patterns in Go](https://refactoring.guru/design-patterns/go)

### DevOps
- [The Twelve-Factor App](https://12factor.net/)
- [Heroku Deployment](https://devcenter.heroku.com/articles/getting-started-with-go)
- [AWS Lambda Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)

---

## 📞 Support

For issues or questions:
1. Check troubleshooting section
2. Review code comments
3. Check CI/CD logs
4. Open GitHub issue with details

---

**Project Status**: ✅ Production Ready  
**Last Updated**: May 30, 2026  
**Go Version**: 1.21+  
**License**: MIT

---

Happy Building! 🚀
