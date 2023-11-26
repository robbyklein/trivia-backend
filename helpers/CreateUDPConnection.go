package helpers

import (
	"net"
)

func CreateUDPConnection() (*net.UDPConn, error) {
	// Get address
	address, err := net.ResolveUDPAddr("udp", BuildAddress())

	if err != nil {
		return nil, err
	}

	// Create connection
	connection, err := net.DialUDP("udp", nil, address)

	if err != nil {
		return nil, err
	}

	return connection, nil
}
