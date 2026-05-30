# Production REST API - Implementation Guide

## 📚 Learning Path Through This Project

### Part 1: Architecture & Design (2-3 hours)
- [ ] Review REST API design principles
- [ ] Understand MongoDB CRUD patterns
- [ ] Learn Kafka producer/consumer architecture
- [ ] Study middleware patterns

**Resources**:
- [REST API Best Practices](https://restfulapi.net/)
- [MongoDB Design Patterns](https://docs.mongodb.com/guides/server/introduction/)
- [Kafka Architecture](https://kafka.apache.org/intro)

### Part 2: MongoDB Integration (4-5 hours)
- [ ] Setup MongoDB connection
- [ ] Create repository pattern
- [ ] Implement CRUD operations
- [ ] Handle MongoDB errors
- [ ] Create indexes

**Key Files**:
- `internal/database/mongodb.go` - Connection setup
- `internal/database/user_repository.go` - CRUD operations

**Tasks**:
```go
// 1. Create a user
user := &models.User{...}
createdUser, err := userRepo.Create(ctx, user)

// 2. Read a user
user, err := userRepo.GetByID(ctx, "userID")

// 3. Update a user
err := userRepo.Update(ctx, "userID", updatedUser)

// 4. Delete a user
err := userRepo.Delete(ctx, "userID")
```

### Part 3: Kafka Integration (4-5 hours)
- [ ] Setup Kafka producer
- [ ] Setup Kafka consumer
- [ ] Publish events
- [ ] Consume and handle events
- [ ] Error handling

**Key Files**:
- `internal/kafka/producer.go` - Publishing events
- `internal/kafka/consumer.go` - Consuming events

**Event Flow**:
```
User Action → Handler → Repository → Publish Event to Kafka
                                   → Consumer processes event
                                   → Update cache/external systems
```

### Part 4: HTTP API & Handlers (3-4 hours)
- [ ] Setup Chi router
- [ ] Create request/response models
- [ ] Implement handlers
- [ ] Add error handling
- [ ] Request validation

**Key Files**:
- `internal/handlers/user_handler.go` - User endpoints
- `cmd/server/main.go` - Server setup

### Part 5: Testing & Deployment (3-4 hours)
- [ ] Write unit tests
- [ ] Write integration tests
- [ ] Test with real MongoDB/Kafka
- [ ] Create Docker setup
- [ ] Test deployment

**Key Files**:
- `tests/models_test.go` - Model tests
- `Dockerfile` - Container image
- `docker-compose.yml` - Local development

---

## 🔑 Key Concepts Explained

### 1. Repository Pattern

The repository pattern abstracts data access:

```go
type UserRepository interface {
    Create(ctx context.Context, user *User) (*User, error)
    GetByID(ctx context.Context, id string) (*User, error)
    Update(ctx context.Context, id string, user *User) error
    Delete(ctx context.Context, id string) error
}
```

**Benefits**:
- Decouples business logic from data layer
- Easy to mock for testing
- Supports multiple data sources
- Single responsibility principle

### 2. Event-Driven Architecture

Events decouple services:

```
Service A writes data → Publishes event to Kafka
                     → Service B consumes event
                     → Service B updates its own data
```

**Benefits**:
- Loose coupling
- Scalability
- Asynchronous processing
- Audit trail

### 3. Middleware Chain

Middleware processes requests sequentially:

```
Request → Logger → Auth → Business Logic → Response
```

### 4. MongoDB Connection Pooling

Efficient connection management:

```go
// Connection pool handles multiple concurrent requests
client.WithOptions(opts)
// Reuses connections instead of creating new ones
```

---

## 🚀 Running the Project

### Quick Start (5 minutes)
```bash
cd 09-projects/06-production-rest-api

# Start all services
docker-compose up -d

# Run API
go run ./cmd/server/main.go
```

### Manual Setup (15 minutes)
```bash
# 1. Start MongoDB
docker run -d -p 27017:27017 mongo:6

# 2. Start Kafka stack
docker-compose up -d zookeeper kafka

# 3. Run application
make build
make run
```

---

## 📝 API Endpoints

### User Management

#### Create User
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepass123",
  "phone": "+1-555-0123",
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip_code": "10001",
    "country": "USA"
  }
}
```

**Response**:
```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "email": "john@example.com",
  "status": "active",
  "role": "user",
  "created_at": "2026-05-30T12:00:00Z",
  "updated_at": "2026-05-30T12:00:00Z"
}
```

#### Get User
```http
GET /api/v1/users?id=507f1f77bcf86cd799439011
```

#### List Users
```http
GET /api/v1/users
```

#### Update User
```http
PUT /api/v1/users?id=507f1f77bcf86cd799439011
Content-Type: application/json

{
  "name": "Jane Doe",
  "status": "inactive"
}
```

#### Delete User
```http
DELETE /api/v1/users?id=507f1f77bcf86cd799439011
```

---

## 🧪 Testing Examples

### Unit Test Example
```go
func TestCreateUser(t *testing.T) {
    user := &models.User{
        Name: "Test User",
        Email: "test@example.com",
    }
    
    if user.Name == "" {
        t.Error("Expected name, got empty")
    }
}
```

### Integration Test Pattern
```go
func TestUserHandlerIntegration(t *testing.T) {
    // 1. Setup MongoDB test database
    // 2. Create handler with test repo
    // 3. Make HTTP request
    // 4. Verify response
    // 5. Check database state
}
```

---

## 🔍 Debugging Tips

### View MongoDB Data
```bash
# Connect to MongoDB
docker-compose exec mongo mongosh

# List databases
show dbs

# Use database
use go_api_db

# View collections
show collections

# Find all users
db.users.find()

# Find specific user
db.users.findOne({email: "john@example.com"})
```

### Check Kafka Topics
```bash
# List topics
docker-compose exec kafka kafka-topics --bootstrap-server localhost:9092 --list

# Describe topic
docker-compose exec kafka kafka-topics --bootstrap-server localhost:9092 \
  --describe --topic user-events

# View messages
docker-compose exec kafka kafka-console-consumer --bootstrap-server localhost:9092 \
  --topic user-events --from-beginning
```

### View API Logs
```bash
# Docker logs
docker-compose logs -f api

# Or if running locally
go run ./cmd/server/main.go 2>&1 | tee app.log
```

---

## 📊 Performance Considerations

### 1. Database Indexing
```go
// Indexes speed up queries
db.users.createIndex({email: 1}, {unique: true})
db.products.createIndex({sku: 1}, {unique: true})
```

### 2. Connection Pooling
```go
// Configure for optimal performance
maxConnections := 100
idleConnections := 10
```

### 3. Kafka Partitions
```
// Partitions enable parallel processing
user-events: 3 partitions
product-events: 3 partitions
```

### 4. Pagination
```go
// Retrieve users in batches
users, err := userRepo.List(ctx, 0, 20)  // First 20
users, err := userRepo.List(ctx, 20, 20) // Next 20
```

---

## 🐛 Common Issues & Solutions

### MongoDB Connection Failed
```bash
# Check if MongoDB is running
docker-compose ps

# Check logs
docker-compose logs mongo

# Restart MongoDB
docker-compose restart mongo
```

### Kafka Connection Issues
```bash
# Verify Kafka is running
docker-compose logs kafka

# Check broker connectivity
docker-compose exec kafka kafka-broker-api-versions \
  --bootstrap-server localhost:9092
```

### API Port Already in Use
```bash
# Change port in .env
PORT=8081

# Or kill existing process
lsof -i :8080
kill -9 <PID>
```

---

## 📚 Further Learning

### Study These Patterns
1. **Repository Pattern** - Data abstraction
2. **Factory Pattern** - Object creation
3. **Singleton Pattern** - Single instance
4. **Observer Pattern** - Event handling

### Explore These Topics
1. **Transaction Management** - ACID compliance
2. **Caching Strategies** - Redis integration
3. **Rate Limiting** - API protection
4. **Logging & Monitoring** - Observability

### Production Considerations
1. **Security** - HTTPS, JWT, RBAC
2. **Scaling** - Horizontal scaling, load balancing
3. **Monitoring** - Metrics, alerts, logging
4. **Disaster Recovery** - Backup, failover

---

## 🎯 Next Steps

1. **Extend the API**
   - Add product endpoints (similar to user endpoints)
   - Add search functionality
   - Add filtering and sorting

2. **Improve Testing**
   - Add integration tests
   - Setup CI/CD pipeline
   - Increase coverage to 80%+

3. **Add Features**
   - Implement pagination properly
   - Add caching layer (Redis)
   - Add authentication (JWT)
   - Add rate limiting

4. **Production Hardening**
   - Add request validation
   - Implement error tracking (Sentry)
   - Add metrics collection (Prometheus)
   - Setup log aggregation (ELK)

---

## 📖 Resources

### Chi Router
- [Chi Documentation](https://github.com/go-chi/chi)
- [Routing Examples](https://github.com/go-chi/chi/tree/master/examples)

### MongoDB Go Driver
- [Driver Documentation](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [CRUD Examples](https://www.mongodb.com/docs/drivers/go/current/crud/)

### Kafka & segmentio/kafka-go
- [Kafka GO Driver](https://github.com/segmentio/kafka-go)
- [Examples](https://github.com/segmentio/kafka-go/tree/main/examples)

### Best Practices
- [12 Factor App](https://12factor.net/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [SOLID Principles](https://en.wikipedia.org/wiki/SOLID)

---

**Happy Building! 🚀**

Remember: Production-grade code requires attention to error handling, logging, testing, and operational concerns. This project demonstrates best practices in those areas.
