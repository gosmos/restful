restful
=======

A library that simplifies routing configuration of RESTful resources
by reducing it to one line of code for each resource. Concept is inspired
by [REST resource routing configuration in Ruby on Rails framework](
http://guides.rubyonrails.org/routing.html#resource-routing-the-rails-default)
and [Model-View-Controller](
http://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller)
architectural pattern.

Fetch gosmos/restful library to your go workspace!

```bash
go get github.com/gosmos/restful
```

Hello, World!
-------------

Following example shows a way of using the library to create REST service
that lists resources managed by the system. Code is heavily commented
to facilitate fast understanding of what's going on.

```go
package main

// We will need only gosmos/restful package and some http server.
import (
  "github.com/gosmos/restful"
  "net/http"
)

// First thing to do, is to define controller structure, which will be
// responsible for handling calls to RESTful API.
type MyToysController struct {
  // Memory storage implemented on map is used for simplicity.
  // Normally we would want some persistence in here.
  storage map[string]interface{}
}
// Contructor function.
func NewMyToysController() *MyToysController {
  controller := &MyToysController { make(map[string]interface{}) }
  return controller
}

// Our controller needs to implement at least one interface for handling
// REST calls (check controller.go for all possible interfaces).
// Interface restful.Indexer will be our choice for this example.

// Index() method will be called by the library on each access to root
// of our RESTful resource ("/") that uses GET method (see HTTP/1.1 spec).
// It must return a map of type map[string]interface{}. Returned map
// will be encoded to JSON and written into the body of HTTP response.
func (controller *MyToysController) Index() map[string]interface{} {
  return controller.storage
}

// It's good to check if proper interface is implemented by the controller,
// so we will face early crash in case of method signature error.
var _ restful.Indexer = &MyToysController{}

// Lastly, we need to create router, configure our RESTful resource,
// and start HTTP server. Voilà!
func main() {
  router := restful.NewRouter()
  router.HandleResource("/api/mytoys", NewMyToysController())
  http.ListenAndServe(":8080", router)
}
```

Compiling and running code above will result in fully operational RESTful
resource. If you're using bash just paste following code into the terminal.

```bash
wget http://raw.github.com/gosmos/restful/master/example/hello/main.go
go build main.go
./main
```

Empty program output signifies that everything went well and HTTP server
is running. You should be able to see empty JSON object after accessing
[localhost:8080/api/mytoys/](http://localhost:8080/api/mytoys/).

Object is empty because map returned from *MyToysController.Index()*
is empty. Try addding some code in controller's constructor
that inserts content to *controler.storage* and see if your data
is being encoded to [JSON](http://json.org/) properly.
To see your changes in the browser you need to recompile and restart
example program.

Have fun! :sweat_smile:

Google App Engine
-----------------

Library is fully GAE-compatible. There are few things to pay attention to:
 1. Invoking *http.ListenAndServe* is not allowed on GAE.
    Use *http.Handle()* or *http.HandleFunc* instead.
 2. Using package *main* and declaring *main()* function is prohibited.
    Create router in *init()* function of another package
    or as a global variable.

Consult [official GAE docs](
https://developers.google.com/appengine/docs/go/gettingstarted/helloworld)
for details.

Why lack of ListenAndServe method?
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

