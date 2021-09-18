package auth

import (
	"fmt"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, req *http.Request) {
	claim, err := VerifyJWTAndGetClaim(req)
	if err != nil {
		fmt.Println("ALREADY LOGOUT")
		http.Error(w, "PLEASE LOGIN FIRST", http.StatusBadRequest)
		return
	}
	claim.ExpiresAt = time.Now().Unix()
	fmt.Println("LOGOUT SUCCESSFULLY")
	fmt.Fprint(w, "LOGOUT SUCCESSFULLY")
}
