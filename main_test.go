package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/robbyklein/trivia-backend/helpers"
	"github.com/robbyklein/trivia-backend/protobuf"
	"google.golang.org/protobuf/proto"
)

func startServerListener() {
	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := udp.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Failed to read message")
		}

		handleMessage(buffer[:n], clientAddr)
	}
}

func TestRegistration(t *testing.T) {
	// Start server message listener
	go startServerListener()

	// Create a connection to send messages
	connection, err := helpers.CreateUDPConnection()

	if err != nil {
		t.Fatalf("Failed to connect to UDP server")
	}

	// Create a register message
	registerMessage := &protobuf.TriviaMessage{
		Type: protobuf.MessageType_MESSAGE_REGISTER,
		Register: &protobuf.RegisterMessage{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		},
	}

	// Serialize it
	data, err := proto.Marshal(registerMessage)

	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	// Send it
	_, err = connection.Write(data)

	if err != nil {
		t.Fatalf("Failed to write to UDP server: %v", err)
	}

	// Add delay to allow server to process the message.
	time.Sleep(1 * time.Second)

	// Create buffer for response
	buffer := make([]byte, 1024)

	// Read it
	n, _, err := connection.ReadFromUDP(buffer)

	if err != nil {
		t.Fatalf("Failed to read response from UDP server: %v", err)
	}

	fmt.Println(string(buffer[:n]))

}

func TestLogin(t *testing.T) {
	// Start server message listener
	go startServerListener()

	// Create a connection to send messages
	connection, err := helpers.CreateUDPConnection()

	if err != nil {
		t.Fatalf("Failed to connect to UDP server")
	}

	// Create a register message
	signInMessage := &protobuf.TriviaMessage{
		Type: protobuf.MessageType_MESSAGE_SIGN_IN,
		SignIn: &protobuf.SignInMessage{
			Email:    "test@example.com",
			Password: "password123",
		},
	}

	// Serialize it
	data, err := proto.Marshal(signInMessage)

	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	// Send it
	_, err = connection.Write(data)

	if err != nil {
		t.Fatalf("Failed to write to UDP server: %v", err)
	}

	// Add delay to allow server to process the message.
	time.Sleep(1 * time.Second)
}
