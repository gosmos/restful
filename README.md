restful
=======

Concept of gosmos/restful is inspired by
[REST resource routing configuration in Ruby on Rails framework](
http://guides.rubyonrails.org/routing.html#resource-routing-the-rails-default).
It aims to simplify RESTful resource routing configuration,
by reducing it to one line of code for each resource.

Fetch gosmos/restful library to your go workspace!

```bash
go get github.com/gosmos/restful
```

Show me the Code!
-----------------

```go
package main

import (
  "github.com/gosmos/restful"
  "net/http"
)

func main()
  router := restful.NewRouter()
  router.AddResource("/api/mytoys", NewMyToysController())
  http.ListenAndServe(":8080", router)
}
```

Creating RESTful resource with gosmos/restful library
can be divided into three simple steps:
 1. create router,
 2. configure resource,
 3. and set the router as HTTP handler.

The last thing to do, to make it work, is to define controller structure
and at least one method for handling REST calls.

```go
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

// Contructs the controller.
func NewMyToysController() *MyToysController {
  return &MyToysController { make(map[string]interface{}), 0 }
}
```

Compiling and running code above will result in fully operational
RESTful resource. Going to http://localhost:8080/api/mytoys/
should display empty json object.

Convention Over Configuration
-----------------------------

TODO

Google App Engine
-----------------

Library is fully GAE-compatible. There are two things to pay attention to:
 1. Invoking *http.ListenAndServe()* is not allowed on GAE.
 2. Using package *main* is prohibited.

Consult [official GAE docs](
https://developers.google.com/appengine/docs/go/gettingstarted/helloworld)
for details.

Why lack of ListenAndServeHTTP method?
--------------------------------------

Current implementation is based on
[gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) library,
so implementing this method would cost just 3 lines of code,
but we didn't want to add functionality just because it's easy
or it (kind of) fits in here. Implementing this method
would create another dependency to mux library,
which we may want to remove in the future.
Restful is a HTTP router, not a HTTP server.
Server functionality is (pretty well) covered
in [net/http](http://golang.org/pkg/net/http/) package.

