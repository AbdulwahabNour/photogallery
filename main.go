package main

import (
	"fmt"
	"net/http"
)

func handlerFunc( w http.ResponseWriter, r * http.Request){
	 fmt.Fprint(w, "<h1>Hello</h1>")
}
func main() {
	fmt.Println("Hello")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}