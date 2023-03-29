package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sberapi/internal/store"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, store store.Store) {
	router.Use(commonMiddleware)

	RegisterAuthHandlers(router, store)
	RegisterEmployeeHandlers(router, store)
	RegisterClientHandlers(router, store)
}

func respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, map[string]string{"error": err.Error(), "code": fmt.Sprint(code)})
}

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			respondError(w, r, code, err)
		}
	}
}
