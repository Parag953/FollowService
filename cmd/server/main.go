package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"FollowService/configs"
	"FollowService/internal/domain/service"
	"FollowService/internal/infrastructure/database"
	"FollowService/internal/infrastructure/graphql"

	"github.com/go-chi/chi/v5"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load configuration
	config := configs.Load()

	// Initialize database
	db, err := initDatabase(config.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := database.NewUserRepository(db)
	followRepo := database.NewFollowRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, followRepo)

	// Initialize GraphQL resolver
	rootResolver := graphql.NewRootResolver(userService)

	// Load GraphQL schema
	schemaData, err := loadSchema()
	if err != nil {
		log.Fatalf("Failed to load schema: %v", err)
	}

	// Parse GraphQL schema
	schema := gql.MustParseSchema(string(schemaData), rootResolver)

	// Initialize HTTP server
	server := initHTTPServer(schema)

	log.Printf("Server is running on http://localhost:%s/graphiql", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, server))
}

func initDatabase(dbConfig configs.DatabaseConfig) (*sql.DB, error) {
	// Check if we should use PostgreSQL or SQLite
	// Use PostgreSQL if any of these conditions are met:
	// 1. DB_HOST is explicitly set (even if localhost)
	// 2. DB_USER is set (indicates intention to use PostgreSQL)
	// 3. DB_PASSWORD is set
	if dbConfig.User != "" || dbConfig.Password != "" || os.Getenv("USE_POSTGRES") == "true" {
		// Use PostgreSQL for production/Docker
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
		
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to open postgres connection: %w", err)
		}
		
		// Test the connection
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("failed to ping postgres: %w", err)
		}
		
		log.Printf("Connected to PostgreSQL at %s:%s", dbConfig.Host, dbConfig.Port)
		return db, nil
	}
	
	// Use SQLite for local development
	os.MkdirAll("data", 0755)
	db, err := sql.Open("sqlite3", "data/followservice.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite connection: %w", err)
	}

	// Run migrations (create tables)
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	// Seed initial data
	if err := seedData(db); err != nil {
		log.Printf("Warning: Failed to seed data: %v", err)
	}

	log.Println("Connected to SQLite database")
	return db, nil
}

func runMigrations(db *sql.DB) error {
	// Create users table
	userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`

	if _, err := db.Exec(userTable); err != nil {
		return err
	}

	// Create follows table
	followsTable := `
    CREATE TABLE IF NOT EXISTS follows (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        follower_id TEXT NOT NULL,
        followee_id TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (follower_id) REFERENCES users(id),
        FOREIGN KEY (followee_id) REFERENCES users(id),
        UNIQUE(follower_id, followee_id)
    )`

	if _, err := db.Exec(followsTable); err != nil {
		return err
	}

	return nil
}

func seedData(db *sql.DB) error {
	// Insert initial users (user1 to user5)
	users := []struct {
		id   string
		name string
	}{
		{"user1", "Alice"},
		{"user2", "Bob"},
		{"user3", "Charlie"},
		{"user4", "Diana"},
		{"user5", "Eve"},
	}

	for _, user := range users {
		_, err := db.Exec("INSERT OR IGNORE INTO users (id, name) VALUES (?, ?)", user.id, user.name)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadSchema() ([]byte, error) {
	schemaFile := "configs/schema.graphql"
	if len(os.Args) >= 2 {
		schemaFile = os.Args[1]
	}

	return os.ReadFile(schemaFile)
}

func initHTTPServer(schema *gql.Schema) *chi.Mux {
	router := chi.NewRouter()

	// GraphQL endpoint
	router.Handle("/query", &relay.Handler{Schema: schema})

	// GraphiQL interface
	router.Get("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(page))
	})

	return router
}
