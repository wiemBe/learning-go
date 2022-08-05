package main

import (
	"fmt"
	"github.com/wiemBe/learning-go/pkg/config"
	"github.com/wiemBe/learning-go/pkg/handlers"
	"github.com/wiemBe/learning-go/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))
	listenAndServe := http.ListenAndServe(portNumber, nil)

	if listenAndServe != nil {
		fmt.Println(listenAndServe)

	}

}
