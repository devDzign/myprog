package handlers

import (
	"net/http"
	"webHello/pkg/config"
	"webHello/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRFToken  string
	Flash     string
	Warning   string
	Error     string
}

// NewRepo creates a new repositories
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "home.page.html", &TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, request *http.Request) {
	var stringMap map[string]string
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.html", &TemplateData{StringMap: stringMap})
}
