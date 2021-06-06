package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/models"
)
 
const (
	host = "localhost"
	port = 5433
	user = "postgres"
	password = "admin"
	dbname = "lenlocked"
)

func main() {
   
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := models.NewUserService(psqlInfo)
	if   err != nil{
		panic(err)
	}
   defer db.Close()
 
  newStatic := controllers.NewStatic()
  newUserControllers  := controllers.NewUser(db)
  
 r := mux.NewRouter()
 
 r.Handle("/",newStatic.HomeView).Methods("GET")
 r.Handle("/contact", newStatic.ContactView).Methods("GET")
 r.Handle("/faq", newStatic.Faq).Methods("GET")
 r.HandleFunc("/signup", newUserControllers.New).Methods("GET")
 r.HandleFunc("/signup", newUserControllers.Create).Methods("POST")
 
 http.ListenAndServe(":3000", r)
 
}