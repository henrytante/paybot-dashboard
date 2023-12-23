package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadHtml()  {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func RunHtml(w http.ResponseWriter, template string, dados interface{})  {
	LoadHtml()
	templates.ExecuteTemplate(w, template, dados)
}