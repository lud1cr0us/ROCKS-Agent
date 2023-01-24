package main

import (
	"fmt"
	"rocks_agent/error"
	"net"
	"rocks_agent/handler"
)

const (
	MASTER_ADDR = "localhost"
	MASTER_PORT = "7878"
	MASTER_TYPE = "udp"
)

 func handleUDPConnection(conn *net.UDPConn) {

         buffer := make([]byte, 1024)

				 // Reading from UDP Connection into buffer
         n, addr, err := conn.ReadFromUDP(buffer);
				 ROCKSError.Encounter(ROCKSError.ErrorInformation{ErrorType: "FATAL", Error: err})

				 // handling metric given from master
				 metricResponse := handler.HandleMetric(string(buffer[:n]))
         message := []byte(metricResponse)

				 // Respond to Master (addr) with the Metric Response
         _, err = conn.WriteToUDP(message, addr);
				 ROCKSError.Encounter(ROCKSError.ErrorInformation{ErrorType: "FATAL", Error: err})
 }

 func main() {
         service := MASTER_ADDR + ":" + MASTER_PORT

				 // Resolving UDP Address of Client + Error Handling
         udpAddr, err := net.ResolveUDPAddr(MASTER_TYPE, service);
				 ROCKSError.Encounter(ROCKSError.ErrorInformation{ErrorType: "FATAL", Error: err})

				 // Listening on UDP for Metrics
         ln, err := net.ListenUDP(MASTER_TYPE, udpAddr);
				 ROCKSError.Encounter(ROCKSError.ErrorInformation{ErrorType: "FATAL", Error: err})

				 // Message when Listening
         fmt.Printf("ROCKS Agent is ready for Recieving Metric Requests on %v\n", MASTER_PORT)
         defer ln.Close()

         for {
                 handleUDPConnection(ln)
         }

 }

