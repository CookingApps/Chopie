### 🍔 Chopie

A scalable and production-ready **food delivery backend** built with Golang, designed to power modern food ordering platforms similar to **Chowdeck**, Uber Eats, and Glovo.

This project is currently **in active development**, focusing on clean architecture, performance, reliability, and real-world production practices.

![GitHub](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-000000?style=for-the-badge&logo=gin&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

---

### ✨ Features

**Currently Implemented / In Progress:**

- User registration and authentication (JWT + Refresh tokens)
- Restaurant management and profile
- Menu & Category management
- Shopping cart functionality
- Order placement and tracking
- Paystack payment gateway integration (initialize + webhook)
- Background job processing with RabbitMQ
- Redis caching for performance
- Rate limiting and input validation
- Full Docker & Docker Compose setup

**Planned / Upcoming:**

- Real-time order tracking
- Rider/delivery agent system
- Geospatial queries (nearby restaurants)
- Reviews and ratings
- Promotions & discounts
- Admin dashboard

---

### 🏗️ Architecture

This project follows **Clean Architecture** (Layered Architecture) to ensure maintainability, testability, and scalability.

```
Chopie/
├── cmd/app/                    # Application entry point
├── internal/
│   ├── config/                 # Configuration (Viper)
│   ├── model/                  # Database models (GORM)
│   ├── dto/                    # Request & Response DTOs
│   ├── repository/             # Database access layer
│   ├── service/                # Business logic layer
│   ├── handler/                # HTTP handlers (Gin)
│   ├── middleware/             # Auth, Rate limiting, Logging
│   ├── queue/                  # RabbitMQ producers & consumers
│   └── util/                   # Helpers (JWT, errors, password hashing)
├── migrations/                 # Database migrations
├── docker-compose.yml
├── Dockerfile
└── README.md
```

**Key Design Principles:**

- Separation of concerns (Handler → Service → Repository)
- Dependency Injection
- Repository pattern for easy testing and swapping databases
- Graceful error handling and logging with Zap

_I would add the architecture diagram in a bit_

---

### 🚀 How to Run Locally

#### Prerequisites

- Docker and Docker Compose installed
- Git

#### Steps

1. **Clone the repository**

   ```bash
   git clone https://github.com/CookingApps/Chopie.git
   cd Chopie
   ```

2. **Copy environment variables**

   ```bash
   cp .env.example .env
   ```

3. **Start all services (PostgreSQL, Redis, RabbitMQ + App)**

   ```bash
   docker-compose up -d --build
   ```

4. **Run database migrations** (once the app is running)

   ```bash
   # This will be automated via the app in future versions
   ```

5. **Test the API**
   Open your browser or Postman and visit:
   ```
   http://localhost:8080/health
   ```

**Available Services:**

- App → `http://localhost:8080`
- PostgreSQL → Port 5432
- Redis → Port 6379
- RabbitMQ Management → `http://localhost:15672` (guest/guest)

---

### 🛠️ Tech Stack

- **Language**: Go 1.23
- **Framework**: Gin
- **ORM**: GORM v2
- **Database**: PostgreSQL (Neon)
- **Cache**: Redis
- **Queue**: RabbitMQ
- **Payment**: Paystack
- **Auth**: JWT
- **Logging**: Zap
- **Containerization**: Docker + Docker Compose

---

### 📋 API Endpoints

**Base URL**: `http://localhost:8080/api/v1`

- **Auth**: `/auth/register`, `/auth/login`, `/auth/refresh`
- **Restaurants**: `/restaurants`, `/restaurants/:id`
- **Menu**: `/restaurants/:id/menu`
- **Cart**: `/cart`, `/cart/items`
- **Orders**: `/orders`, `/orders/:id`, `/orders/:id/status`
- **Payments**: `/payments/webhook`

**Note**: Full Postman documentation will be added soon.

---

### 📸 Screenshots / Demo

---

### 🔄 Current Status

**In Active Development** — Core features (Auth, Restaurant, Menu, Cart, Order, Payment) are being built step by step.

**Next Milestones:**

- Complete Order flow with Paystack integration
- Add background job processing
- Implement comprehensive tests
- Deploy to Render

---

### 🤝 Contributing

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](https://github.com/CookingApps/Chopie/issues).

---

**Cooked by CookingApps**

---
