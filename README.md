# Subly - A Subscription Management Web Application with Email Integration

Subly is a Go-based web application that provides subscription plan management with integrated email notifications and user authentication. It offers a secure and scalable solution for managing user subscriptions with features like email verification, session management, and invoice generation.

The application is built using modern Go practices and leverages PostgreSQL for data persistence, Redis for session management, and MailHog for email handling in development. It implements a clean architecture pattern with separation of concerns between handlers, models, and business logic.

## Repository Structure

```
.
├── cmd/web/                    # Main application code
│   ├── config.go              # Application configuration and server setup
│   ├── db.go                  # Database and Redis connection management
│   ├── handlers.go            # HTTP request handlers
│   ├── mailer.go              # Email service implementation
│   ├── main.go                # Application entry point
│   ├── middleware.go          # HTTP middleware functions
│   ├── render.go              # Template rendering logic
│   ├── routes.go              # URL routing definitions
│   └── signer.go              # URL signing utilities
├── data/                      # Data models and business logic
│   ├── models.go              # Core model definitions
│   ├── plan.go                # Subscription plan model
│   └── user.go                # User model and authentication
└── docker-compose.yml         # Docker services configuration
```

## Usage Instructions

### Prerequisites

- Go
- Docker and Docker Compose
- Makefile

Rename the `.env.example` to `.env` and fill in the required environment variables.

### Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd subly
```

2. Start the infrastructure services:

```bash
docker-compose up -d
```

3. Install Go dependencies:

```bash
go mod download
```

4. Run the application:

```bash
make start
```

5. Access the application at `http://localhost/`.
6. Access MailHog at `http://localhost:8025/` to view sent emails.
