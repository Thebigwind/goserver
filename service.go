package webFile

import (
	"fmt"

	. "goserver/server"
)

const (
	DEFAULT_REST_SERV string = "127.0.0.1"
	DEFAULT_REST_PORT string = "9090"
)

func ServerStart() error {
	restAddr := fmt.Sprintf(":%s", DEFAULT_REST_PORT)
	serv := NewRESTServer(restAddr)
	serv.StartRESTServer()
	return nil
}
