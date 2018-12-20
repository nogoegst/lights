package main

import (
	"log"
	"net/http"

	"github.com/nogoegst/lights"
	"github.com/nogoegst/lights/frontend"
	"periph.io/x/periph/host/rpi"
)

func main() {
	if err := lights.Init(); err != nil {
		log.Fatal(err)
	}
	router := lights.NewRouter()

	lamp := lights.NewLight("lamp", rpi.P1_33)
	router.Add(lamp)

	cuties := lights.NewLight("cuties", rpi.P1_35)
	router.Add(cuties)

	http.Handle("/lights/", http.StripPrefix("/lights", router))

	frontend, err := frontend.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", frontend)

	if err := http.ListenAndServe(":http", nil); err != nil {
		log.Fatal(err)
	}
}
