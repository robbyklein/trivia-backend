package controllers

import (
	"fmt"
	"net"

	"github.com/robbyklein/trivia-backend/helpers"
	"github.com/robbyklein/trivia-backend/models"
	"github.com/robbyklein/trivia-backend/protobuf"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

func RegistrationsCreate(db *gorm.DB, udp *net.UDPConn, clientAddress *net.UDPAddr, msg *protobuf.RegisterMessage) {
	// Validate password
	if !helpers.IsPasswordValid(msg.Password) {
		helpers.RespondWithError(udp, clientAddress, "Password must be 8+ characters")
		return
	}

	// Create password hash
	hash, err := helpers.HashPassword(msg.Password)

	if err != nil {
		helpers.RespondWithError(udp, clientAddress, "Failed to has password")
		return
	}

	// Create a user
	user := models.User{
		Email:        msg.Email,
		Username:     msg.Username,
		PasswordHash: hash,
	}

	result := db.Create(&user)

	if result.Error != nil {
		helpers.RespondWithError(udp, clientAddress, "Username or email already taken")
		return
	}

	// Create a response
	response := &protobuf.Response{
		Status: protobuf.ResponseStatus_RESPONSE_SUCCESS,
		Text:   "Registration success! Please confirm your email",
	}

	// Serialize it
	data, err := proto.Marshal(response)

	if err != nil {
		fmt.Println("Failed to serialize response")
		return
	}

	helpers.RespondToClient(udp, clientAddress, data)
}
