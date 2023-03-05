package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sberapi/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterEmployeeHandlers() {
	employeeRoute := s.router.PathPrefix("/employee").Subrouter()

	protectedRoute := s.router.PathPrefix("/employee").Subrouter()
	protectedRoute.Use(s.authentificateEmployee)

	employeeRoute.HandleFunc("/", s.handleEmployeeCreate()).Methods("POST")
	protectedRoute.HandleFunc("/{id:[0-9]+}/", s.handleEmployeeById()).Methods("GET")

	employeeRoute.HandleFunc("/directions/", s.handleDirections()).Methods("POST", "GET")
	employeeRoute.HandleFunc("/directions/{id:[0-9]+}/", s.handleDirectionById()).Methods("GET")
	
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
		e, err := s.store.Employee().Find(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		e.Sanitize()
		s.respond(w, r, http.StatusOK, e)
	}
}

func (s *Server) handleDirectionById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		dir, err := s.store.Direction().Find(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.respond(w, r, http.StatusOK, dir)

	}
}
func (s *Server) handleDirections() http.HandlerFunc {

	postRequest := &struct {
		Name string `json:"name"`
	}{}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := json.NewDecoder(r.Body).Decode(postRequest); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			direction := &model.Direction{Name: postRequest.Name}

			if err := s.store.Direction().Create(direction); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			s.respond(w, r, http.StatusOK, direction)

		} else if r.Method == "GET" {
			directions, err := s.store.Direction().All()
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			s.respond(w, r, http.StatusOK, directions)
			return

		} else {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Wrong HTTP method"))
		}

	}
}
