package shuttle

import (
	"fmt"
	"goRecycleBuffer"
	"net"
)

type HandleRequest func(net.Conn, int, []byte)

var bufferSize = 8192
var rb = goRecycleBuffer.Init(bufferSize)

const workerCount = 4

func Start(host, port string, handler HandleRequest) {
	goRecycleBuffer.Init(bufferSize)

	addr := host + ":" + port
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	fmt.Println("Listening on " + addr)

	for i := 0; i < workerCount; i++ {
		go accept(l, handler)
	}

	done := make(chan bool)
	<-done
}

func accept(l net.Listener, handler HandleRequest) {
	for {
		conn, err := l.Accept()
		if err == nil {
			go worker(conn, handler)
		}
	}
}

func worker(conn net.Conn, handler HandleRequest) {
	var readBuffer []byte
	for {
		readBuffer = <-rb.Get
		readLen, err := conn.Read(readBuffer)

		if readLen > 0 && err == nil {
			handler(conn, readLen, readBuffer)
			rb.Give <- readBuffer
		} else {
			rb.Give <- readBuffer
			return
		}
	}
}
