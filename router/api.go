package router

import (
	"boilerplate/app"
	"boilerplate/middleware"
	"github.com/go-chi/chi/v5"
)

func (a api) Init(r *chi.Mux, apiInstance app.IAPI) {
	userRouter := chi.NewRouter()
	r.Use(middleware.ApplyBasicAuthorizer(apiInstance))

	userRouter.Route("/company", func(rr chi.Router) {

	})
}
