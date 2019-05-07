package router

import (
	rr "github.com/julienschmidt/httprouter"
)

var (
	r *rr.Router
)

func Router() *rr.Router {
	if r == nil {
		r = rr.New()
	}

	return r
}