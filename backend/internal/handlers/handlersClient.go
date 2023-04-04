package handlers

import (
	"encoding/json"
	"net/http"
	"sberapi/internal/model"
	"sberapi/internal/store"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterClientHandlers(router *mux.Router, store store.Store) {
	route := router.PathPrefix("/clients").Subrouter()
	authorizedRoute := router.PathPrefix("/clients").Subrouter()
	authorizedRoute.Use(authentificateEmployee(store))

	route.HandleFunc("/", handleClientCreate(store)).Methods("POST")
	route.HandleFunc("/", handleClientsGetAll(store)).Methods("GET")

	route.HandleFunc("/{id:[0-9]+}/", handleClientById(store)).Methods("GET")

	route.HandleFunc("/contacts/", handleContactCreate(store)).Methods("POST")
	route.HandleFunc("/contacts/", handleContactsGetAll(store)).Methods("GET")
}

func handleClientById(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

		client, err := store.Client().Find(id)
		if err != nil {
			respondError(w, r, http.StatusNotFound, err)
			return
		}
		respond(w, r, http.StatusOK, client)
	}
}

func handleClientCreate(store store.Store) http.HandlerFunc {
	type Request struct {
		Name string `json:"name"`
		Inn  string `json:"inn"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		data := &Request{}

		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		client := &model.Client{Name: data.Name, Inn: data.Inn}

		err := store.Client().Create(client)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		respond(w, r, http.StatusOK, client)

	}
}

func handleClientsGetAll(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := store.Client().All()
		if err != nil {
			respondError(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, clients)
	}
}

func handleContactCreate(store store.Store) http.HandlerFunc {
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
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		client := model.Contact{
			Fullname: req.Fullname,
			Client:   &model.Client{ID: req.ClientId},
			Phone:    req.Phone,
			Email:    req.Email,
			Position: req.Position,
		}

		err := store.Contact().Create(&client)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		respond(w, r, http.StatusOK, client)
	}

}

func handleContactsGetAll(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//FIXME: add filters
		filters := &model.ContactFilters{}
		contacts, err := store.Contact().All(filters)
		if err != nil {
			respondError(w, r, http.StatusInternalServerError, err)
			return
		}
		respond(w, r, http.StatusOK, contacts)
	}
}
