package views

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

var (
    LayoutDir = "views/layouts/"
    templateDir = "views/"
    TemplateExt = ".gohtml"
)

func NewView( templ ...string) *View {
 
     addTemplateDir(templ)
 
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
func (v *View)ServeHTTP(w http.ResponseWriter, r * http.Request){
        v.Render(w, nil) 
}
func (v *View) Render(w http.ResponseWriter, data interface{})  {

      switch data.(type){
      case Data:
      default:
        data = Data{Yield: data}
      }

       var buf bytes.Buffer

       if  err := v.Template.ExecuteTemplate(&buf, v.Body, data); err!= nil{
             http.Error(w, AlertMsgGeneric, http.StatusInternalServerError)
             return
       }
       io.Copy(w, &buf)
}

func layoutFiles() []string{
    files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
    if err != nil {
        panic(err)
    }
    
    return files
}

//addTemplateDir take slice of string
//as file path and it appends templateDir and template extension
//
//Eg the input {"contact"} the output would be {"view/contact.gohtml"} 
//if templateDir is 'view/' and TemplateExt is '.gohtml' 
func addTemplateDir(files []string){
    for i, f := range files{
            files[i] = templateDir + f + TemplateExt
    }
}




  
