package app

import (
	"fmt"
	"github.com/go-chi/chi"
)

// Initialize the router
func (s *Server) InitRouter() chi.Router {
	r := chi.NewRouter()
	pattern := fmt.Sprintf("/%s/%s", s.Name, s.Version)
	r.Mount(pattern, InitApi(s))
	return r
}

func InitApi(s *Server) chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {

	})

	return r

}
