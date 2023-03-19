package server

import (
	"encoding/json"
	"net/http"
	"sberapi/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterClientHandlers() {
	route := s.router.PathPrefix("/clients").Subrouter()
	authorizedRoute := s.router.PathPrefix("/clients").Subrouter()
	authorizedRoute.Use(s.authentificateEmployee)

	route.HandleFunc("/", s.handleClientCreate()).Methods("POST")
	route.HandleFunc("/", s.handleClientsGetAll()).Methods("GET")

	route.HandleFunc("/{id:[0-9]+}/", s.handleClientById()).Methods("GET")

	route.HandleFunc("/contacts/", s.handleContactCreate()).Methods("POST")

}

func (s *Server) handleClientById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

		client, err := s.store.Client().Find(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.respond(w, r, http.StatusOK, client)
	}
}

func (s *Server) handleClientCreate() http.HandlerFunc {
	type Request struct {
		Name string `json:"name"`
		Inn  string `json:"inn"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		data := &Request{}

		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		client := &model.Client{Name: data.Name, Inn: data.Inn}

		err := s.store.Client().Create(client)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, client)

	}
}

func (s *Server) handleClientsGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := s.store.Client().All()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, clients)
	}
}

func (s *Server) handleContactCreate() http.HandlerFunc {
	type Request struct {
		Fullname string  `json:"fullname"`
		ClientId uint64  `json:"client_id"`
		Phone    *string `json:"phone,omitempty"`
		Email    *string `json:"email,omitempty"`
		Position *string `json:"position,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		client := model.Contact{
			Fullname: req.Fullname,
			Client:   &model.Client{ID: req.ClientId},
			Phone:    req.Phone,
			Email:    req.Email,
			Position: req.Position,
		}

		err := s.store.Contact().Create(&client)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, client)
	}

}
