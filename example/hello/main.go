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

package main

// We will need only gosmos/restful package and some http server.
import (
  "github.com/gosmos/restful"
  "net/http"
)

// First thing to do, is to define controller structure, which will be
// responsible for handling calls to RESTful API.
type MyToysController struct {
  // Memory storage implemented on map is used for simplicity.
  // Normally we would want some persistence in here.
  storage map[string]interface{}
}
// Contructor function.
func NewMyToysController() *MyToysController {
  controller := &MyToysController { make(map[string]interface{}) }
  return controller
}

// Our controller needs to implement at least one interface for handling
// REST calls (check controller.go for all possible interfaces).
// Interface restful.Indexer will be our choise for this example.

// Index() method will be called by the library on each access to root
// of our RESTful resource ("/") that uses GET method (see HTTP/1.1 spec).
// It must return a map of type map[string]interface{}. Returned map
// will be encoded to JSON and written into the body of HTTP response.
func (controller *MyToysController) Index() map[string]interface{} {
  return controller.storage
}

// It's good to check if proper interface is implemented by the controller,
// so we will face early crash in case of method signature error.
var _ restful.Indexer = &MyToysController{}

// Lastly, we need to create router, configure our RESTful resource,
// and start HTTP server. Voilà!
func main() {
  router := restful.NewRouter()
  router.AddResource("/api/mytoys", NewMyToysController())
  http.ListenAndServe(":8080", router)
}

