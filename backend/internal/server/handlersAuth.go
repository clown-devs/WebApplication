package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) RegisterAuthHandlers() {
	s.router.HandleFunc("/token/login/", s.handleCreateOrGetToken()).Methods("POST")
}

func (s *Server) handleCreateOrGetToken() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e, err := s.store.Employee().FindByUsername(req.Username)
		if err != nil || !e.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Incorrect username or password"))
			return
		}

		token, err := s.store.Token().Find(e.ID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if token == nil {
			token, err = s.store.Token().Create(e.ID, s.config.Salt)
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}
		s.respond(w, r, http.StatusOK, map[string]interface{}{
			"auth_token:": token.Token,
			"employee_id": token.EmployeeId,
		})
	}
}
