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
  log.Printf("adding resource %v", path)

  api := router.impl.PathPrefix(path).Subrouter()
  registered := make([]string, 0, 5)

  if lister, ok := controller.(Lister); ok {
    api.Path("/").Methods("GET").HandlerFunc(NewListerHandler(lister))
    registered = append(registered, "Lister")
  }
  if getter, ok := controller.(Getter); ok {
    api.Path("/{id:[0-9]+}").Methods("GET").HandlerFunc(NewGetterHandler(getter))
    registered = append(registered, "Getter")
  }
  if adder, ok := controller.(Adder); ok {
    api.Path("/").Methods("POST").HandlerFunc(NewAdderHandler(adder))
    registered = append(registered, "Adder")
  }
  if replacer, ok := controller.(Replacer); ok {
    api.Path("/{id:[0-9]+}").Methods("PUT").HandlerFunc(NewReplacerHandler(replacer))
    registered = append(registered, "Replacer")
  }
  if deleter, ok := controller.(Deleter); ok {
    api.Path("/{id:[0-9]+}").Methods("DELETE").HandlerFunc(NewDeleterHandler(deleter))
    registered = append(registered, "Deleter")
  }

  if len(registered) == 0 {
    log.Panicf("controller %#v doesn't implement any REST interfaces", controller)
  }
  log.Printf("registered controller %#v implementing interfaces %v", controller, registered)
}

func (router *Router) ServeHTTP(out http.ResponseWriter, in *http.Request) {
  router.impl.ServeHTTP(out, in)
}

