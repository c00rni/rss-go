package main

import (
	"github.com/c00rni/rss-go/internal/database"
	"log"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := extractAuthorization(r, "ApiKey ")
		if err != nil {
			log.Println(err)
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get the user form thr database.")
			return
		}

		handler(w, r, user)
	})
}
