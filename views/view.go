package views

import "html/template"





func NewView(layout string, files ...string) *View{
        files = append(files,  "views/layouts/body.gohtml", "views/layouts/footer.gohtml")
        t:= template.Must(template.ParseFiles(files...))
        return &View{Template:t, Layout: layout}
}

type View struct{
    Template  *template.Template
    Layout string
}

