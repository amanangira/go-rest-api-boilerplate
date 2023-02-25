package middleware

import (
	"net/http"
	"superman/api"
)

type IMiddleware interface {
	Apply(a api.IAPI) func(http.Handler) http.Handler
}
