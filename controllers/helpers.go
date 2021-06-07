package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)


func parseForm(req * http.Request, dst interface{}) error {

        if err := req.ParseForm(); err != nil{
                panic(err)  
        }
        var decoder = schema.NewDecoder()
        
        err := decoder.Decode(dst, req.PostForm)
        if err != nil{
                panic(err)  
        }
        return nil 
}

 