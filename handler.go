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
  "encoding/json"
  "net/http"
  "log"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func NewListerHandler(lister Lister) HandlerFunc {
  return func (out http.ResponseWriter, req *http.Request) {
    elems := lister.All()
    if elems == nil {
      log.Panicf("nil map returned from %#v\n", lister)
    }

    encoder := json.NewEncoder(out)
    if err := encoder.Encode(elems); err != nil {
      panic(err)
    }
  }
}

func NewGetterHandler(getter Getter) HandlerFunc {
  return func (out http.ResponseWriter, in *http.Request) {
    id := getResourceId(in)

    elem := getter.Get(id)

    bytes, _ := json.Marshal(map[string]interface{} { id : elem })
    if _, err := out.Write(bytes); err != nil {
      panic(err)
    }
  }
}

func NewAdderHandler(adder Adder) HandlerFunc {
  return func (out http.ResponseWriter, in *http.Request) {
    decoder := json.NewDecoder(in.Body)
    elem := adder.New()
    if elem == nil {
      log.Panicf("Creator.New on %#v returned nil", adder)
    }
    if err := decoder.Decode(&elem); err != nil {
      panic(err)
    }

    id := adder.Add(elem)

    bytes, _ := json.Marshal(map[string]interface{} { id : elem })
    if _, err := out.Write(bytes); err != nil {
      panic(err)
    }
  }
}

func NewReplacerHandler(replacer Replacer) HandlerFunc {
  return func (out http.ResponseWriter, in *http.Request) {
    id := getResourceId(in)

    decoder := json.NewDecoder(in.Body)
    elem := replacer.New()
    if elem == nil {
      log.Panicf("Creator.New on %#v returned nil", replacer)
    }
    if err := decoder.Decode(&elem); err != nil {
      panic(err)
    }

    replacer.Replace(id, elem)

    bytes, _ := json.Marshal(map[string]interface{} { id : elem })
    if _, err := out.Write(bytes); err != nil {
      panic(err)
    }
  }
}

func NewDeleterHandler(deleter Deleter) HandlerFunc {
  return func (out http.ResponseWriter, in *http.Request) {
    id := getResourceId(in)

    ok := deleter.Delete(id)

    bytes, _ := json.Marshal(map[string]bool { "ok": ok })
    if _, err := out.Write(bytes); err != nil {
      panic(err)
    }
  }
}

func getResourceId(in *http.Request) string {
  id, ok := mux.Vars(in)["id"]
  if !ok {
    log.Panicln("resource id not found in request data")
  }
  return id
}

