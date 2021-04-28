package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/views"
)
 
var (
  homeTempl * views.View //0xc00012a3e0
  contactTempl * views.View
  faq   *  views.View
  signupTemp * views.View 
)
 
func home(w http.ResponseWriter, req *http.Request){
      w.Header().Set("Content-Type", "text/html")
      
      if err := homeTempl.Template.ExecuteTemplate(w,contactTempl.Body  , nil); err != nil{
         panic(err)
      } 
}

func contact (w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "text/html")

  if err := contactTempl.Template.ExecuteTemplate(w, contactTempl.Body ,nil); err != nil{
    panic( err)
  }
  
}
func faqPage( w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-Type", "text/html")
   if err := faq.Template.ExecuteTemplate(w, faq.Body, nil); err != nil{
      panic(err)
   }
}
 
 
func main() {
  homeTempl = views.NewView("views/home.gohtml")//0xc00012a3e0
  contactTempl = views.NewView("views/contact.gohtml")
  faq = views.NewView("views/faq.gohtml")

  newUserControllers  := controllers.NewUser()

  
 r := mux.NewRouter()
 
 r.HandleFunc("/", home).Methods("GET")
 r.HandleFunc("/faq", faqPage).Methods("GET")
 r.HandleFunc("/contact", contact).Methods("GET")
 r.HandleFunc("/signup", newUserControllers.New).Methods("GET")
 r.HandleFunc("/signup", newUserControllers.Create).Methods("POST")
 ///views/layouts/js/jquery-3.1.1.min.js
 
 http.ListenAndServe(":3000", r)
 
}