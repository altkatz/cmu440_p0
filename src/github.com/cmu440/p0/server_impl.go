// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

var countChannel chan bool = make(chan bool)
var clientCount int = 0

type keyValueServer struct {
	// TODO: implement this!
	listener net.Listener
}

var pServer *keyValueServer

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	pServer = &keyValueServer{}
	init_db()
	return pServer
}

func incrCount() {
	for {
		incr := <-countChannel
		if incr {
			clientCount += 1
		} else {
			clientCount -= 1
		}
	}
}

func (kvs *keyValueServer) Start(port int) error {
	// TODO: implement this!
	var err error
	kvs.listener, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	go incrCount()
	if err == nil {
		go accpetor(kvs)
	} else {
		fmt.Println("start error", err)
	}
	return err
}
func accpetor(kvs *keyValueServer) {
	listener := kvs.listener
	for {
		conn, err := listener.Accept()
		if err == nil {
			countChannel <- true
			go serveConn(conn)
		}
	}
}

var getCmd = []byte("get")
var putCmd = []byte("put")

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func serveConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Bytes()
		cmd := line[0:3]
		if testEq(cmd, getCmd) {
			fmt.Println("is get")
		} else if testEq(cmd, putCmd) {
			fmt.Println("is put")
		}
	}
	countChannel <- false
	conn.Close()
}

func (kvs *keyValueServer) Close() {
	// TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
	return clientCount
}

// TODO: add additional methods/functions below!
