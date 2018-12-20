package lights

import (
	"encoding/json"
	"net/http"

	"github.com/bmizerany/pat"
)

type Router struct {
	lights map[string]*Light
	m      *pat.PatternServeMux
}

func NewRouter() *Router {
	r := &Router{
		lights: make(map[string]*Light),
		m:      pat.New(),
	}
	r.m.Get("/", http.HandlerFunc(r.listLightsHandler))
	return r
}

func (r *Router) Add(l *Light) {
	r.lights[l.Name()] = l
	r.m.Get("/"+l.Name()+"/:action", l)
}

func (r *Router) listLightsHandler(w http.ResponseWriter, req *http.Request) {
	var ret []string
	for _, light := range r.lights {
		ret = append(ret, light.Name())
	}
	json.NewEncoder(w).Encode(ret)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.m.ServeHTTP(w, req)
}
