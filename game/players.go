package game

import (
	"net"
	"time"
)

type ActivePlayer struct {
	userId        int
	score         int
	lastUpdate    time.Time
	clientAddress *net.UDPAddr
}

var ActivePlayers map[int]ActivePlayer

func AddToGame(userId int, clientAddress *net.UDPAddr) bool {
	if len(ActivePlayers) > 11 || userId == 0 {
		return false
	}

	newPlayer := ActivePlayer{
		score:         0,
		userId:        userId,
		lastUpdate:    time.Now(),
		clientAddress: clientAddress,
	}

	if ActivePlayers == nil {
		ActivePlayers = make(map[int]ActivePlayer)
	}

	ActivePlayers[userId] = newPlayer

	return true
}

func RemoveFromGame(userId int) {
	delete(ActivePlayers, userId)
}
