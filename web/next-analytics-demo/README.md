# Next.js Analytics Demo

A comprehensive demo application showcasing the Go Analytics SDK integration with Next.js, featuring a product catalog, shopping cart functionality, and real-time event tracking.

## 🚀 Features

- **Product Catalog** - Interactive product cards with images and descriptions
- **Shopping Cart** - Add items, view cart, and checkout functionality
- **Event Tracking** - Pageviews, clicks, and conversions
- **State Management** - Zustand for cart state management
- **Responsive Design** - Mobile-first Tailwind CSS styling
- **Real-time Analytics** - Live event tracking to Go backend

## 📋 Prerequisites

- **Node.js** 18+
- **npm** or **yarn**
- **Go Analytics API** running on `localhost:8080`
- **Redis** and **PostgreSQL** (via Docker Compose)

## 🛠️ Installation & Setup

### 1. Navigate to Demo Directory

```bash
cd web/next-analytics-demo
```

### 2. Install Dependencies

```bash
npm install
# or
yarn install
```

### 3. Start the Development Server

```bash
npm run dev
# or
yarn dev
```

## 🎯 Demo Features

### **Analytics Tracking**

- **Pageview Events** - Automatic page load tracking
- **Click Events** - Button interactions and product clicks
- **Conversion Events** - Checkout completion tracking
- **Real-time Data** - Events sent to Go backend immediately

## 📊 Event Tracking Examples

### **Pageview Tracking**

```typescript
// Automatic on page load via AnalyticsProvider
trackEvent('user_007', 'pageview', {
  page_title: document.title,
  referrer: document.referrer,
});
```

### **Click Tracking**

```typescript
// Product card interactions
trackEvent('user_007', 'click', {
  element: 'add_to_cart_button',
  product_id: 'product-1',
  product_name: 'Wireless Headphones',
  product_price: '$99.99',
  product_category: ['headphones', 'electronics'],
});
```

### **Conversion Tracking**

```typescript
// Checkout completion
trackEvent('user_007', 'conversion', {
  order_id: 'order_123-2024-01-01T10:00:00Z',
  order_total: '379.97',
  products: [
    {
      product_id: 'product-1',
      product_name: 'Wireless Headphones',
      product_price: '$99.99',
      product_quantity: 1,
      product_category: ['headphones', 'electronics'],
    },
  ],
});
```

## 🧪 Testing the Demo

### **1. Start the Backend**

```bash
# From project root
docker-compose up -d
cd api
go run main.go
```

### **2. Start the Demo**

```bash
cd web/next-analytics-demo
npm run dev
```

### **3. Test Events**

1. **Visit** `http://localhost:3000`
2. **Check Console** - Pageview event should appear
3. **Click Products** - Add to cart and view details
4. **Check Cart** - Add multiple items
5. **Checkout** - Complete purchase
6. **Monitor Backend** - Events should appear in Go server logs

### **4. Verify Data**

- **Browser Console** - Debug logs from SDK
- **Go Server Logs** - Event processing confirmation
- **PostgreSQL** - Stored event data
- **Redis** - Event queue status

## 📄 License

MIT License - see LICENSE file for details

---

**Built with ❤️ using Next.js, React, TypeScript, and Tailwind CSS**
