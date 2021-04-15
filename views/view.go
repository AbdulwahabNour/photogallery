package views

import "html/template"

func NewView(templ ...string) *View {
    // append footer to tmplates
     templ = append(templ,"views/layouts/footer.gohtml", "views/layouts/header.gohtml")
     
     t, err := template.ParseFiles(templ...);
     if err != nil {
         panic(err)
     }
     return &View{Template: t}
}

type View struct{
     Template *template.Template
}