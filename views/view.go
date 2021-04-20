package views

import (
	"html/template"
	"path/filepath"
)

var (
    LayoutDir = "views/layouts/"
    TemplateExt = ".gohtml"
)

func NewView( templ ...string) *View {
    // append footer to tmplates
     templ = append(templ,layoutFiles()...)
       
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

func layoutFiles() []string{
    files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
    if err != nil {
        panic(err)
    }
    
    return files
}

  