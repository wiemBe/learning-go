package handlers

import (
	"github.com/wiemBe/learning-go/pkg/render"
	"net/http"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
