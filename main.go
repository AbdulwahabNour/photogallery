package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)
 
var (
  homeTempl * views.View
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

func signup(w http.ResponseWriter, req * http.Request){
      w.Header().Set("Content-Type", "text/html")
      if err := signupTemp.Template.ExecuteTemplate(w, faq.Body, nil); err != nil{
        panic( err)
      }

}
 
func main() {
  homeTempl = views.NewView("views/home.gohtml")
  contactTempl = views.NewView("views/contact.gohtml")
  faq = views.NewView("views/faq.gohtml")
  signupTemp = views.NewView("views/signup.gohtml")

  
 r := mux.NewRouter()
 r.HandleFunc("/", home)
 r.HandleFunc("/faq", faqPage)
 r.HandleFunc("/contact", contact)
 r.HandleFunc("/signup", signup)
 ///views/layouts/js/jquery-3.1.1.min.js
 
 http.ListenAndServe(":3000", r)
 
}