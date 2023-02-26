package server

import (
	"net/http"

	. "sberapi/internal/config"
	"sberapi/internal/model"
	"sberapi/internal/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	Logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		Logger: logrus.New(),
		router: mux.NewRouter(),
		store:  nil,
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.Logger.Info("Configuring routers...")
	s.configureRouter()

	s.Logger.Info("Configuring database...")
	if err := s.configureStore(); err != nil {
		return err
	}

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

func (s *Server) configureStore() error {
	st := store.New(s.config)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	// Debug working!!! Don't pass!
	u, err := s.store.Employee().Create(&model.Employee{
		Firstname:  "Vladimir",
		Secondname: "Putin",
		Thirdname:  "Vladimirovich",
		Password:   "ebal",
		Username:   "putin",
	})
	
	//u, err := s.store.Employee().FindByUsername("putin")
	s.Logger.Debug(u)
	if err != nil {
		return err
	}

	//---End of debugging work---
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
