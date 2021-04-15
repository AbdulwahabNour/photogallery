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
      if err := homeTempl.Template.Execute(w, nil); err != nil{
         panic(err)
      } 
}

func contact (w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "text/html")
  if err := contactTempl.Template.Execute(w, nil); err != nil{
    panic( err)
  }
}
func main() {
  homeTempl = views.NewView("views/home.gohtml",  "views/layouts/footer.gohtml")
  contactTempl = views.NewView("views/contact.gohtml", "views/layouts/footer.gohtml")
  
 r := mux.NewRouter()
 r.HandleFunc("/", home)
 r.HandleFunc("/contact", contact)
 http.ListenAndServe(":3000", r)
 
}