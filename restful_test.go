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
  "net/http"
  "net/http/httptest"
  "testing"
  "strings"
  "fmt"
)

type testLister struct {
  TimesCalled int
  ReturnedMap map[string]interface{}
}
func NewTestLister() *testLister {
  return &testLister{ 0, make(map[string]interface{}) }
}
func (lister *testLister) All() map[string]interface{} {
  lister.TimesCalled += 1
  return lister.ReturnedMap
}

func TestAllMethodCalled(t *testing.T) {
  fakeController := NewTestLister()

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  if fakeController.TimesCalled != 1 {
    t.Errorf("Expected 1 call, got %v.", fakeController.TimesCalled)
  }
}

func TestJsonResponseAfterReturningEmptyMapFromAll(t *testing.T) {
  fakeController := NewTestLister()

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  expectedJson := "{}"
  actualJson := strings.TrimSpace(string(w.Body.Bytes()))
  if expectedJson != actualJson {
    t.Errorf("Expected response '%v', got '%v'.", expectedJson, actualJson)
  }
}

func TestJsonResponseAfterReturningEmptyMapWithOneString(t *testing.T) {
  fakeController := NewTestLister()
  id0 := "test"
  fakeController.ReturnedMap[id0] = id0

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  expectedJson := fmt.Sprintf("{\"%v\":\"%v\"}", id0, id0)
  actualJson := strings.TrimSpace(string(w.Body.Bytes()))
  if expectedJson != actualJson {
    t.Errorf("Expected response '%v', got '%v'.", expectedJson, actualJson)
  }
}

type testStruct struct {
  Test string `json:"test"`
}

func TestJsonResponseAfterReturningEmptyMapWithTwoStructs(t *testing.T) {
  fakeController := NewTestLister()
  var (
    id0 string = "0"
    id1 string = "1"
  )
  fakeController.ReturnedMap[id0] = testStruct{id0}
  fakeController.ReturnedMap[id1] = testStruct{id1}

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  expectedJson := fmt.Sprintf("{\"%v\":{\"test\":\"%v\"},\"%v\":{\"test\":\"%v\"}}", id0, id0, id1, id1)
  actualJson := strings.TrimSpace(string(w.Body.Bytes()))
  if expectedJson != actualJson {
    t.Errorf("Expected response '%v', got '%v'.", expectedJson, actualJson)
  }
}

func TestPanicWhenReturningNilMapFromLister(t *testing.T) {
  fakeController := new(testLister) // ReturnedMap is nil

  router := NewRouter()
  w := httptest.NewRecorder()

  defer func() {
    if err := recover(); err == nil {
      t.Error("Should panic when returned map is nil, but didnt.")
    }
  }()
  createResourceAndServeARequest(router, "/test", "/", fakeController, w)
}

func createResourceAndServeARequest(router *Router,
    resource string, request string, controller interface{},
    out http.ResponseWriter) {

  router.AddResource(resource, controller)

  url := "http://www.domain.com"+ resource + request
  req, _ := http.NewRequest("GET", url, nil)

  router.ServeHTTP(out, req)
}

