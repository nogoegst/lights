package lights

import (
	"periph.io/x/periph/host"
)

func Init() error {
	_, err := host.Init()
	return err
}
