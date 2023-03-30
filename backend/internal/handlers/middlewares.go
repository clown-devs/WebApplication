package handlers

import (
	"fmt"
	"net/http"
	"sberapi/internal/store"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func authentificateEmployee(next http.Handler, store store.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Authorization"]) != 1 {
			respondError(w, r, http.StatusUnauthorized, fmt.Errorf("Wrong Authorization Header"))
			return
		}
		requestToken := r.Header["Authorization"][0]
		valid, err := store.Token().Exists(requestToken)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		if !valid {
			respondError(w, r, http.StatusUnauthorized, fmt.Errorf("Invalid token"))
			return
		}

		fmt.Println(requestToken)

		next.ServeHTTP(w, r)

	})
}
