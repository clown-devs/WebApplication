package server

import (
	"net/http"

	. "sberapi/internal/config"
	"sberapi/internal/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  nil,
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting server...")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureStore() error {
	st := store.New(s.config)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *Server) handleIndex() http.HandlerFunc {
	// Такое определние позволяет нам определять дополнительные поля здесь, перед функцией,
	// Так называемое замыкание.. Поэтому я не передаю в роутер http функцию напрямую
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, web!"))
	}
}
