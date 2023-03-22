package server

import (
	"fmt"
	"net/http"
)

// TODO: Rewrite middlewares to different module
func (s *Server) commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) authentificateEmployee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Authorization"]) != 1 {
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Wrong Authorization Header"))
			return
		}
		requestToken := r.Header["Authorization"][0]

		fmt.Println(requestToken)

		next.ServeHTTP(w, r)
		//token, err := s.store.Token().Find()
	})
}
