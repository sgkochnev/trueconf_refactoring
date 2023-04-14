package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type router struct {
	chi.Router
}

func NewRouter(uCase UseCase) *router {
	r := &router{
		chi.NewRouter(),
	}

	r.endpoints(uCase)

	return r
}

func (r *router) endpoints(uCase UseCase) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	userCnt := NewUserController(uCase)

	rUsers := func(r chi.Router) {
		r.Get("/", userCnt.SearchUsers)
		r.Post("/", userCnt.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", userCnt.GetUser)
			r.Patch("/", userCnt.UpdateUser)
			r.Delete("/", userCnt.DeleteUser)
		})
	}

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", rUsers)
		})
	})
}
