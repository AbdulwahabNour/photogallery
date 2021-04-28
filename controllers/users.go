package controllers

import (
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

func(u *User)New(w http.ResponseWriter, r * http.Request){
    
   if err:= u.NewView.Render(w,nil); err != nil {
        panic(err)
   }

}
 