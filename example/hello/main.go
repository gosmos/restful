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

// First thing to do, is defining controller structure, which is responsible
// for handling calls to RESTful API of objects identified by string keys.
type MyToysController struct {
  // Memory storage implemented on map is used for simplicity.
  // Normally we would want some persistence in here.
  storage map[string]interface{}
}
// Contructs the controller.
func NewMyToysController() *MyToysController {
  controller := &MyToysController { make(map[string]interface{}) }
  return controller
}

// Our controller needs at least one method for handling REST calls.

// Index() method will be called when accessing root ("/") of our RESTful
// resource using GET method (see HTTP/1.1 spec). It must return a map 
// of type map[string]interface{}. Returned map will be encoded to json
// and written into body of HTTP response.
func (controller *MyToysController) Index() map[string]interface{} {
  return controller.storage
}

// Its good to check if proper interface is implemented (in this case
// restful.Indexer) to force early crash in case of method signature error.
var _ restful.Indexer = &MyToysController{}

// Lastly, we create router, configure our RESTful resource, and start
// HTTP server. Voilà!
func main() {
  router := restful.NewRouter()
  router.AddResource("/api/mytoys", NewMyToysController())
  http.ListenAndServe(":8080", router)
}

