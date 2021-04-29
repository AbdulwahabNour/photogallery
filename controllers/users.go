package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

//function to parse signup page new page

func NewUser() *User{
    return &User{ NewView: views.NewView("views/users/new.gohtml")}
}

type User struct{
     NewView *views.View
}
type SignupForm struct{
    Email string `schema:"email"`
    Password string `schema:"password"`

}
//Render the signup page (views/users/new.gohtml) 
//create  a new user account
//GET /signup
func(u *User)New(w http.ResponseWriter, r * http.Request){
    
   if err:= u.NewView.Render(w,nil); err != nil {
        panic(err)
   }

}

//Create  is used to process th signup form when submit the form
//POST /signup
func(u *User)Create(w http.ResponseWriter, req *http.Request){
     var dataForm SignupForm
     
     if err := parseForm(req, &dataForm); err != nil{
          panic(err)
     }
       fmt.Fprintln(w, dataForm)
}