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

import (
  "github.com/gosmos/restful"
  "net/http"
)

// Example controller will handle calls to RESTful API
// of simple integers identified by string keys.
type MyToysController struct {
  storage map[string]interface{}
  nextKey int
}

// Index() method will be called when accessing root ("/")
// of the RESTful resource using GET method (see HTTP/1.1 spec).
// It must always return a map from string to interface{}.
// Returned map will be encoded to json and returned in the body
// of HTTP response.
func (controller *MyToysController) Index() map[string]interface{} {
  return controller.storage
}
// Create() method will be invoked when accessing root of restful
// resource using POST method. JSON object from request body will be
// decoded and passed in here as argument.
func (controller *MyToysController) Create(newElement interface{}) string {
  newElementInt := newElement.(int)
  key := string(controller.nextKey)
  controller.nextKey += 1
  controller.storage[key] = newElementInt;
  return key
}
// New() method creates new zeroed instance of resource object.
// Incoming data will be decoded from json into returned instance
// and passed to Create() method.
func (controller *MyToysController) New() interface{} {
  return 0
}

// Contructs the controller.
func NewMyToysController() *MyToysController {
  return &MyToysController { make(map[string]interface{}), 0 }
}

func main() {
  router := restful.NewRouter()
  router.AddResource("/api/mytoys", NewMyToysController())
  http.ListenAndServe(":8080", router)
}

