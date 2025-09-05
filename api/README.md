# Analytics API - Go Backend Service

A high-performance Go-based analytics API server that processes, queues, and stores user events from web applications.

## 🚀 Overview

This Go service provides a robust backend for the Analytics SDK, handling event ingestion, queuing with Redis streams, and persistent storage in PostgreSQL. Built with consumer groups for reliable event processing and CORS support for cross-origin requests.

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Handlers │    │   Redis Streams │    │   PostgreSQL    │
│                 │───▶│                 │───▶│                 │
│ - Pageview      │    │ - Event Queue   │    │ - Event Storage │
│ - Click         │    │ - Consumer Groups│   │ - Analytics DB  │
│ - Conversion    │    │ - Reliable Proc │    │ - Data Queries  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 📋 Prerequisites

- **Go 1.21+**
- **Redis 6.0+** (for event queuing)
- **PostgreSQL 13+** (for data storage)
- **Docker & Docker Compose** (for local development)

## 🛠️ Installation & Setup

### 1. Clone and Navigate

```bash
git clone git@github.com:aqmnguyen/go-analytics-sdk.git
cd go-analytics-sdk/api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Environment Setup

The service uses minimal environment configuration:

```bash
# Server Configuration (optional - defaults to 8080)
export PORT=8080
```

**Note**: Database and Redis connections are currently hardcoded for development:

### 4. Start Infrastructure

```bash
# From project root
docker-compose up -d
```

### 5. Run the Service

```bash
# Development mode with hot reload
air

# Or run directly
go run main.go
```

The API will be available at `http://localhost:8080`

## 📊 Event Types

### **Pageview Events**

```go
type PageviewEvent struct {
    ClientId  string `json:"client_id" required:"true"`
    UserId    string `json:"user_id" required:"true"`
    EventType string `json:"event_type" required:"true"`
    EventUrl  string `json:"event_url" required:"true"`
    Referrer  string `json:"referrer,omitempty"`
    IpAddress string `json:"ip_address,omitempty"`
    UserAgent string `json:"user_agent,omitempty"`
}
```

### **Click Events**

```go
type ClickEvent struct {
    ClientId  string `json:"client_id" required:"true"`
    UserId    string `json:"user_id" required:"true"`
    EventType string `json:"event_type" required:"true"`
    EventUrl  string `json:"event_url" required:"true"`
    Element   string `json:"element" required:"true"`
    Referrer  string `json:"referrer,omitempty"`
    IpAddress string `json:"ip_address,omitempty"`
    UserAgent string `json:"user_agent,omitempty"`
}
```

### **Conversion Events**

```go
type ConversionEvent struct {
    ClientId    string                   `json:"client_id" required:"true"`
    UserId      string                   `json:"user_id" required:"true"`
    EventType   string                   `json:"event_type" required:"true"`
    EventUrl    string                   `json:"event_url" required:"true"`
    ProductData []ConversionProductData  `json:"product_data,omitempty"`
}

type ConversionProductData struct {
    ProductId   string  `json:"product_id" required:"true"`
    ProductName string  `json:"product_name" required:"true"`
    Price       float64 `json:"price" required:"true"`
    Quantity    int     `json:"quantity" required:"true"`
    Category    string  `json:"category,omitempty"`
}
```

## 🔌 API Endpoints

### **Health Check**

```http
GET /health
```

Returns server status and basic information.

### **Pageview Tracking**

```http
POST /event/pageview
Content-Type: application/json

{
  "client_id": "client_123",
  "user_id": "user_456",
  "event_type": "pageview",
  "event_url": "https://example.com/page",
  "referrer": "https://google.com",
  "user_agent": "Mozilla/5.0...",
  "ip_address": "192.168.1.1"
}
```

### **Click Tracking**

```http
POST /event/click
Content-Type: application/json

{
  "client_id": "client_123",
  "user_id": "user_456",
  "event_type": "click",
  "event_url": "https://example.com/page",
  "element": "add_to_cart_button",
  "referrer": "https://example.com",
  "user_agent": "Mozilla/5.0...",
  "ip_address": "192.168.1.1"
}
```

### **Conversion Tracking**

```http
POST /event/conversion
Content-Type: application/json

{
  "client_id": "client_123",
  "user_id": "user_456",
  "event_type": "conversion",
  "event_url": "https://example.com/checkout",
  "product_data": [
    {
      "product_id": "prod_123",
      "product_name": "Widget",
      "price": 29.99,
      "quantity": 2,
      "category": "electronics"
    }
  ]
}
```

## 🔧 Configuration

### **Database Configuration**

The service automatically connects to PostgreSQL and creates necessary tables on startup.

### **Redis Configuration**

- Uses Redis streams for event queuing
- Implements consumer groups for reliable processing
- Automatic retry logic for failed events

## 🏃‍♂️ Development

### **Hot Reload with Air**

```bash
# Install air (if not already installed)
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## 🚀 Production Deployment

### **Build Binary**

```bash
# Build for current platform
go build -o analytics-api main.go

# Build for Linux (if deploying to Linux server)
GOOS=linux GOARCH=amd64 go build -o analytics-api main.go
```

### **Environment Variables**

```bash
# Only PORT is currently configurable
export PORT=8080

# Note: Database and Redis connections are hardcoded
# For production, you'll need to modify the source code to use environment variables
```

## 📈 Monitoring & Logging

### **Logging**

- Structured logging with timestamps
- Event processing status
- Error tracking and reporting
- Redis connection status

### **Health Monitoring**

- `/health` endpoint for uptime checks
- Database connection monitoring
- Redis connection monitoring
- Event processing metrics

## 🔍 Troubleshooting

### **Common Issues**

**Database Connection Failed**

```bash
# Check if PostgreSQL is running
docker-compose ps

# Check database logs
docker-compose logs postgres
```

**Redis Connection Failed**

```bash
# Check if Redis is running
docker-compose ps

# Check Redis logs
docker-compose logs redis
```

**Events Not Processing**

```bash
# Check Redis stream
redis-cli XINFO STREAM events:live

# Check consumer group
redis-cli XINFO GROUPS events:live
```

### **Debug Mode**

Enable debug logging by setting environment variable:

```bash
export DEBUG=true
```

## 🧪 Testing

### **Manual Testing**

```bash
# Test health endpoint
curl http://localhost:8080/health

# Test pageview event
curl -X POST http://localhost:8080/event/pageview \
  -H "Content-Type: application/json" \
  -d '{"client_id":"test","user_id":"user123","event_type":"pageview","event_url":"https://example.com"}'
```

### **Load Testing**

```bash
# Install hey (load testing tool)
go install github.com/rakyll/hey@latest

# Run load test
hey -n 1000 -c 10 -m POST -H "Content-Type: application/json" \
  -d '{"client_id":"test","user_id":"user123","event_type":"pageview","event_url":"https://example.com"}' \
  http://localhost:8080/event/pageview
```

## 📝 API Response Format

### **Success Response**

```json
{
  "client_id": "client_123",
  "user_id": "user_456",
  "event_type": "pageview",
  "event_url": "https://example.com/page",
  "referrer": "https://google.com",
  "user_agent": "Mozilla/5.0...",
  "ip_address": "192.168.1.1"
}
```

### **Error Response**

```json
{
  "error": "Missing required fields: client_id, user_id",
  "status": 400
}
```

## 📄 License

MIT License - see LICENSE file for details

---

**Built with ❤️ using Go, Redis, and PostgreSQL**
