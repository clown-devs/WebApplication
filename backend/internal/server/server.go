package server

import (
	"database/sql"
	"net/http"

	. "sberapi/internal/config"
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
		//router: mux.NewRouter(), FIXME:
		store: nil,
	}
}

// Function for starting HTTP. Don't start if using struct in tests
func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.Logger.Info("Configuring database...")
	err := s.configureStore()
	if err != nil {
		return err
	}

	s.Logger.Info("Configuring routers...")
	s.configureRouter() // routes.go

	s.Logger.Info("Server started on port ", s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
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
	return nil
}
