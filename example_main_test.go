package httpmultidir_test

import (
  "net/http"
  "github.com/mh-cbon/http-multidir"
)

var port = ":8080"

// Example_main demonstrates usage of multidir package.
func Example_main() {
  dirs := []string{"static/", "other_static/", "upload/",}
  fileServer := http.FileServer(httpmultidir.Multidir(dirs))
  http.Handle("/prefixed/", http.StripPrefix("/prefixed/", fileServer))
  http.ListenAndServe(port, nil)
}
