package server

import "net/http"

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/users", s.handleIndex())
}

// FIXME:test handle, will removed later
func (s *Server) handleIndex() http.HandlerFunc {
	// Такое определние позволяет нам определять дополнительные поля здесь, перед функцией,
	// Так называемое замыкание.. Поэтому я не передаю в роутер http функцию напрямую
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, web!"))
	}
}
