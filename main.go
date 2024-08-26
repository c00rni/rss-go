package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	port string
}

func main() {
	mux := http.NewServeMux()
	lErr := godotenv.Load()
	if lErr != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	apiCfg := apiConfig{
		port: port,
	}

	mux.HandleFunc("/v1/healthz", handleHealthz)
	mux.HandleFunc("/v1/err", handleError)

	srv := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + apiCfg.port,
	}

	log.Printf("Serving on port: %s\n", apiCfg.port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Error: w%", err)
	}
}
