package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	MyDb "restapi/package/Database"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET LOGIN METHOD STARTED")
	//checking is already login or not
	_, err := VerifyJWTAndGetClaim(r)
	if err == nil {
		fmt.Println("Already Authenticated, USER ID =")
		http.Error(w, "PLEASE LOGOUT FIRST", http.StatusForbidden)
		return
	}
	// decoding credentials
	var credentials Credentials
	fmt.Println(r.Body)
	error := json.NewDecoder(r.Body).Decode(&credentials)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//checking credentials in database
	user, exist, err := MyDb.VerfiyCredentials(credentials.Email, credentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if exist {
		token, err := getJWT(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("LOGIN SUCCESSFULLY")
		w.Header().Set("Content-type", "application/json")
		w.Write(js)
		return
	}
	// if credentials not valid
	fmt.Println("INVALID LOGIN DETAILS")
	w.Header().Set("Content-type", "application/json")
	http.Error(w, "INVALID CREDENTIALS", http.StatusBadRequest)

}
