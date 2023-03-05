package server

import (
	"fmt"
	"net/http"
)

func (s *Server) authentificateEmployee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestToken := r.Header["Authorization"]
		fmt.Println(requestToken)

		next.ServeHTTP(w, r)
		//token, err := s.store.Token().Find()
	})
}
