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

func AnswerSubmit(db *gorm.DB, udp *net.UDPConn, clientAddress *net.UDPAddr, msg *protobuf.AnswerMessage) {
	// Decode token
	claims, err := helpers.DecodeJWTToken(msg.Token)

	if err != nil {
		helpers.RespondWithError(udp, clientAddress, "Invalid auth token")
		return
	}

	// Find the user
	var user models.User
	if result := db.First(&user, claims["sub"]); result.Error != nil {
		helpers.RespondWithError(udp, clientAddress, "Failed to find user")
		return
	}

	// Create a response
	response := &protobuf.Response{
		Status: protobuf.ResponseStatus_RESPONSE_SUCCESS,
		Text:   "Correct!",
	}

	// Serialize it
	data, err := proto.Marshal(response)

	if err != nil {
		fmt.Println("Failed to serialize response")
		return
	}

	helpers.RespondToClient(udp, clientAddress, data)
}
