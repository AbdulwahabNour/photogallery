package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

//function to parse signup page new page

func NewUser(us *models.UserService) *User{
   
    return &User{ NewView: views.NewView("users/new"),
                  LoginView:views.NewView("users/login"), 
                  userServ:us}
}

type User struct{
     NewView *views.View
     LoginView *views.View
     userServ *models.UserService
}
type SignupForm struct{
     Email string `schema:"email,required"`
     Name string `schema:"name,required"`
     Password string `schema:"password"`
}

type LoginForm struct{
     Email string `schema:email,required`
     Password string `schema:password`
} 
 

//Create  is used to process th signup form when submit the form
//POST /signup
func(u *User)Create(w http.ResponseWriter, req *http.Request){
     var dataForm SignupForm
     
     parseForm(req, &dataForm)
     
     user := models.User{
          Email: dataForm.Email,
          Name: dataForm.Name,
          Password: dataForm.Password,
     }
     if err := u.userServ.Create(&user); err != nil{
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
     }
     fmt.Fprintln(w, dataForm,user.Password )
}
//Handle login Post request
//By check email address and password is correct or no
//and login if are correct

func (u *User)Login(w http.ResponseWriter, req *http.Request){
        form := LoginForm{}
        parseForm(req, &form)
        fmt.Fprintln(w, form)
        
}