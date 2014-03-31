package restful

type Lister interface {
  All() map[int64]interface{}
}
type Getter interface {
  Get(int64) interface{}
}
type Creator interface {
  New() interface{}
}
type Adder interface {
  Creator
  Add(interface{}) int64
}
type Replacer interface {
  Creator
  Replace(int64, interface{})
}
type Deleter interface {
  Delete(int64) bool
}

