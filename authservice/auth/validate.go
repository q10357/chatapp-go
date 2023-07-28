package auth

import (
	"fmt"
	"net/http"
)

// adds user to db
func ValidateHandler(rw http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	fmt.Printf("Authorization Header: %s\n", authHeader)
	fmt.Printf("This is the header: %s\n", authHeader)
}
