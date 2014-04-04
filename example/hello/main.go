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

// we need restful package and some http server
import (
  "github.com/gosmos/restful"
  "net/http"
)

// Example controller handles calls to RESTful API
// of simple integers identified by string keys.
type MyToysController struct {
  storage map[string]interface{}
}

// Contructs the controller.
func NewMyToysController() *MyToysController {
  return &MyToysController { make(map[string]interface{}) }
}

// Index() method will be called when accessing root ("/")
// of the RESTful resource using GET method (see HTTP/1.1 spec).
// It must always return a map from string to interface{}.
// Returned map will be encoded to json and returned in the body
// of HTTP response.
func (controller *MyToysController) Index() map[string]interface{} {
  return controller.storage
}

// Its good to check if proper interface is implemented
// to force early crash in case of method signature error.
var _ = MyToysController{}.(restful.Indexer)

func main() {
  router := restful.NewRouter()
  router.AddResource("/api/mytoys", NewMyToysController())
  http.ListenAndServe(":8080", router)
}

