package handlers

import (
	"github.com/AseelSaghir7/Bookings/pkg/config"
	"github.com/AseelSaghir7/Bookings/pkg/models"
	"github.com/AseelSaghir7/Bookings/pkg/render"
	"net/http"
)

// Repo - The reason I'm saving main's app variable to new variable. Because I don't want to make changes in original variable instead we will use copy of it and use it for render and handlers pkg
//Repo - the repository used by the handlers
var Repo *Repository

// Repository is the repository type
//Saving into struct because in future we may add more things in struct...
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
