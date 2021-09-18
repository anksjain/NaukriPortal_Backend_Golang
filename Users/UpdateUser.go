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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update USER METHOD STARTED")
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
	if user.Id != user_id {
		w.Write([]byte("User Not Found"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var updated MyDb.User
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updated.Id = user.Id
	if updated.Password == "" {
		updated.Password = user.Password
	}
	if updated.First_Name == "" {
		updated.First_Name = user.First_Name
	}
	if updated.Last_Name == "" {
		updated.Last_Name = user.Last_Name
	}
	if updated.City == "" {
		updated.City = user.City
	}
	if updated.Experience == 0 {
		updated.Experience = user.Experience
	}
	_, err = MyDb.UpdateUser(updated)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated Successfully"))
}
