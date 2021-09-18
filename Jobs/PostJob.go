package jobs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
)

func PostJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("######### Post JOB STARTED ###########")
	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	if claim.Role != "Recruiter" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var job MyDb.Job
	err = json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	}
	job.User_id = claim.User_ID
	job.Company_id = claim.Company_id
	err = MyDb.PostJob(job)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	}
	fmt.Println("JOB ADDED SUCCESSFULLY")
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(job)
	if err != nil {
		log.Fatalln(err)
	}

}
