// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"strconv"
)

type keyValueServer struct {
	// TODO: implement this!
	listener net.Listener
}

var pServer *keyValueServer

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	pServer = &keyValueServer{}
	return pServer
}

func (kvs *keyValueServer) Start(port int) error {
	// TODO: implement this!
	var err error
	kvs.listener, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	fmt.Println("start error", err)
	return err
}

func (kvs *keyValueServer) Close() {
	// TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
	// TODO: implement this!
	return -1
}

// TODO: add additional methods/functions below!
