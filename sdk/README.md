# Go Analytics SDK

A lightweight TypeScript SDK for tracking user analytics events on websites and web applications.

## 🚀 Features

- **TypeScript Support** - Full type safety and IntelliSense
- **Event Tracking** - Pageviews, clicks, and conversions
- **Graceful Error Handling** - Non-blocking event sending
- **Browser Compatibility** - Works in all modern browsers
- **Minimal Bundle Size** - Optimized for performance
- **Debug Mode** - Development-friendly logging

## 📦 Installation

```bash
npm install @just-baiting/go-analytics-sdk
```

## 🛠️ Quick Start

### 1. Initialize the SDK

```typescript
import { AnalyticsSDK } from '@just-baiting/go-analytics-sdk';

const analytics = new AnalyticsSDK({
  clientId: 'your-client-id',
  baseUrl: 'https://your-api-endpoint.com', // Optional, defaults to localhost:8080
  debug: true, // Optional, enables console logging
});
```

### 2. Track Events

```typescript
// Track a pageview
await analytics.sendEvent({
  user_id: 'user_123',
  event_type: 'pageview',
  event_url: window.location.href,
  event_data: {
    page_title: document.title,
    referrer: document.referrer,
  },
});

// Track a click
await analytics.sendEvent({
  user_id: 'user_123',
  event_type: 'click',
  event_url: window.location.href,
  event_data: {
    element: 'add_to_cart_button',
    product_id: 'prod_123',
    product_name: 'Widget',
  },
});

// Track a conversion
await analytics.sendEvent({
  user_id: 'user_123',
  event_type: 'conversion',
  event_url: window.location.href,
  event_data: {
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

## 📋 API Reference

### **AnalyticsSDK Class**

#### Constructor

```typescript
new AnalyticsSDK(config: AnalyticsConfig)
```

**Parameters:**

- `config.clientId` (string, required) - Your unique client identifier
- `config.baseUrl` (string, optional) - API endpoint URL (default: `http://localhost:8080`)
- `config.debug` (boolean, optional) - Enable debug logging (default: `false`)

#### Methods

##### `sendEvent(event: Event)`

Sends an analytics event to the server.

**Parameters:**

- `event.user_id` (string, required) - Unique user identifier
- `event.event_type` (string, required) - Event type: `'pageview'`, `'click'`, or `'conversion'`
- `event.event_url` (string, required) - Current page URL
- `event.event_data` (object, optional) - Additional event data

**Returns:** `Promise<any | null>` - Server response or `null` on error

##### `getConfig()`

Returns the current SDK configuration.

**Returns:** `AnalyticsConfig` - Current configuration object

## 🎯 Event Types

### **Pageview Events**

Track page visits and navigation:

```typescript
{
  user_id: 'user_123',
  event_type: 'pageview',
  event_url: 'https://example.com/page',
  event_data: {
    page_title: 'Product Page',
    referrer: 'https://google.com',
    user_agent: navigator.userAgent
  }
}
```

### **Click Events**

Track user interactions:

```typescript
{
  user_id: 'user_123',
  event_type: 'click',
  event_url: 'https://example.com/page',
  event_data: {
    element: 'add_to_cart_button',
    product_id: 'prod_123',
    product_name: 'Widget',
    product_price: 29.99
  }
}
```

### **Conversion Events**

Track e-commerce transactions:

```typescript
{
  user_id: 'user_123',
  event_type: 'conversion',
  event_url: 'https://example.com/checkout',
  event_data: {
    order_id: 'order_456',
    order_total: 99.99,
    products: [
      {
        product_id: 'prod_123',
        product_name: 'Widget',
        price: 29.99,
        quantity: 2,
        category: 'electronics'
      }
    ]
  }
}
```

## 🔧 Configuration

### **AnalyticsConfig Interface**

```typescript
interface AnalyticsConfig {
  clientId: string; // Required: Your client identifier
  clientKey?: string; // Optional: Additional authentication key
  baseUrl?: string; // Optional: API endpoint (default: localhost:8080)
  debug?: boolean; // Optional: Enable debug logging
}
```

## 🧪 Development

### **Build the SDK**

```bash
# Install dependencies
npm install

# Build for production
npm run build

# Build in watch mode
npm run dev

# Run tests
npm test

# Type checking
npm run typecheck
```

## 📄 License

MIT License - see LICENSE file for details

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

---

**Built with ❤️ using TypeScript**
