package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"boilerplate/app"
	"boilerplate/app/container"
	middleware2 "boilerplate/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

var isDev *bool
var port string

func init() {
	isDev = flag.Bool("dev", false, "To run server in dev mode")
	flag.Parse()

	port = ""
	if *isDev {
		if err := godotenv.Load(); err != nil {
			log.Printf(err.Error())
		}
	}

	if envPort, exists := os.LookupEnv(app.EnvServerPortKey); exists {
		port = envPort
	} else {
		port = "80"
	}

}

func newMainRouter() *chi.Mux {
	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware2.ApplyPanicRecovery)
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	return mainRouter
}

func newAPIRouter() *chi.Mux {
	router := chi.NewRouter()
	c := container.NewContainer()

	router.Route("/user", func(rr chi.Router) {
		rr.Post("/", c.UserController.Create)
		rr.Get("/{user_id}", c.UserController.Get)
	})

	return router
}

func startServer(router http.Handler) {
	fmt.Printf("\nServer started on port %s. Do ctrl+c to exit... \n", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}

func muxDebugLogger(router *chi.Mux) error {
	return chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: %s has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
}

func main() {
	// Setup main router
	mainRouter := newMainRouter()

	// Setup routes
	mainRouter.Mount("/api", newAPIRouter())
	mainRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	if app.IsDebug() {
		if logErr := muxDebugLogger(mainRouter); logErr != nil {
			log.Print(logErr)
		}
	}

	startServer(mainRouter)
}
