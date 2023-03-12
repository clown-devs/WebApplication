package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sberapi/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterClientHandlers() {
	route := s.router.PathPrefix("/clients").Subrouter()
	authorizedRoute := s.router.PathPrefix("/clients").Subrouter()
	authorizedRoute.Use(s.authentificateEmployee)

	route.HandleFunc("/", s.handleClients()).Methods("POST", "GET")
	route.HandleFunc("/{id:[0-9]+}/", s.handleClientById()).Methods("GET")

	route.HandleFunc("/contacts/", s.handleContacts()).Methods("POST")

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

func (s *Server) handleClients() http.HandlerFunc {
	type PostRequest struct {
		Name string `json:"name"`
		Inn  string `json:"inn"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			data := &PostRequest{}

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

		if r.Method == "GET" {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("GET is not implemented yet"))
		}
	}
}

func (s *Server) handleContacts() http.HandlerFunc {
	type PostData struct {
		Fullname string  `json:"fullname"`
		ClientId uint64  `json:"client_id"`
		Phone    *string `json:"phone,omitempty"`
		Email    *string `json:"email,omitempty"`
		Position *string `json:"position,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			postData := &PostData{}
			if err := json.NewDecoder(r.Body).Decode(postData); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			client := model.Contact{
				Fullname: postData.Fullname,
				Client:   &model.Client{ID: postData.ClientId},
				Phone:    postData.Phone,
				Email:    postData.Email,
				Position: postData.Position,
			}

			err := s.store.Contact().Create(&client)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			s.respond(w, r, http.StatusOK, client)
		}
	}
}
