package router

import (
	"context"
	"net/http"

	"github.com/goshield/interfaces"
	"github.com/goshield/tools"
	"github.com/julienschmidt/httprouter"
)

type httpRouter struct {
	httprouter                *httprouter.Router
	errorHandler              interfaces.ErrorHandler
	beforeDispatchMiddlewares []interfaces.Middleware
	afterDispatchMiddlewares  []interfaces.Middleware
}

// NewHTTPRouter returns an instance of Route using httprouter package
func NewHTTPRouter(errorHandler interfaces.ErrorHandler) Router {
	return &httpRouter{
		httprouter:                httprouter.New(),
		errorHandler:              errorHandler,
		beforeDispatchMiddlewares: make([]interfaces.Middleware, 0),
		afterDispatchMiddlewares:  make([]interfaces.Middleware, 0),
	}
}

func (hR *httpRouter) BeforeDispatch(middlewares ...interfaces.Middleware) {
	hR.beforeDispatchMiddlewares = append(hR.beforeDispatchMiddlewares, middlewares...)
}

func (hR *httpRouter) AfterDispatch(middlewares ...interfaces.Middleware) {
	hR.afterDispatchMiddlewares = append(hR.afterDispatchMiddlewares, middlewares...)
}

func (hR *httpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hR.httprouter.ServeHTTP(w, r)
}

func (hR *httpRouter) ROUTE(method string, path string, middlewares ...interfaces.Middleware) {
	mws := make([]interfaces.Middleware, 0)
	mws = append(mws, hR.beforeDispatchMiddlewares...)
	mws = append(mws, middlewares...)
	mws = append(mws, hR.afterDispatchMiddlewares...)
	hR.httprouter.Handle(method, path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		rb := tools.NewBag()
		for _, param := range params {
			rb.Set(param.Key, param.Value)
		}
		r = r.WithContext(context.WithValue(r.Context(), requestBagKey{}, rb))
		for _, mw := range mws {
			ctx, err := mw(w, r)
			if ctx != nil {
				r = r.WithContext(ctx)
			}
			if err != nil {
				if hR.errorHandler != nil {
					hR.errorHandler(w, r, err)
				} else {
					panic(err)
				}
				return
			}
		}
	})
}

func (hR *httpRouter) GET(path string, middlewares ...interfaces.Middleware) {
	hR.ROUTE(http.MethodGet, path, middlewares...)
}

func (hR *httpRouter) POST(path string, middlewares ...interfaces.Middleware) {
	hR.ROUTE(http.MethodPost, path, middlewares...)
}

func (hR *httpRouter) PUT(path string, middlewares ...interfaces.Middleware) {
	hR.ROUTE(http.MethodPut, path, middlewares...)
}

func (hR *httpRouter) PATCH(path string, middlewares ...interfaces.Middleware) {
	hR.ROUTE(http.MethodPatch, path, middlewares...)
}

func (hR *httpRouter) DELETE(path string, middlewares ...interfaces.Middleware) {
	hR.ROUTE(http.MethodDelete, path, middlewares...)
}

func (hR *httpRouter) OPTIONS(path string, middlewares ...interfaces.Middleware) {
	hR.ROUTE(http.MethodOptions, path, middlewares...)
}
