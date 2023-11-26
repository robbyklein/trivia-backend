package helpers

import (
	"net"
)

func RespondToClient(udp *net.UDPConn, clientAddress *net.UDPAddr, response []byte) error {
	// Send the response back to the client\
	var err error
	_, err = udp.WriteToUDP(response, clientAddress)

	if err != nil {
		return err
	}

	return nil
}
