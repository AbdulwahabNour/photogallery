package main

import (
	"log"
	"net/http"

	"github.com/AbdulwahabNour/photogallery/views"
	"github.com/gorilla/mux"
)

var (
	homeTemplate *views.View
)

func main() {
	homeTemplate = views.NewView("body","views/home.gohtml")
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
	err := homeTemplate.Template.ExecuteTemplate(w, homeTemplate.Layout,data)
	if err != nil {
		log.Fatalln(err)
	}
}
