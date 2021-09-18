package companies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	Auth "restapi/package/Auth"

	MyDb "restapi/package/Database"

	"github.com/gorilla/mux"
)

func GetCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########GET Compnay  STARTED #############")

	_, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-type", "application/json")
	var company_id int
	params := mux.Vars(r)
	company_id, err = strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something wrong with Id"))
		return
	}
	company, err := MyDb.GetCompany(company_id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if company.Id != company_id {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Company not exist"))
		return
	}
	json.NewEncoder(w).Encode(company)
}
