package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)
 
var templ * template.Template

 
func home(w http.ResponseWriter, req *http.Request){
      w.Header().Set("Content-Type", "text/html")
      if err := templ.Execute(w, nil); err != nil{
         panic(err)
      }
}
func main() {
   var err error
   templ, err = template.ParseFiles("views/home.gohtml")
   if err != nil{
     panic(err)
   }
 r := mux.NewRouter()
 r.HandleFunc("/", home)
 
 http.ListenAndServe(":3000", r)
}