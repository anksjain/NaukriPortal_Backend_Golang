package jobs

import (
	"encoding/json"
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
	"strconv"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("######### GET JOBS STARTED ###########")
	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var jobs []MyDb.Job
	if claim.Role == "Candidate" {
		var applied bool = false
		fmt.Println("In Candidate")
		for key, val := range r.Form {
			fmt.Println("My key ->", key, "and Values ->", val)
			if key == "type" && val[0] == "applied" {
				applied = true
				break
			}
		}
		jobs, err = MyDb.GetJobsByCandidate(claim.User_ID, applied)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application-json")
		err = json.NewEncoder(w).Encode(jobs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else if claim.Role == "Recruiter" {
		fmt.Println("In Recruiter")
		jobs, err = MyDb.GetJobsByRecruiter(claim.User_ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application-json")
		err = json.NewEncoder(w).Encode(jobs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else if claim.Role == "Admin" {
		fmt.Println("In Admin")
		for key, val := range r.Form {
			fmt.Println("My key ->", key, "and Values ->", val)
			if key == "applied" {
				user_id, err := strconv.Atoi(val[0])
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				jobs, err = MyDb.GetJobsByCandidate(user_id, true)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-type", "application-json")
				err = json.NewEncoder(w).Encode(jobs)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
				return
			}

		}
		jobs, err = MyDb.GetJobs()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application-json")
		err = json.NewEncoder(w).Encode(jobs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusBadGateway)

}
