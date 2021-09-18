package jobs

import (
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	"strconv"

	MyDb "restapi/package/Database"

	"github.com/gorilla/mux"
)

func ApplyJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("######### APPLY TO JOB STARTED ###########")
	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	if claim.Role != "Candidate" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var job_id int
	params := mux.Vars(r)
	job_id, err = strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(job_id)
	applied, err := MyDb.ApplyJob(job_id, claim.User_ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if applied <= 0 {
		w.Write([]byte("Invalid Job"))
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	fmt.Println("APPLIED TO JOB SUCCESSFULLY")
	fmt.Fprint(w, "APPLIED TO JOB SUCCESSFULLY")
}
