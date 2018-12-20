package lights

import (
	"net/http"

	"periph.io/x/periph/conn/gpio"
)

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
