package connection

import
(
	"net"
	"fmt"
)

const (
	LISTEN_ADDR = "0.0.0.0"
	LISTEN_PORT = "9999"
	LISTEN_TYPE = "tcp"
)

type TCPConnection struct {
	Conn net.Conn
}

func GetConnection() TCPConnection {
	var addr = fmt.Sprintf("%v:%v", LISTEN_ADDR, LISTEN_PORT)
	listener, _ := net.Listen(LISTEN_TYPE, addr)
	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		return TCPConnection{Conn: connection}
	}
}