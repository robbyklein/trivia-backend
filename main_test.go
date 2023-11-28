package main

import (
	"testing"
	"time"

	"github.com/robbyklein/trivia-backend/helpers"
	"github.com/robbyklein/trivia-backend/protobuf"
	"google.golang.org/protobuf/proto"
)

var buffer = make([]byte, 1024)
var token string

func TestRegistration(t *testing.T) {

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
	connection, _ := helpers.CreateUDPConnection()
	_, err = connection.Write(data)

	if err != nil {
		t.Fatalf("Failed to write to UDP server: %v", err)
	}

	// Add delay to allow server to process the message.
	time.Sleep(1 * time.Second)

	// Read it
	n, _, err := connection.ReadFromUDP(buffer)

	if err != nil {
		t.Fatalf("Failed to read response from UDP server: %v", err)
	}

	// Deserialize it
	var msg protobuf.Response
	err = proto.Unmarshal(buffer[:n], &msg)

	if msg.Status != protobuf.ResponseStatus_RESPONSE_SUCCESS {
		t.Fatalf("Recieved a non success status: %v", err)
	}
}

func TestSignin(t *testing.T) {
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
	connection, _ := helpers.CreateUDPConnection()
	_, err = connection.Write(data)

	if err != nil {
		t.Fatalf("Failed to write to UDP server: %v", err)
	}

	// Add delay to allow server to process the message.
	time.Sleep(1 * time.Second)

	// Read it
	n, _, err := connection.ReadFromUDP(buffer)

	if err != nil {
		t.Fatalf("Failed to read response from UDP server: %v", err)
	}

	// Deserialize it
	var msg protobuf.Response
	err = proto.Unmarshal(buffer[:n], &msg)

	if msg.Status != protobuf.ResponseStatus_RESPONSE_SUCCESS {
		t.Fatalf("Recieved a non success status: %v", err)
	}

	token = msg.Text
}

func TestAnswer(t *testing.T) {
	// Create a register message
	answerMessage := &protobuf.TriviaMessage{
		Type: protobuf.MessageType_MESSAGE_ANSWER,
		Answer: &protobuf.AnswerMessage{
			Answer: "The beatles",
			Token:  token,
		},
	}

	// Serialize it
	data, err := proto.Marshal(answerMessage)

	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	// Send it
	connection, _ := helpers.CreateUDPConnection()
	_, err = connection.Write(data)

	if err != nil {
		t.Fatalf("Failed to write to UDP server: %v", err)
	}

	// Add delay to allow server to process the message.
	time.Sleep(1 * time.Second)

	// Read it
	n, _, err := connection.ReadFromUDP(buffer)

	if err != nil {
		t.Fatalf("Failed to read response from UDP server: %v", err)
	}

	// Deserialize it
	var msg protobuf.Response
	err = proto.Unmarshal(buffer[:n], &msg)

	if msg.Status != protobuf.ResponseStatus_RESPONSE_SUCCESS {
		t.Fatalf("Recieved a non success status: %v", err)
	}

	token = msg.Text
}
