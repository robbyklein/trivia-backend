package controllers

import (
	"net"

	"github.com/robbyklein/trivia-backend/game"
	"github.com/robbyklein/trivia-backend/helpers"
	"github.com/robbyklein/trivia-backend/protobuf"
	"gorm.io/gorm"
)

func JoinGame(db *gorm.DB, udp *net.UDPConn, clientAddress *net.UDPAddr, msg *protobuf.AnswerMessage) {
	// Decode token
	claims, err := helpers.DecodeJWTToken(msg.Token)

	if err != nil {
		helpers.RespondWithError(udp, clientAddress, "Invalid auth token")
		return
	}

	// Join game
	if ok := game.AddToGame(claims["userId"].(int)); !ok {
		helpers.RespondWithError(udp, clientAddress, "Game already full")
		return
	}

	// Alert all players

}
