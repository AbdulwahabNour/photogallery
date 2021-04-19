package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)
 
var (
  homeTempl * views.View
  contactTempl * views.View
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
 
func main() {
  homeTempl = views.NewView("views/layouts/home.gohtml")
  contactTempl = views.NewView("views/layouts/contact.gohtml")
  
 r := mux.NewRouter()
 r.HandleFunc("/", home)
 r.HandleFunc("/contact", contact)
 ///views/layouts/js/jquery-3.1.1.min.js
 
 http.ListenAndServe(":3000", r)
 
}