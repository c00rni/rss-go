package main

import (
	"encoding/json"
	"github.com/c00rni/rss-go/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}

	type response struct {
		Id         uuid.UUID `json:"id"`
		CreateDate time.Time `json:"created_at"`
		UpdateDate time.Time `json:"updated_at"`
		Name       string    `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	inputData := request{}
	err := decoder.Decode(&inputData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	today := time.Now()

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: today,
		UpdatedAt: today,
		Name:      inputData.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}
