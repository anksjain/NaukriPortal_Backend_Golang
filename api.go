package main

import (
	"fmt"
	"log"
	"net/http"
	Auth "restapi/package/Auth"
	Companies "restapi/package/Companies"
	Jobs "restapi/package/Jobs"
	Users "restapi/package/Users"

	MyDb "restapi/package/Database"
	middle "restapi/package/Middleware"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter()
	app_port := ":" + MyDb.GetEnv("APP_PORT")
	//ADDED FOR CORS
	myRouter.Use(middle.Middleware)
	//AUTH
	myRouter.HandleFunc("/login", Auth.GetLogin).Methods(http.MethodPost, "OPTIONS")
	myRouter.HandleFunc("/register", Auth.Register).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/logout", Auth.Logout).Methods("GET", "OPTIONS")

	// USERS (ADMIN)
	myRouter.HandleFunc("/users", Users.GetUsers).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/users/{id}", Users.GetUser).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/users/{id}", Users.DeleteUser).Methods("DELETE", "OPTIONS")
	myRouter.HandleFunc("/users/{id}", Users.UpdateUser).Methods("PUT", "OPTIONS")

	//JOBS
	myRouter.HandleFunc("/jobs", Jobs.GetJobs).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/jobs", Jobs.PostJob).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/jobs/{id}", Jobs.GetJob).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/jobs/{id}", Jobs.ApplyJob).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/jobs/{id}", Jobs.UpdateJob).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/jobs/{id}", Jobs.DeleteJob).Methods("DELETE", "OPTIONS")

	//Company
	myRouter.HandleFunc("/company", Companies.GetCompanies).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/company/{id}", Companies.GetCompany).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/company", Companies.CreateCompany).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/company/{id}", Companies.DeleteCompany).Methods("DELETE", "OPTIONS")

	fmt.Println("PORT RUNNING :", app_port)
	MyDb.MigrationSeed()
	log.Fatalln(http.ListenAndServe(app_port, myRouter))
}
