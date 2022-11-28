package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"placementAPI/http/company"
	"placementAPI/http/student"
	"placementAPI/store"
	"placementAPI/svc"
)

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)

		return
	}
	str1 := store.NewCompanyStorer(db)
	str2 := store.NewStudentStorer(db)
	service1 := svc.NewCompanyServiceHandler(str1)
	service2 := svc.NewStudentServiceHandler(str2)
	h1 := company.NewCompanyHandler(service1)
	h2 := student.NewStudentHandler(service2)
	router := mux.NewRouter()
	router.HandleFunc("/companies", h1.GetCompanyData).Methods("GET")
	router.HandleFunc("/companies", h1.PostCompanyData).Methods("POST")
	router.HandleFunc("/companies", h1.UpdateCompanyData).Methods("PUT")
	router.HandleFunc("/companies", h1.DeleteCompanyData).Methods("DELETE")

	router.HandleFunc("/students", h2.GetStudentData).Methods("GET")
	router.HandleFunc("/students", h2.PostStudentData).Methods("POST")
	router.HandleFunc("/students", h2.UpdateStudentData).Methods("PUT")
	router.HandleFunc("/students", h2.DeleteStudentData).Methods("DELETE")

	//http.HandleFunc("/student", h1.StudentsHandler)
	//http.HandleFunc("/company", h2.CompHandler)
	http.ListenAndServe(":8000", nil)
}
