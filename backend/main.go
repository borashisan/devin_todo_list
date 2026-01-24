package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/internal/handler"
	"backend/internal/infrastructure/db"
	"backend/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	database, err := db.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Println("Successfully connected to database")

	// Setup layers (dependency injection)
	todoRepo := db.NewTodoRepository(database)
	repoAdapter := db.NewTodoRepositoryAdapter(todoRepo)
	todoUsecase := usecase.NewTodoUsecase(repoAdapter)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	// Setup router
	router := handler.NewRouter(todoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
