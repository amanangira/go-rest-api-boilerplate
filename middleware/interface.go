package middleware

import (
	"net/http"

	"boilerplate/app"
)

type IMiddleware interface {
	Apply(a app.IAPI) func(http.Handler) http.Handler
}
