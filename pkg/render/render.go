package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {

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

var tCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	//check to see if we already have the template in our cache

	_, inMap := tCache[t]

	if !inMap {
		// you need create template and add in cache
		log.Println("Create template and add to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cache template")
	}
	// we have template in cache
	tmpl = tCache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.go.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tCache[t] = tmpl

	return nil
}
