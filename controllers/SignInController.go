package controllers

import (
	"net"

	"github.com/robbyklein/trivia-backend/protobuf"
	"gorm.io/gorm"
)

func SignInCreate(db *gorm.DB, udp *net.UDPConn, clientAddress *net.UDPAddr, msg *protobuf.SignInMessage) {

}
