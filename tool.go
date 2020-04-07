package router

import (
	"net/http"

	"github.com/goshield/interfaces"
	"github.com/goshield/tools"
)

type requestBagKey struct{}

// RequestBag returns an instance of Bag in request
func RequestBag(r *http.Request) interfaces.Bag {
	pb, ok := r.Context().Value(requestBagKey{}).(interfaces.Bag)
	if ok {
		return pb
	}
	return tools.NewBag()
}
