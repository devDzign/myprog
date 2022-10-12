package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// create a template cache

	cache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested templete from cache
	t, ok := cache[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)

	if err != nil {
		log.Fatal(err)
	}
	/// render the template

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template) same as 40
	myCache := map[string]*template.Template{}

	// get all the files named *.page.html from ./templates

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	// range through all file ending with *.page.html

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.go.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*go.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
