package helpers

import (
	"fmt"
	"net"

	"github.com/robbyklein/trivia-backend/protobuf"
	"google.golang.org/protobuf/proto"
)

func RespondWithError(udp *net.UDPConn, clientAddress *net.UDPAddr, msg string) ([]byte, error) {
	// Create an error response
	response := &protobuf.Response{
		Status: protobuf.ResponseStatus_RESPONSE_ERROR,
		Text:   msg,
	}

	// Serialize the error response
	data, err := proto.Marshal(response)

	if err != nil {
		fmt.Println("Failed to serialize error response:", err)
		return nil, err
	}

	return data, nil
}
