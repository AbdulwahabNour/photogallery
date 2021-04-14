package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)
 
var (
  homeTempl * template.Template
  contactTempl * template.Template
)


 
func home(w http.ResponseWriter, req *http.Request){
      w.Header().Set("Content-Type", "text/html")
      if err := homeTempl.Execute(w, nil); err != nil{
         panic(err)
      } 
}

func contact (w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "text/html")
  if err := contactTempl.Execute(w, nil); err != nil{
    panic( err)
  }
}
func main() {
   var err error
   homeTempl, err = template.ParseFiles("views/home.gohtml")
   if err != nil{
     panic(err)
   }

  contactTempl, err = template.ParseFiles("views/contact.gohtml")
  if err != nil{
    panic(err)
  }

 r := mux.NewRouter()
 r.HandleFunc("/", home)
 r.HandleFunc("/contact", contact)
 http.ListenAndServe(":3000", r)
}