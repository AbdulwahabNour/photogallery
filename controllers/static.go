package controllers

import "lenslocked.com/views"

 
func NewStatic()* Static{
        return &Static{ 
                        HomeView: views.NewView("static/home") ,
                        ContactView:views.NewView("static/contact"),
                        Faq: views.NewView("static/faq"),
                       }
}

type Static struct{
    HomeView * views.View
    ContactView * views.View
    Faq *views.View
 }
