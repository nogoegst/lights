package lights

import (
	"net/http"

	"github.com/bmizerany/pat"
)

type Router struct {
	lights map[string]*Light
	m      *pat.PatternServeMux
}

func NewRouter() *Router {
	return &Router{
		lights: make(map[string]*Light),
		m:      pat.New(),
	}
}

func (r *Router) Add(l *Light) {
	r.lights[l.Name()] = l
	r.m.Get("/"+l.Name()+"/:action", l)
}

func (r *Router) List() []*Light {
	var ret []*Light
	for _, light := range r.lights {
		ret = append(ret, light)
	}
	return ret
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.m.ServeHTTP(w, req)
}
