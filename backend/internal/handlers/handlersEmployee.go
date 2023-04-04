package handlers

import (
	"encoding/json"
	"net/http"
	"sberapi/internal/model"
	"sberapi/internal/store"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func RegisterEmployeeHandlers(router *mux.Router, store store.Store) {
	employeeRoute := router.PathPrefix("/employee").Subrouter()

	authorizedRoute := router.PathPrefix("/employee").Subrouter()
	authorizedRoute.Use(authentificateEmployee(store))

	employeeRoute.HandleFunc("/", handleEmployeeCreate(store)).Methods("POST")
	employeeRoute.HandleFunc("/", handleEmployeesGetAll(store)).Methods("GET")
	//FIXME:
	//authorizedRoute.HandleFunc("/{id:[0-9]+}/", handleEmployeeById()).Methods("GET")
	employeeRoute.HandleFunc("/{id:[0-9]+}/", handleEmployeeById(store)).Methods("GET")
	authorizedRoute.HandleFunc("/current/", handleCurrentUser(store)).Methods("GET", "PUT")

	employeeRoute.HandleFunc("/directions/", handleDirectionsCreate(store)).Methods("POST")
	employeeRoute.HandleFunc("/directions/", handleDirectionsGetAll(store)).Methods("GET")
	employeeRoute.HandleFunc("/directions/{id:[0-9]+}/", handleDirectionById(store)).Methods("GET")

}

func handleEmployeesGetAll(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filters := &model.EmployeeFilters{}

		filters.DirectionId, _ = strconv.ParseUint(r.URL.Query().Get("direction_id"), 10, 64)
		filters.ClientId, _ = strconv.ParseUint(r.URL.Query().Get("client_id"), 10, 64)

		employees, err := store.Employee().All(filters)

		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusOK, employees)
	}
}

func handleEmployeeCreate(store store.Store) http.HandlerFunc {
	type request struct {
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		request := &request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.Employee{
			Fullname: request.Fullname,
			Username: request.Username,
			Password: request.Password,
		}

		if err := store.Employee().Create(e); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		e.Sanitize()
		respond(w, r, http.StatusOK, e)
	}
}

func handleEmployeeById(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		e, err := store.Employee().Find(id)
		if err != nil {
			respondError(w, r, http.StatusNotFound, err)
			return
		}
		respond(w, r, http.StatusOK, e)
	}
}

func handleCurrentUser(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.Fields(r.Header["Authorization"][0])[1]
		e, err := store.Employee().FindByToken(token)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		if r.Method == "GET" {
			respond(w, r, http.StatusOK, e)

		} else if r.Method == "PUT" {
			if err := json.NewDecoder(r.Body).Decode(e); err != nil {
				respondError(w, r, http.StatusBadRequest, err)
				return
			}

			err := store.Employee().UpdatePublic(e)
			if err != nil {
				respondError(w, r, http.StatusBadRequest, err)
				return
			}
			respond(w, r, http.StatusOK, e)
		}
	}
}

func handleDirectionById(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		dir, err := store.Direction().Find(id)
		if err != nil {
			respondError(w, r, http.StatusNotFound, err)
			return
		}
		respond(w, r, http.StatusOK, dir)

	}
}

func handleDirectionsCreate(store store.Store) http.HandlerFunc {
	type Request struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		request := &Request{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		direction := &model.Direction{Name: request.Name}
		if err := store.Direction().Create(direction); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusOK, direction)

	}
}

func handleDirectionsGetAll(store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		directions, err := store.Direction().All()
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusOK, directions)
		return

	}

}
