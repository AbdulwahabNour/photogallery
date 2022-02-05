package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func main() {
	homeTemplate = template.Must(template.ParseFiles("views/home.gohtml"))
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":8080", r)

}

func home(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
	}{
		Name: "Ahmed",
	}
	err := homeTemplate.Execute(w, data)
	if err != nil {
		log.Fatalln(err)
	}
}
