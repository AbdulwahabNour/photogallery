package main

import (
	"fmt"
	"net/http"
)

func main(){
  http.HandleFunc("/",  handlerFunc)
  http.ListenAndServe(":8080", nil)
}
func handlerFunc(w http.ResponseWriter, r *http.Request){
    
      fmt.Fprintf(w, "<h1>Welcome </h1> ")
}