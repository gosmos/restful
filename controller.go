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

/*
  Lists all objects contained in this REST resource.
*/
type Indexer interface {
  /*
    Invoked by the router when call to root of the resource ("/")
    is received via GET HTTP method. Returned map will be JSON-encoded
    and writen into HTTP response.
  */
  Index() map[string]interface{}
}

/*
  Returns object of given id contained in this REST resource.
*/
type Shower interface {
  /*
    Invoked by the router when class to resource object ("/{id}")
    is received via GET HTTP method. Returned object will be put in
    a response map under key of resource id. Map will be JSON-encoded
    and written into HTTP response.
  */
  Show(string) interface{}
}

/*
  Adds object to this REST resource.
*/
type Creator interface {
  /*
    Invoked by the router when call to root of the resource ("/")
    if received via POST HTTP method. Returned object will be used
    to decode JSON-encoded body of HTTP request into. After filling
    with data, object will be passed to Create() method.
  */
  New() interface{}
  /*
    Invoked by the router after decoding body of HTTP request into
    the object returned from New() method. Method is responsible
    for generating id for given object and returning it. Returned
    id will be used as a key of new object in a response map.
    Map will be JSON-encoded and written into HTTP response.
  */
  Create(interface{}) string
}

/*
  Replaces content of an object contained in this REST resource.
*/
type Updater interface {
  /*
    Invoked by the router when call to resource object ("/{id}")
    if received via PUT HTTP method. Returned object will be used
    to decode JSON-encoded body of HTTP request into. After filling
    with data, object will be passed to Update() method.
  */
  New() interface{}
  /*
    Invoked by the router after decoding body of HTTP request into
    the object returned from New() method. If this method finishes
    thithout panic, response map will contain true value under index
    "ok". Map will be JSON-encoded and written into HTTP response.
  */
  Update(string, interface{})
}

/*
  Deletes object from this REST resource.
*/
type Deleter interface {
  /*
    Invoked by the router when class to resource object ("/{id}")
    is received via DELETE HTTP method. If this method finishes
    thithout panic, response map will contain true value under index
    "ok". Map will be JSON-encoded and written into HTTP response.
  */
  Delete(string)
}

