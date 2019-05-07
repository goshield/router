package router

import (
	"net/http"

	. "github.com/goshield/tools"
)

func Params(r *http.Request) Bag {
	pb, ok := r.Context().Value(KeyRequestParams).(Bag)
	if ok {
		return pb
	}
	return NewBag()
}