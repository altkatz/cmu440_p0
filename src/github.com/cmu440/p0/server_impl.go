// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
)

type keyValueServer struct {
	// TODO: implement this!
}

var pServer *keyValueServer

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	pServer = &keyValueServer{}
	init_db()
	return pServer
}

func (kvs *keyValueServer) Start(port int) error {
	// TODO: implement this!
	var (
		err      error
		listener net.Listener
	)
	listener, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err == nil {
		fmt.Println("before serve", listener)
		go serve(listener)
	} else {
		fmt.Println("start error", err)
	}
	return err
}
func serve(listener net.Listener) {
	defer listener.Close()
	fmt.Println("in listener:", listener)
	for {
		fmt.Println("before accept")
		conn, err := listener.Accept()
		fmt.Println("after_accept")
		if err == nil {
			go serveConn(conn)
		}
	}
}

var getCmd = []byte("get")
var putCmd = []byte("put")

func serveConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Bytes()
		cmd := line[0:3]
		fmt.Println("command_is:", cmd)
		if bytes.Equal(cmd, getCmd) {
		} else if bytes.Equal(cmd, putCmd) {
		}
	}
}

func (kvs *keyValueServer) Close() {
	// TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
	// TODO: implement this!
	return 0
}

// TODO: add additional methods/functions below!
