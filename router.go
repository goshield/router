package router

import (
	"net/http"

	"github.com/goshield/interfaces"
)

// Router is a routing service
type Router interface {
	http.Handler

	ROUTE(method string, path string, middlewares ...interfaces.Middleware)
	GET(path string, middlewares ...interfaces.Middleware)
	POST(path string, middlewares ...interfaces.Middleware)
	PUT(path string, middlewares ...interfaces.Middleware)
	PATCH(path string, middlewares ...interfaces.Middleware)
	DELETE(path string, middlewares ...interfaces.Middleware)
	OPTIONS(path string, middlewares ...interfaces.Middleware)

	BeforeDispatch(...interfaces.Middleware)
	AfterDispatch(...interfaces.Middleware)
}
