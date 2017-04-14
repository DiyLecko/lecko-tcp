package arbiter

import (
	"net"
	"shuttle"
)

func Start(host, port string) {
	shuttle.Start(host, port, handleRequest)
}

func handleRequest(conn net.Conn, readLen int, readBuffer []byte) {
	//	fmt.Println(string(readBuffer[:readLen]))
}
