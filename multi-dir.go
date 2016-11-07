package httpmultidir

import (
	"errors"
	http "net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// A Multidir implements FileSystem using the native file system restricted to a
// set of specific directory tree to treat in order.
//
// An empty Multidir always returns not found.
// An empty value in Multidir is always ignored.
type Multidir []string

// NotFound is an error which represents a not found resource, it is likely to return a 404.
var NotFound = os.ErrNotExist

// Open locate name within available directories.
// The first directory to contain a corresponding file,
// returns that resource.
// If the matched name is a directory, it returns NotFound error.
func (d Multidir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	var dirs []string
	dirs = d
	for _, dir := range dirs {
		if dir != "" {
			f, err := os.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
			if err != nil {
				continue
			}
			d, err := f.Stat()
			if d.IsDir() {
				continue
			}
			return f, nil
		}
	}
	return nil, NotFound
}
