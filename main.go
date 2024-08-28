package main

import _ "github.com/lib/pq"

import (
	"database/sql"
	"github.com/c00rni/rss-go/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	mux := http.NewServeMux()
	lErr := godotenv.Load()
	if lErr != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error: w%", err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux.HandleFunc("/v1/healthz", handleHealthz)
	mux.HandleFunc("/v1/err", handleError)
	mux.HandleFunc("POST /users", apiCfg.handleCreateUser)
	mux.HandleFunc("GET /users", apiCfg.middlewareAuth(handleGetUser))
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))

	srv := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}

	log.Printf("Serving on port: %s\n", port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Error: w%", err)
	}
}
