package server

import (
	"database/sql"
	"net/http"

	. "sberapi/internal/config"
	"sberapi/internal/model"
	"sberapi/internal/store"
	"sberapi/internal/store/sqlstore"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	Logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		Logger: logrus.New(),
		router: mux.NewRouter(),
		store:  nil,
	}
}

// Function for starting HTTP. Don't start if using struct in tests
func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.Logger.Info("Configuring routers...")
	s.configureRouter() // routes.go

	s.Logger.Info("Configuring database...")
	s.configureStore()

	s.Logger.Info("Server started")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)
	return nil
}

func newDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil { // database ping
		return nil, err
	}

	return db, nil
}

func (s *Server) configureStore() error {
	db, err := newDB(s.config.DbConnStr)
	if err != nil {
		return err
	}
	s.store = sqlstore.New(db)

	// Debug working!!! Don't pass!
	u := &model.Employee{
		Firstname:  "Vladimir",
		Secondname: "Putin",
		Thirdname:  "Vladimirovich",
		Password:   "ebal",
		Username:   "putin",
	}

	err = s.store.Employee().Create(u)

	u, _ = s.store.Employee().Find(u.ID)
	s.Logger.Debug(u)
	if err != nil {
		return err
	}

	//---End of debugging work---
	return nil
}
