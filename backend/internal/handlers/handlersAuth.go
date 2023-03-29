package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sberapi/internal/store"

	"github.com/gorilla/mux"
)

func RegisterAuthHandlers(router *mux.Router, store store.Store) {
	router.HandleFunc("/token/login/", handleCreateOrGetToken(store)).Methods("POST")
}

func handleCreateOrGetToken(store store.Store) http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		e, err := store.Employee().FindByUsername(req.Username)
		if err != nil || !e.ComparePassword(req.Password) {
			respondError(w, r, http.StatusUnauthorized, fmt.Errorf("Incorrect username or password"))
			return
		}

		token, err := store.Token().Find(e.ID)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		if token == nil {
			token, err = store.Token().Create(e.ID, "SaltySalt")
			if err != nil {
				respondError(w, r, http.StatusInternalServerError, err)
				return
			}
		}
		respond(w, r, http.StatusOK, map[string]interface{}{
			"auth_token:": token.Token,
			"employee":    e,
		})
	}
}
