package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET USER METHOD STARTED")
	if Auth.CorsMiddleware(w, r) {
		return
	}
	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	if claim.Role != "Admin" {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	user_id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, eror := MyDb.GetUser(user_id)
	if eror != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Id != 0 {
		json.NewEncoder(w).Encode(user)
		return
	}
	http.Error(w, "USER NOT FOUND", http.StatusBadRequest)
}
