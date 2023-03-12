package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// DI and Middleware
func (s *Server) configureRouter() {
	s.router = mux.NewRouter().PathPrefix("/api/v3").Subrouter()
	s.router.Use(s.commonMiddleware)

	s.RegisterEmployeeHandlers()
	s.RegisterAuthHandlers()
	s.RegisterClientHandlers()
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error(), "code": fmt.Sprint(code)})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.error(w, r, code, err)
		}
	}
}
