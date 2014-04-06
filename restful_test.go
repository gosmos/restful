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

type testIndexer struct {
  TimesCalled int
  ReturnedMap map[string]interface{}
}
func NewTestIndexer() *testIndexer {
  return &testIndexer{ 0, make(map[string]interface{}) }
}
func (indexer *testIndexer) Index() map[string]interface{} {
  indexer.TimesCalled += 1
  return indexer.ReturnedMap
}

func TestAllMethodCalled(t *testing.T) {
  fakeController := NewTestIndexer()

  router, w := NewRouter(), httptest.NewRecorder()
  createResourceAndServeARequest(router, "/test", "/", fakeController, w)

  if fakeController.TimesCalled != 1 {
    t.Errorf("Expected 1 call, got %v.", fakeController.TimesCalled)
  }
}

func TestJsonResponseAfterReturningEmptyMapFromAll(t *testing.T) {
  fakeController := NewTestIndexer()

  router, w := NewRouter(), httptest.NewRecorder()
  createResourceAndServeARequest(router, "/test", "/", fakeController, w)

  expectedJson := "{}"
  actualJson := strings.TrimSpace(string(w.Body.Bytes()))
  if expectedJson != actualJson {
    t.Errorf("Expected response '%v', got '%v'.", expectedJson, actualJson)
  }
}

func TestJsonResponseAfterReturningEmptyMapWithOneString(t *testing.T) {
  fakeController := NewTestIndexer()
  id0 := "test"
  fakeController.ReturnedMap[id0] = id0

  router, w := NewRouter(), httptest.NewRecorder()
  createResourceAndServeARequest(router, "/test", "/", fakeController, w)

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
  fakeController := NewTestIndexer()
  id0, id1 := "0", "1"
  fakeController.ReturnedMap[id0] = testStruct{id0}
  fakeController.ReturnedMap[id1] = testStruct{id1}

  router, w := NewRouter(), httptest.NewRecorder()
  createResourceAndServeARequest(router, "/test", "/", fakeController, w)

  expectedJson := fmt.Sprintf("{\"%v\":{\"test\":\"%v\"},\"%v\":{\"test\":\"%v\"}}", id0, id0, id1, id1)
  actualJson := strings.TrimSpace(string(w.Body.Bytes()))
  if expectedJson != actualJson {
    t.Errorf("Expected response '%v', got '%v'.", expectedJson, actualJson)
  }
}

func TestPanicWhenReturningNilMapFromIndexer(t *testing.T) {
  fakeController := new(testIndexer) // ReturnedMap is nil

  router, w := NewRouter(), httptest.NewRecorder()

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

  router.HandleResource(resource, controller)

  url := "http://www.domain.com"+ resource + request
  req, _ := http.NewRequest("GET", url, nil)

  router.ServeHTTP(out, req)
}

