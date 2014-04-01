package restful

import (
  "github.com/gorilla/mux"
  "net/http"
  "log"
)

type Router struct {
  impl *mux.Router
}

func NewRouter() *Router {
  return &Router{ mux.NewRouter() }
}

func (router *Router) AddResource(path string, controller interface{}) {
  api := router.impl.PathPrefix(path).Subrouter()
  registered := false

  if lister, ok := controller.(Lister); ok {
    api.Path("/").Methods("GET").HandlerFunc(NewListerHandler(lister))
    registered = true
  }
  if getter, ok := controller.(Getter); ok {
    api.Path("/{id:[0-9]+}").Methods("GET").HandlerFunc(NewGetterHandler(getter))
    registered = true
  }
  if adder, ok := controller.(Adder); ok {
    api.Path("/").Methods("POST").HandlerFunc(NewAdderHandler(adder))
    registered = true
  }
  if replacer, ok := controller.(Replacer); ok {
    api.Path("/{id:[0-9]+}").Methods("PUT").HandlerFunc(NewReplacerHandler(replacer))
    registered = true
  }
  if deleter, ok := controller.(Deleter); ok {
    api.Path("/{id:[0-9]+}").Methods("DELETE").HandlerFunc(NewDeleterHandler(deleter))
    registered = true
  }

  if !registered {
    log.Panicf("controller %#v doesn't implement any REST interfaces", controller)
  }
}

func (router *Router) ServeHTTP(out http.ResponseWriter, in *http.Request) {
  router.impl.ServeHTTP(out, in)
}

