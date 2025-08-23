# FollowService

A GraphQL-based follow/unfollow service built with Go, demonstrating clean architecture principles and supporting both SQLite and PostgreSQL databases.

## 🏗️ Tech Stack

- **Language**: Go 1.18+
- **Web Framework**: Chi Router v5
- **GraphQL**: graph-gophers/graphql-go
- **Database**: 
  - SQLite (local development)
  - PostgreSQL 15 (production/Docker)
- **Architecture**: Domain-Driven Design (DDD)
- **Containerization**: Docker & Docker Compose

## 📁 Project Structure

```
FollowService/
├── cmd/server/          # Application entry point
├── internal/
│   ├── domain/          # Business logic layer
│   │   ├── entity/      # Domain entities
│   │   ├── repository/  # Repository interfaces
│   │   └── service/     # Business services
│   └── infrastructure/ # External concerns
│       ├── database/    # Database implementations
│       └── graphql/     # GraphQL resolvers & models
├── configs/            # Configuration files
├── migrations/         # Database migrations
├── docker/            # Docker compose files
└── docs/              # Documentation
```

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher
- Docker & Docker Compose (for PostgreSQL setup)
- Git

### 📥 Installation

1. **Clone the repository**
```bash
git clone <repository-url>
cd FollowService
```

2. **Install dependencies**
```bash
go mod tidy
```

## 🛠️ Local Development Setup

### Option 1: SQLite (Quick Start)

Perfect for development and testing:

```bash
# Run with SQLite (default)
go run ./cmd/server
```

The application will:
- Create a SQLite database at `data/followservice.db`
- Seed 5 initial users (user1-user5)
- Start server on http://localhost:8080

### Option 2: PostgreSQL with Docker

For production-like environment:

1. **Start PostgreSQL container**
```bash
docker-compose -f docker/docker-compose.yml up postgres -d
```

2. **Set environment variables**
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=followuser
export DB_PASSWORD=followpass
export DB_NAME=followservice
```

3. **Run the application**
```bash
go run ./cmd/server
```

### Option 3: Full Docker Setup

Run both application and database in containers:

```bash
docker-compose -f docker/docker-compose.yml up
```

## 🔧 Configuration

The application uses environment-based configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8080` | HTTP server port |
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `followuser` | Database username |
| `DB_PASSWORD` | `followpass` | Database password |
| `DB_NAME` | `followservice` | Database name |
| `USE_POSTGRES` | `false` | Force PostgreSQL usage |

## 🗄️ Database Access

### SQLite
```bash
# Access SQLite database
sqlite3 data/followservice.db

# Common commands
.tables                 # List tables
.schema users          # Show table structure
SELECT * FROM users;   # Query data
```

### PostgreSQL
```bash
# Access PostgreSQL via Docker
docker exec -it docker-postgres-1 psql -U followuser -d followservice

# Common commands
\dt                    # List tables
\d users              # Describe table
SELECT * FROM users;  # Query data
\q                    # Quit
```

## 🌐 API Usage

### GraphiQL Interface
Navigate to http://localhost:8080/graphiql for interactive GraphQL playground.

### Sample Queries & Mutations

### Sample Queries & Mutations

**Pre-seeded Users:** `user1` (Alice), `user2` (Bob), `user3` (Charlie), `user4` (Diana), `user5` (Eve)

#### Create a New User
```graphql
mutation {
  createUser(name: "Parag") {
    Id
    name
  }
}
```

#### Follow a User
```graphql
mutation {
  followUser(myId: "user1", targetId: "user5")
}
```
*Returns `true` if successful, `false` if already following or user doesn't exist*

#### Unfollow a User
```graphql
mutation {
  unfollowUser(myId: "user1", targetId: "user5")
}
```

#### Get User's Followers
```graphql
query {
  followers(Id: "user5") {
    Id
    name
  }
}
```

#### Get Users Someone is Following
```graphql
query {
  followings(Id: "user1") {
    Id
    name
  }
}
```

## 🧪 Testing

### Manual Testing
1. Start the server (SQLite or PostgreSQL)
2. Visit http://localhost:8080/graphiql
3. Run the sample mutations and queries above
4. Verify data persistence by querying the database directly

### Database Verification
```bash
# For SQLite
sqlite3 data/followservice.db "SELECT * FROM follows;"

# For PostgreSQL
docker exec -it docker-postgres-1 psql -U followuser -d followservice -c "SELECT * FROM follows;"
```

## 🏗️ Architecture Features

### ✅ Clean Architecture Benefits
- **Domain-Driven Design**: Clear separation of business logic
- **Repository Pattern**: Database abstraction for testing
- **Dependency Injection**: Proper service initialization
- **Environment Configuration**: Production-ready config management

### ✅ Production Ready Features
- **Database Persistence**: SQLite for development, PostgreSQL for production
- **Docker Support**: Containerized deployment
- **Health Checks**: PostgreSQL container health monitoring
- **Migrations**: Automatic database schema setup
- **Error Handling**: Comprehensive error responses
- **Input Validation**: Business rule enforcement

## 🚀 Deployment

### Docker Deployment
```bash
# Build and run with Docker Compose
docker-compose -f docker/docker-compose.yml up --build

# Scale the application
docker-compose -f docker/docker-compose.yml up --scale app=3
```

### Production Environment Variables
```bash
export DB_HOST=your-postgres-host
export DB_PORT=5432
export DB_USER=your-db-user
export DB_PASSWORD=your-secure-password
export DB_NAME=followservice
export SERVER_PORT=8080
```

## 📚 Additional Resources

- **GraphQL Schema**: Located at `configs/schema.graphql`
- **Database Migrations**: See `migrations/` directory
- **Docker Configuration**: Check `docker/docker-compose.yml`