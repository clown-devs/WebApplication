package server

import (
	"encoding/json"
	"net/http"
	"sberapi/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterEmployeeHandlers() {
	s.router.HandleFunc("/api/v3/employee/", s.handleEmployeeCreate()).Methods("POST")
	s.router.HandleFunc("/api/v3/employee/{id:[0-9]+}/", s.handleEmployeeById()).Methods("GET")
}

func (s *Server) handleEmployeeCreate() http.HandlerFunc {
	type request struct {
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.Employee{
			Fullname: request.Fullname,
			Username: request.Username,
			Password: request.Password,
		}

		if err := s.store.Employee().Create(e); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		e.Sanitize()
		s.respond(w, r, http.StatusOK, e)
	}
}

func (s *Server) handleEmployeeById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		e, err := s.store.Employee().Find(uint64(id))
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		e.Sanitize()
		s.respond(w, r, http.StatusOK, e)
	}
}
