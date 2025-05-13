
# StockVisualization

StockVisualization is a fullstack application designed to provide stock market insights. It consists of two main components:

1. **StockBack**: A backend service built with Go for managing stock data and APIs.
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
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Backend**:
  - RESTful APIs for stock data.
  - Database integration with PostgreSQL using GORM.
  - Modular architecture with services, handlers, and repositories.

- **Frontend**:
  - Interactive UI for stock visualization.
  - Responsive design for desktop and mobile.
  - Component-based architecture with Vue 3.

---

## Tech Stack

### Backend (StockBack)
- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM

### Frontend (StockVI)
- **Framework**: Vue 3
- **Build Tool**: Vite
- **Testing**: Vitest
- **Styling**: Scoped CSS

---

## Project Structure

### Backend (`StockBack`)

```
StockBack/
├── cmd/                # Application entry point
├── internal/
│   ├── core/           # Core domain models and services
│   ├── handlers/       # HTTP handlers
│   ├── repositories/   # Database repositories
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksums
```

### Frontend (`StockVI`)
```
StockVI/
├── public/             # Static assets
├── src/
│   ├── components/     # Vue components
│   ├── stores/         # State management
│   ├── assets/         # Images, styles, etc.
├── vite.config.ts      # Vite configuration
├── package.json        # NPM dependencies
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
   - Update the database connection string in the environment variables.
5. **Run the backend**:
   ```sh
   go run cmd/app.go
   ```

### Frontend Setup

1. **Install Node.js**: Ensure you have Node.js 16+ installed.
2. **Navigate to the frontend directory**:
   ```sh
   cd StockVisualization/StockVI
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

### Frontend
- The main Vue app is in `StockVI/src/App.vue`.
- Components are located in `StockVI/src/components`.

---

## Testing

### Backend
- Add unit tests for services and handlers.
- Run tests using:
  ```sh
  go test ./...
  ```

### Frontend
- Unit tests are written using Vitest.
- Run tests using:
  ```sh
  npm run test:unit
  ```

