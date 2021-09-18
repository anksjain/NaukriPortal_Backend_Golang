package users

import (
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########DELETE USER STARTED #############")

	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	if claim.Role != "Admin" {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-type", "application/json")
	params := mux.Vars(r)
	user_id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something wrong with Id"))
		return
	}
	deleted, err := MyDb.DeleteUser(user_id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if deleted <= 0 {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Not Valid Details"))
		return
	}
	w.Write([]byte("Deleted"))
}
