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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))
	listenAndServe := http.ListenAndServe(portNumber, nil)

	if listenAndServe != nil {
		fmt.Println(listenAndServe)

	}

}
