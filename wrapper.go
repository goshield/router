package router

import (
	"context"
	"net/http"
	
	. "github.com/goshield/interfaces"
	. "github.com/goshield/tools"
	rr "github.com/julienschmidt/httprouter"
)

func Wrap(eh ErrorHandler, mws ...Middleware) rr.Handle {
	return func(w http.ResponseWriter, r *http.Request, pr rr.Params) {
		pb := NewBag()
		for _, param := range pr {
			pb.Set(param.Key, param.Value)
		}
		r = r.WithContext(context.WithValue(r.Context(), KeyRequestParams, pb))
		for _, mw := range mws {
			ctx, err := mw(w, r)
			if ctx != nil {
				r = r.WithContext(ctx)
			}
			if err != nil && eh != nil {
				eh(w, r, err)
				return
			}
		}
	}
}