package helpers

import (
	"fmt"
	"os"
)

func BuildAddress() string {
	// Build address
	base := os.Getenv("ROOT_URL")
	if base == "" {
		base = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf("%s:%s", base, port)
}
