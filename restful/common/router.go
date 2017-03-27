package common

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	*httprouter.Router
}

func wrapHandler(h http.Handler) httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
	return httprouter.Handle(fn)
}

func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

func (r *Router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
}
