package views

import "html/template"

func NewView( templ ...string) *View {
    // append footer to tmplates
     templ = append(templ,"views/layouts/body.gohtml","views/layouts/navbar.gohtml")
       
     t, err := template.ParseFiles(templ...);
     if err != nil {
         panic(err)
     }

     return &View{Template: t, Body: "body"}
}

type View struct{
     Template *template.Template
     Body string
}

  