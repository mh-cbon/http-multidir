package main_test

import (
  "net/http"
  "github.com/mh-cbon/http-multidir"
)

func main() {
  dirs := []string{"static/", "other_static/", "upload/",}
  fileServer := http.FileServer(httpmultidir.Multidir(dirs))
  http.Handle("/prefixed/", http.StripPrefix("/prefixed/", fileServer))
}
