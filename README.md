# StockVisualization

StockVisualization is a fullstack application designed to provide stock market insights. It consists of two main components:

1. **StockBack**: A backend service built with Go (Gin) for managing stock data, APIs, and business logic.
2. **StockVI**: A frontend application built with Vue 3 and Vite for visualizing stock data.

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Development](#development)
- [Testing](#testing)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Backend**:
  - RESTful APIs for stock data and ratings.
  - Modular, layered architecture (domain, services, ports, handlers, repositories).
  - Database integration with PostgreSQL using GORM.
  - Unit tests for domain, ports, and services (with mocks).
  - Docker support for containerized deployment.
  - Environment variable configuration for database and external APIs.

- **Frontend**:
  - Interactive UI for stock visualization.
  - Responsive design for desktop and mobile.
  - Component-based architecture with Vue 3 and Pinia.
  - Unit testing with Vitest.

---

## Tech Stack

### Backend (StockBack)
- **Language**: Go (1.24+)
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Testing**: Go test, Testify
- **Containerization**: Docker

### Frontend (StockVI)
- **Framework**: Vue 3
- **Build Tool**: Vite
- **Testing**: Vitest
- **Styling**: TailwindCSS

---

## Project Structure

### Backend (`StockBack`)

```
StockBack/
├── cmd/                # Application entry point (main.go/app.go)
├── internal/
│   ├── core/
│   │   ├── domain/     # Domain models
│   │   ├── services/   # Business logic/services
│   │   ├── ports/      # Interfaces/ports for repositories/services
│   ├── handlers/       # HTTP handlers and DTOs
│   ├── repositories/   # Database repositories
├── tests/              # Unit tests (mirroring internal/core structure, using mocks)
├── go.mod
├── go.sum
├── Dockerfile
```

### Frontend (`StockVI`)
```
StockVI/
├── public/             # Static assets
├── src/
│   ├── components/     # Vue components
│   ├── stores/         # State management (Pinia)
│   ├── assets/         # Images, styles, etc.
├── vite.config.ts
├── package.json
```

---

## Setup Instructions

### Backend Setup

1. **Install Go**: Ensure you have Go 1.24+ installed.
2. **Clone the repository**:
   ```sh
   git clone https://github.com/your-repo/StockVisualization.git
   cd StockVisualization/StockBack
   ```
3. **Install dependencies**:
   ```sh
   go mod tidy
   ```
4. **Set up the database**:
   - Create a PostgreSQL database.
   - Set the `DATABASE_URL` environment variable (and any others needed, e.g., `DATA_SOURCE`, `AUTH_TOKEN`).
5. **Run the backend (development):**
   ```sh
   go run cmd/app.go
   ```
6. **Build the backend (production):**
   ```sh
   go build -o stock-back ./cmd/app.go
   ./stock-back
   ```
7. **(Optional) Run with Docker:**
   ```sh
   docker build -t stock-back .
   docker run -e DATABASE_URL=your_db_url -p 8080:8080 stock-back
   ```

### Frontend Setup

1. **Install Node.js**: Ensure you have Node.js 16+ installed.
2. **Navigate to the frontend directory**:
   ```sh
   cd ../StockVI
   ```
3. **Install dependencies**:
   ```sh
   npm install
   ```
4. **Run the development server**:
   ```sh
   npm run dev
   ```

---

## Development

### Backend
- The backend entry point is located at `StockBack/cmd/app.go`.
- Core domain models are in `StockBack/internal/core/domain`.
- Services and business logic are in `StockBack/internal/core/services`.
- Ports/interfaces are in `StockBack/internal/core/ports`.
- Handlers (HTTP) are in `StockBack/internal/handlers`.
- Unit tests (with mocks) are in `StockBack/tests/core/`.

### Frontend
- The main Vue app is in `StockVI/src/App.vue`.
- Components are located in `StockVI/src/components`.

---

## Testing

### Backend
- Unit tests are provided for domain, ports, and services.
- Tests use mocks for dependencies.
- Run all tests with coverage:
  ```sh
  go test -coverprofile=coverage.out ./...
  ```
- To view coverage in VS Code:
  - Run the above command.
  - Use the command palette: `Go: Toggle Coverage in Workspace`.

### Frontend
- Unit tests are written using Vitest.
- Run tests using:
  ```sh
  npm run test:unit
  ```

---

<!-- ## Deployment

### Backend

- **Build the binary:**
  ```sh
  go build -o stock-back ./cmd/app.go
  ```
- **Or use Docker:**
  ```sh
  docker build -t stock-back .
  docker run -e DATABASE_URL=your_db_url -p 8080:8080 stock-back
  ```
- **(Optional) Use Docker Compose for DB + App:**
  ```yaml
  version: '3.8'
  services:
    db:
      image: postgres:15
      environment:
        POSTGRES_DB: stockdb
        POSTGRES_USER: user
        POSTGRES_PASSWORD: pass
      ports:
        - "5432:5432"
    app:
      build: .
      environment:
        DATABASE_URL: postgres://user:pass@db:5432/stockdb?sslmode=disable
      ports:
        - "8080:8080"
      depends_on:
        - db
  ```

### Frontend

- **Build for production:**
  ```sh
  npm run build
  ```
- **Deploy the `dist/` folder to your static hosting provider.**

--- -->

## Troubleshooting

- **Backend not responding:**  
  - Check that the server is running and listening on the expected port (default Gin port is `8080`).
  - Ensure all required environment variables are set.
  - Check logs for errors or panics.
- **Database connection issues:**  
  - Verify your `DATABASE_URL` and database status.

