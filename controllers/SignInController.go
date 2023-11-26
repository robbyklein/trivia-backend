package controllers

import (
	"fmt"
	"net"

	"github.com/robbyklein/trivia-backend/helpers"
	"github.com/robbyklein/trivia-backend/models"
	"github.com/robbyklein/trivia-backend/protobuf"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

func SignInCreate(db *gorm.DB, udp *net.UDPConn, clientAddress *net.UDPAddr, msg *protobuf.SignInMessage) {
	// Find the user
	var user models.User

	if result := db.First(&user, "email = ?", msg.Email); result.Error != nil {
		helpers.RespondWithError(udp, clientAddress, "Failed to find user")
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(msg.Password)); err != nil {
		helpers.RespondWithError(udp, clientAddress, "Incorrect password")
	}

	// Create token
	token := helpers.GenerateHMAC(int(user.ID))

	// Create a response
	response := &protobuf.Response{
		Status: protobuf.ResponseStatus_RESPONSE_SUCCESS,
		Text:   token,
	}

	// Serialize it
	data, err := proto.Marshal(response)

	if err != nil {
		fmt.Println("Failed to serialize response")
		return
	}

	helpers.RespondToClient(udp, clientAddress, data)
}
