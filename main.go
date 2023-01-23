package main

import (
	"bufio"
	"rocks_agent/connection"
	"rocks_agent/handler"
	"strings"
)

func main() {
	conn := connection.GetConnection().Conn
	clientReader := bufio.NewReader(conn)
	for {
		clientRequest, err := clientReader.ReadString('\n')
		clientRequest = strings.TrimSpace(clientRequest)
		handler.HandleMetric(clientRequest, err, conn)
	}

}
