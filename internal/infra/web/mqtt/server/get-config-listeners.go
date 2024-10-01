package server

import (
	"fmt"

	"github.com/mochi-mqtt/server/v2/listeners"
)

func GetListeners(host, id string, port int) listeners.Config {

	return listeners.Config{
		Address: fmt.Sprintf("%s:%d", host, port),
		ID:      id,
	}
}
