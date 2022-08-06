package handlers

import (
	"github.com/wiemBe/learning-go/pkg/config"
	"github.com/wiemBe/learning-go/pkg/render"
	"net/http"
)

var Repo *Repository

// TemplateData holds  data sent from handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html", &TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["hello"] = "there"

	// send data to template
	render.RenderTemplate(w, "about.html", &TemplateData{
		StringMap: stringMap,
	})
}
