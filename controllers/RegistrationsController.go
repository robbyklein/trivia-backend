package controllers

import (
	"fmt"
	"net"

	"github.com/robbyklein/trivia-backend/helpers"
	"github.com/robbyklein/trivia-backend/models"
	"github.com/robbyklein/trivia-backend/protobuf"
	"gorm.io/gorm"
)

func RegistrationsCreate(db *gorm.DB, udp *net.UDPConn, clientAddress *net.UDPAddr, msg *protobuf.RegisterMessage) {
	// Validate password
	if !helpers.IsPasswordValid(msg.Password) {
		fmt.Println("Invalid password respond")
	}

	// Create password hash
	hash, err := helpers.HashPassword(msg.Password)

	if err != nil {
		fmt.Println("Failed to hash password respond")
	}

	// Create a user
	user := models.User{
		Email:        msg.Email,
		Username:     msg.Username,
		PasswordHash: hash,
	}

	result := db.Create(&user)

	if result.Error != nil {
		fmt.Println("Failed to create password respond")
	}

	helpers.RespondToClient(udp, clientAddress)
}
