package initializers

import (
	"net"

	"github.com/robbyklein/trivia-backend/helpers"
)

func CreateUDPServer() (*net.UDPConn, error) {
	address := helpers.BuildAddress()
	udpAddress, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		return nil, err
	}

	// Create a UDP socket
	udp, err := net.ListenUDP("udp", udpAddress)

	if err != nil {
		return nil, err
	}

	return udp, nil
}
