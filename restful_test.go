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
  "bytes"
  "net/http"
  "encoding/json"
  "net/http/httptest"
  "testing"
)

type testLister struct {
  TimesCalled int
  ReturnedMap map[int64]interface{}
}
func (lister *testLister) All() map[int64]interface{} {
  lister.TimesCalled += 1
  return lister.ReturnedMap
}

func TestAllMethodCalled(t *testing.T) {
  fakeController := new(testLister)

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  if fakeController.TimesCalled != 1 {
    t.Errorf("Expected 1 call, got %+v.", fakeController.TimesCalled)
  }
}

func TestJsonResponseAfterReturningEmptyMapFromAll(t *testing.T) {
  fakeController := new(testLister)

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  expectedJson, _ := json.Marshal(fakeController.All())
  actualJson := w.Body.Bytes()
  if !bytes.Equal(expectedJson, actualJson) {
    t.Errorf("Expected response %+v, got %+v.", expectedJson, actualJson)
  }
}

func TestJsonResponseAfterReturningEmptyMapWithOneString(t *testing.T) {
  fakeController := new(testLister)
  fakeController.ReturnedMap = map[int64]interface{} { 0 : "test" }

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  expectedJson, _ := json.Marshal(fakeController.All())
  actualJson := w.Body.Bytes()
  if !bytes.Equal(expectedJson, actualJson) {
    t.Errorf("Expected response %+v, got %+v.", expectedJson, actualJson)
  }
}

type testStruct struct {
  test int
}

func TestJsonResponseAfterReturningEmptyMapWithTwoStructs(t *testing.T) {
  fakeController := new(testLister)
  fakeController.ReturnedMap = map[int64]interface{} {
    0 : testStruct{0}, 1 : testStruct{1},
  }

  router := NewRouter()
  router.AddResource("/test", fakeController)

  url := "http://www.domain.com/test/"
  req, _ := http.NewRequest("GET", url, nil)

  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  expectedJson, _ := json.Marshal(fakeController.All())
  actualJson := w.Body.Bytes()
  if !bytes.Equal(expectedJson, actualJson) {
    t.Errorf("Expected response %+v, got %+v.", expectedJson, actualJson)
  }
}

