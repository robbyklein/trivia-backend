package main

import (
	"fmt"
	"net"
	"os"

	"github.com/robbyklein/trivia-backend/controllers"
	"github.com/robbyklein/trivia-backend/initializers"
	"github.com/robbyklein/trivia-backend/models"
	"github.com/robbyklein/trivia-backend/protobuf"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

var db *gorm.DB
var udp *net.UDPConn

func init() {
	var err error

	// Connect to databse
	db, err = initializers.ConnectToDatabase()

	if err != nil {
		panic("Failed to connect to database")
	}

	// Drop tables
	if os.Getenv("GO_ENV") == "development" {
		db.Migrator().DropTable(models.User{})
	}

	// Create tables
	initializers.SyncDatabase(db)

	// Create udp server
	udp, err = initializers.CreateUDPServer()

	if err != nil {
		panic("Failed to create udp server")
	}
}

func main() {
	buffer := make([]byte, 1024)

	for {
		n, clientAddress, err := udp.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		handleMessage(buffer[:n], clientAddress)
	}
}

func handleMessage(data []byte, clientAddress *net.UDPAddr) {
	// Parse message
	var msg protobuf.TriviaMessage
	err := proto.Unmarshal(data, &msg)

	if err != nil {
		fmt.Println(err)
	}

	// Handle it
	switch msg.Type {
	case protobuf.MessageType_MESSAGE_REGISTER:
		controllers.RegistrationsCreate(db, udp, clientAddress, msg.Register)
	case protobuf.MessageType_MESSAGE_SIGN_IN:
		controllers.SignInCreate(db, udp, clientAddress, msg.SignIn)
	case protobuf.MessageType_MESSAGE_JOIN:
		fmt.Println("We hit a join message")
	case protobuf.MessageType_MESSAGE_ANSWER:
		fmt.Println("We hit a answer message")
	default:
		fmt.Println("Unknown message type")
	}
}
