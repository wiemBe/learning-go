package render

import (
	"bytes"
	"fmt"
	"github.com/wiemBe/learning-go/pkg/config"
	"github.com/wiemBe/learning-go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		// create a template cache
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()

	}
	// get request from cache

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}
	buff := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)

	_ = t.Execute(buff, templateData)
	// render template

	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	// range through it
	for _, page := range pages {
		// it gets rid of the ./ stuff and gives you file name
		name := filepath.Base(page)
		// ts = template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("/templates/*.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.html")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
