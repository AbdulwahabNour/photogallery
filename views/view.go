package views

import (
	"html/template"
	"net/http"
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
//Qazwsxedc11@111z
//Qazwsxedc11@111z2
type View struct{
     Template *template.Template
     Body string
}
func (v *View) Render(w http.ResponseWriter, data interface{}) error{
        err := v.Template.ExecuteTemplate(w, v.Body, data)
        return err
}

func layoutFiles() []string{
    files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
    if err != nil {
        panic(err)
    }
    
    return files
}

  