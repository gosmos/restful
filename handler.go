package restful

import (
  "github.com/gorilla/mux"
  "encoding/json"
  "net/http"
  "strconv"
  "log"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func NewListerHandler(lister Lister) HandlerFunc {
  return func (out http.ResponseWriter, req *http.Request) {
    elems := lister.All()

    bytes, _ := json.Marshal(elems)
    if _, err := out.Write(bytes); err != nil {
      panic(err)
    }
  }
}

func NewGetterHandler(getter Getter) HandlerFunc {
  return func (out http.ResponseWriter, in *http.Request) {
    id := getResourceId(in)

    elem := getter.Get(id)

    bytes, _ := json.Marshal(map[int64]interface{} { id : elem })
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

    bytes, _ := json.Marshal(map[int64]interface{} { id : elem })
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

    bytes, _ := json.Marshal(map[int64]interface{} { id : elem })
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

func getResourceId(in *http.Request) int64 {
  idStr, ok := mux.Vars(in)["id"]
  if !ok {
    log.Panicln("resource id not found in request data")
  }
  id, err := strconv.ParseInt(idStr, 10, 64)
  if err != nil {
    log.Panicf("resource id is not an integer: %d (%s)", idStr, err)
  }
  return id
}

