package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

//function to parse signup page new page

func NewUser(us *models.UserService) *User{
   
    return &User{ NewView: views.NewView("users/new"), userServ:us}
}

type User struct{
     NewView *views.View
     userServ *models.UserService
}
type SignupForm struct{
    Email string `schema:"email,required"`
    Name string `schema:"name,required"`
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

     user := models.User{
          Email: dataForm.Email,
          Name: dataForm.Name,
     }
     if err := u.userServ.Create(&user); err != nil{
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
     }
     fmt.Fprintln(w, dataForm)
}