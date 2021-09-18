package companies

import (
	"encoding/json"
	"fmt"
	"net/http"

	Auth "restapi/package/Auth"

	MyDb "restapi/package/Database"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########Post COMAPANY STARTED #############")

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
	var company MyDb.Company
	err = json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = MyDb.PostCompany(company)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println("Company ADDED SUCCESSFULLY")
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
