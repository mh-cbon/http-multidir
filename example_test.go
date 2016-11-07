package httpmultidir_test

import (
  "net/http"
  "github.com/mh-cbon/http-multidir"
)

func Example() {
  dirs := []string{"static/", "other_static/", "upload/",}
  fileServer := http.FileServer(httpmultidir.Multidir(dirs))
  http.Handle("/prefixed/", http.StripPrefix("/prefixed/", fileServer))
}
