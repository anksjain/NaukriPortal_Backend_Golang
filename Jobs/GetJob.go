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

func GetJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########GET SPECIFIC JOB STARTED #############")

	_, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-type", "application/json")
	var job_id int
	params := mux.Vars(r)
	job_id, err = strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something wrong with Id"))
		return
	}
	job, err := MyDb.GetJob(job_id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if job.Id != job_id {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Job not exist"))
		return
	}
	json.NewEncoder(w).Encode(job)
}
