package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)
 
func main() {
   
  newStatic := controllers.NewStatic()
  newUserControllers  := controllers.NewUser()
  
 r := mux.NewRouter()
 
 r.Handle("/",newStatic.HomeView).Methods("GET")
 r.Handle("/contact", newStatic.ContactView).Methods("GET")
 r.Handle("/faq", newStatic.Faq).Methods("GET")
 r.HandleFunc("/signup", newUserControllers.New).Methods("GET")
 r.HandleFunc("/signup", newUserControllers.Create).Methods("POST")
 ///views/layouts/js/jquery-3.1.1.min.js
 
 http.ListenAndServe(":3000", r)
 
}