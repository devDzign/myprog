package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.go.html")

	if err != nil {
		fmt.Println("Error parsed template:", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsed template:", err)
		return
	}
}
