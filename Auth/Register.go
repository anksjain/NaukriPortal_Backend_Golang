package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	MyDb "restapi/package/Database"

	jwt "github.com/dgrijalva/jwt-go"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE USER METHOD STARTED")

	_, err := VerifyJWTAndGetClaim(r)
	// if no error
	if err == nil {
		fmt.Println("Already Authenticated, USER ID =")
		http.Error(w, "PLEASE LOGOUT FIRST", http.StatusForbidden)
		return
	}
	// if all other error
	if err != http.ErrContentLength && err != jwt.ErrSignatureInvalid {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user MyDb.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exist, err := MyDb.IsUsernameExist(user.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exist {
		fmt.Println("Username already exist")
		http.Error(w, "Username already exist", http.StatusFound)
		return
	}
	eror := MyDb.CreateUser(user)
	if eror != nil {
		fmt.Println(eror)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("USER CREATED SUCCESFULLY")
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode("User Successfully created")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
