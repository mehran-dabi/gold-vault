# GoldVault

GoldVault is a backend service for a gold trading platform, allowing users to buy and sell assets like gold, manage their wallets, and view pricing information. The system is built using microservices with a focus on scalability, reliability, and performance.

## Features

- User management with OTP-based authentication.
- Asset trading functionality with real-time price updates.
- Wallet service allowing users to manage multiple asset types.
- Separate microservices for handling users, assets, pricing, and trades.
- Configurable system flag for allowing purchases even when inventory is low.
- High scalability with gRPC, Redis, and PostgreSQL.
- Optimized caching with Redis for handling high traffic on asset price requests.

## Architecture

**GoldVault** is built using a microservices architecture with the following key components:

- User-Service: Manages user profiles, OTP authentication, and roles.
- Wallet-Service: Manages user wallets, asset balances, and transaction histories.
- Asset-Service: Provides up-to-date asset prices and handles price caching.
- Trade-Service: Manages the trading system, including inventory management and transaction logging.

The project follows a hexagonal architecture to separate business logic from infrastructure concerns. It also supports dependency injection using Uber's fx library.

## Technologies

- **Golang**: The primary programming language for backend services.
- **PostgreSQL**: Used for persistence.
- **Redis**: Used for caching asset prices and managing rate limits.
- **gRPC**: Used for communication between microservices.
- **Docker**: Containerizes the services for easy deployment.
- **Uber Fx**: Dependency injection framework for Go.
- **Swagger**: API documentation for the services.

## Folder Structure

```
.
├── cmd
│   ├── app
│   └── clients
├── internal
│   ├── config
│   ├── core
│   │   ├── application
│   │   │   └── services
│   │   └── domain
│   │       └── entities
│   └── infrastructure
│   │       ├── db
│   │       ├── cache
│   │       └── persistence   
│   ├── interfaces
│   │   ├── api
│   │   └── dto
│   └── server
├── migrations
└── proto
```

## Contributing
If you'd like to contribute to the project, feel free to open a pull request or raise an issue.

## License
This project is licensed under the MIT License.




