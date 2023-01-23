package handler

import (
	"io"
	"net"
	"rocks_agent/response"
	"rocks_agent/stats"
)

func HandleMetric(metric string, err error, conn net.Conn) {
	switch err {
	case nil:
		switch metric {
		case "":
			response.Send(response.ResponseStruct{Code: response.FLD702, Message: "Undefined Metric Parameter"}, conn)
		case "cpu.load":
			response.Send(response.ResponseStruct{Code: response.FLD703, Message: "Feature is currently not available"}, conn)
		case "disk.iops":
			response.Send(response.ResponseStruct{Code: response.FLD703, Message: "Feature is currently not available"}, conn)
		case "disk.free":
			response.Send(stats.GetDiskSpace(), conn)
		default:



	}
	case io.EOF:
		response.Send(response.ResponseStruct{Code: response.FLD701, Message: "Client closed connection unexpected"}, conn)
	}

}