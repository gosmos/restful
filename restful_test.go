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

