package helpers

import (
	"fmt"
	"net"
)

func RespondToClient(udp *net.UDPConn, clientAddress *net.UDPAddr) {
	// Prepare a response message
	responseMessage := "Hello, Client!"

	// Send the response back to the client\
	var err error
	_, err = udp.WriteToUDP([]byte(responseMessage), clientAddress)

	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}
}
