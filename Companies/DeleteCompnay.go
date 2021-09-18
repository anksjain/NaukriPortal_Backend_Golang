package companies

import (
	"fmt"
	"net/http"
	"strconv"

	Auth "restapi/package/Auth"

	MyDb "restapi/package/Database"

	"github.com/gorilla/mux"
)

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##########Delete Compnay  STARTED #############")

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
	var company_id int
	params := mux.Vars(r)
	company_id, err = strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something wrong with Id"))
		return
	}
	deleted, err := MyDb.DeleteCompany(company_id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if deleted <= 0 {

		w.Write([]byte("Invalid Id"))
		return
	}
	w.Write([]byte("Deleted"))
}
