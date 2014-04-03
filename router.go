/*
   Copyright 2014 Maciej Chałapuk

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package restful

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
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

  if indexer, ok := controller.(Indexer); ok {
    api.Path("/").Methods("GET").HandlerFunc(NewIndexerHandler(indexer))
    registered = true
  }
  if shower, ok := controller.(Shower); ok {
    api.Path("/{id}").Methods("GET").HandlerFunc(NewShowerHandler(shower))
    registered = true
  }
  if creator, ok := controller.(Creator); ok {
    api.Path("/").Methods("POST").HandlerFunc(NewCreatorHandler(creator))
    registered = true
  }
  if updater, ok := controller.(Updater); ok {
    api.Path("/{id}").Methods("PUT").HandlerFunc(NewUpdaterHandler(updater))
    registered = true
  }
  if deleter, ok := controller.(Deleter); ok {
    api.Path("/{id}").Methods("DELETE").HandlerFunc(NewDeleterHandler(deleter))
    registered = true
  }

  if !registered {
    fmt.Errorf("%#v doesn't implement any REST interfaces", controller)
  }
}

func (router *Router) ServeHTTP(out http.ResponseWriter, in *http.Request) {
  router.impl.ServeHTTP(out, in)
}

