package lights

import (
	"net/http"

	"github.com/bmizerany/pat"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
)

func Init() error {
	_, err := host.Init()
	return err
}

type Light struct {
	name string
	pin  gpio.PinIO
}

func NewLight(name string, pin gpio.PinIO) *Light {
	return &Light{
		name: name,
		pin:  pin,
	}
}

func (l *Light) Name() string {
	return l.name
}

func (l *Light) ToggleOn() error {
	return l.pin.Out(gpio.High)
}

func (l *Light) ToggleOff() error {
	return l.pin.Out(gpio.Low)
}

func (l *Light) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get(":action")
	switch action {
	case "on":
		if err := l.ToggleOn(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	case "off":
		if err := l.ToggleOff(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "no such an action", http.StatusBadRequest)
}

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
