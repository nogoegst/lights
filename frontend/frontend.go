package frontend

import (
	"net/http"

	_ "github.com/nogoegst/lights/frontend/statik"
	"github.com/rakyll/statik/fs"
)

func New() (http.Handler, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}
	return http.FileServer(statikFS), nil
}
