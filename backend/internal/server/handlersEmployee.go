package server

import (
	"encoding/json"
	"net/http"
	"sberapi/internal/model"
)

func (s *Server) RegisterEmployeeHandlers() {
	s.router.HandleFunc("/api/v3/employee/", s.handleEmployeeCreate())
}

func (s *Server) handleEmployeeCreate() http.HandlerFunc {
	type request struct {
		Firstname  string `json:"first_name"`
		Secondname string `json:"second_name"`
		Thirdname  string `json:"third_name"`
		Username   string `json:"username"`
		Password   string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.Employee{
			Firstname:  request.Firstname,
			Secondname: request.Secondname,
			Thirdname:  request.Thirdname,
			Username:   request.Username,
			Password:   request.Password,
		}

		if err := s.store.Employee().Create(e); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		e.Sanitize()
		s.respond(w, r, http.StatusOK, e)
	}
}
