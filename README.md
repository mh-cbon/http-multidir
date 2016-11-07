# http-multidir

Serve multiple static directories under the same prefix.

Unlike `http.Dir`, `httpmultidir.Multidir` will not list directories content
and will serve only matching files.

# Install

```sh
go get github.com/mh-cbon/http-multidir
glide install github.com/mh-cbon/http-multidir
```

# Usage

```go
package main

import (
	"net/http"
	"github.com/mh-cbon/http-multidir"
)

func main() {
  dirs := []string{"static/", "other_static/", "upload/",}
  fileServer := http.FileServer(httpmultidir.Multidir(dirs))
	http.Handle("/prefixed/", http.StripPrefix("/prefixed/", fileServer))
}
```
