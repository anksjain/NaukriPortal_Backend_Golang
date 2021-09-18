package companies

import (
	"encoding/json"
	"fmt"
	"net/http"
	MyDb "restapi/package/Database"
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("######### GET Companies STARTED ###########")
	// claim, err := Auth.VerifyJWTAndGetClaim(r)
	// if err != nil {
	// 	http.Error(w, "No Access", http.StatusUnauthorized)
	// 	return
	// }
	// if claim.Role == "Candidate" {
	// 	http.Error(w, "No Access", http.StatusUnauthorized)
	// 	return
	// }
	var companies []MyDb.Company
	companies, err := MyDb.GetAllCompany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application-json")
	err = json.NewEncoder(w).Encode(companies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
