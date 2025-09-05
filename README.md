# Go Analytics SDK

A complete analytics solution with a TypeScript SDK for frontend tracking and a Go backend for data processing and storage.

## 🚀 What This Repository Contains

### **Frontend SDK (`/sdk`)**

- TypeScript analytics SDK for tracking user events
- Supports pageviews, clicks, and conversions
- Easy integration with any web application
- Built with TypeScript and bundled with tsdown

### **Backend API (`/api`)**

- Go-based analytics API server
- Redis for event queuing and processing
- PostgreSQL for data storage
- Consumer groups for reliable event processing
- CORS-enabled for cross-origin requests

### **Demo Application (`/web/next-analytics-demo`)**

- Next.js demo application showcasing the SDK
- Product catalog with shopping cart functionality
- Real-time analytics tracking
- Complete e-commerce analytics example

## 📋 Prerequisites

- **Go 1.21+** (for the API server)
- **Node.js 18+** (for the SDK and demo)
- **Docker & Docker Compose** (for databases)
- **Redis** (for event queuing)
- **PostgreSQL** (for data storage)

## 🛠️ Installation & Setup

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd go-analytics-sdk
```

### 2. Start the Infrastructure

```bash
# Start Redis and PostgreSQL using Docker Compose
docker-compose up -d
```

### 3. Set Up the Backend API

```bash
cd api

# Install Go dependencies
go mod tidy

# Run database migrations (if needed)
# The API will create tables automatically on first run

# Start the API server
go run main.go
```

The API will be available at `http://localhost:8080`

### 4. Set Up the Demo Application

```bash
cd web/next-analytics-demo

# Install dependencies
npm install

# Start the development server
npm run dev
```

The demo will be available at `http://localhost:3000`

### 5. Build the SDK (Optional)

```bash
cd sdk

# Install dependencies
npm install

# Build the SDK
npm run build
```

## 🎯 How to Use

### **1. Install the SDK in Your Project**

```bash
npm install <your-sdk-package>
```

### **2. Initialize Analytics**

```typescript
import { AnalyticsSDK } from 'go-analytics-sdk';

const analytics = new AnalyticsSDK({
  clientId: 'your-client-id',
  baseUrl: 'http://localhost:8080', // Your API endpoint
  debug: true, // Enable for development
});
```

### **3. Track Events**

```typescript
// Track pageviews
analytics.sendEvent({
  user_id: 'user123',
  event_type: 'pageview',
  event_url: window.location.href,
  eventData: {
    page_title: document.title,
    referrer: document.referrer,
  },
});

// Track clicks
analytics.sendEvent({
  user_id: 'user123',
  event_type: 'click',
  event_url: window.location.href,
  eventData: {
    element: 'button',
    product_id: 'prod_123',
  },
});
```

### **4. Track Conversions**

```typescript
analytics.sendEvent({
  user_id: 'user123',
  event_type: 'conversion',
  event_url: window.location.href,
  eventData: {
    order_id: 'order_456',
    order_total: 99.99,
    products: [
      {
        product_id: 'prod_123',
        product_name: 'Widget',
        price: 99.99,
        quantity: 1,
      },
    ],
  },
});
```

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   Go API        │    │   Databases     │
│   (Next.js)     │───▶│   Server        │───▶│   Redis + PG    │
│                 │    │                 │    │                 │
│ - SDK           │    │ - Event Handlers│    │ - Event Queue   │
│ - Analytics     │    │ - Redis Streams │    │ - Data Storage  │
│ - Demo App      │    │ - CORS Support  │    │ - Consumer Groups│
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 📊 Event Types

### **Pageview Events**

- Track page visits and navigation
- Automatic URL and referrer capture
- Custom page metadata support

### **Click Events**

- Track user interactions
- Element identification
- Product and action tracking

### **Conversion Events**

- E-commerce transactions
- Order tracking
- Revenue analytics

## 🔧 Configuration

### **SDK Configuration**

```typescript
interface AnalyticsConfig {
  clientId: string; // Required: Your client identifier
  baseUrl?: string; // Optional: API endpoint (default: localhost:8080)
  debug?: boolean; // Optional: Enable debug logging
}
```

### **API Configuration**

- **Port**: Set via `PORT` environment variable (default: 8080)
- **Database**: PostgreSQL connection via environment variables
- **Redis**: Redis connection for event queuing

## 🚀 Production Deployment

### **Backend API**

```bash
# Build the Go binary
cd api
go build -o analytics-api main.go

# Run with environment variables
PORT=8080 ./analytics-api
```

### **Frontend SDK**

```bash
# Build and publish to npm
cd sdk
npm run build
npm publish
```

## 📈 Monitoring & Analytics

- **Real-time event processing** via Redis streams
- **Consumer groups** for reliable event handling
- **PostgreSQL storage** for analytics queries
- **CORS support** for cross-origin tracking

## 🧪 Testing

### **Run the Demo**

1. Start the infrastructure: `docker-compose up -d`
2. Start the API: `cd api && go run main.go`
3. Start the demo: `cd web/next-analytics-demo && npm run dev`
4. Visit `http://localhost:3000` and interact with the products
5. Check the API logs to see events being processed

## 📝 API Endpoints

- `POST /event/pageview` - Track pageview events
- `POST /event/click` - Track click events
- `POST /event/conversion` - Track conversion events
- `GET /health` - Health check endpoint

## 📄 License

MIT License - see LICENSE file for details

**Built with ❤️ using Go, TypeScript, Next.js, Redis, and PostgreSQL**
