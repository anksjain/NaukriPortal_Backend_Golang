package jobs

import (
	"encoding/json"
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########UPDATE JOB STARTED #############")

	claims, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	if claims.Role != "Recruiter" {
		w.Write([]byte("Not a valid user"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var job MyDb.Job
	err = json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	job.User_id = claims.User_ID
	w.Header().Add("Content-type", "application/json")
	params := mux.Vars(r)
	job.Id, err = strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something wrong with Id"))
		return
	}
	updated, err := MyDb.UpdateJob(&job)
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
	w.Write([]byte("Updated"))
}
