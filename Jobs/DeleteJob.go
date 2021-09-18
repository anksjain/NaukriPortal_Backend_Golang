package jobs

import (
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########Delete JOB STARTED #############")

	claims, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	if claims.Role == "Candidate" {
		w.Write([]byte("Not a valid user"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-type", "application/json")
	params := mux.Vars(r)
	job_id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something wrong with Id"))
		return
	}
	if claims.Role == "Admin" {
		claims.User_ID = -1
	}
	updated, err := MyDb.DeleteJob(job_id, claims.User_ID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if updated <= 0 {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Not Valid Details"))
		return
	}
	w.Write([]byte("Deleted"))
}
