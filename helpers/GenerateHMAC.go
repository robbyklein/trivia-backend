package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
)

func GenerateHMAC(userId int) string {
	sharedSecret := os.Getenv("SHARED_SECRET")
	h := hmac.New(sha256.New, []byte(sharedSecret))

	userIdString := strconv.Itoa(userId)
	h.Write([]byte(userIdString))

	return hex.EncodeToString(h.Sum(nil))
}
