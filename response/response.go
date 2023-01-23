package response

import (
	"fmt"
	"net"
)

const (
	SCS600 = "ROCKS.SCS600: "
	SCS601 = "ROCKS.SCS601: "
	SCS602 = "ROCKS.SCS602: "
	SCS603 = "ROCKS.SCS603: "
	SCS604 = "ROCKS.SCS604: "
)

const (
	FLD700 = "ROCKS.FLD700: "
	FLD701 = "ROCKS.FLD701: "
	FLD702 = "ROCKS.FLD702: "
	FLD703 = "ROCKS.FLD703: "
)

type ResponseStruct struct {
	Code string
	Message string
}

func Send(response ResponseStruct, conn net.Conn) {
	_, err := conn.Write([]byte(fmt.Sprintf("%v%v\n", response.Code, response.Message)))
	if err != nil {
		fmt.Print(err)
	}
}