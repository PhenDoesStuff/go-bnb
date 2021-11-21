package handlers

import (
	"net/http"

	"github.com/stephenmontague/go-bnb/pkg/config"
	"github.com/stephenmontague/go-bnb/pkg/models"
	"github.com/stephenmontague/go-bnb/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Next 2 lines of code stores the remote IP in the session with a key to look it up
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Warlocks renders the warlocks room page
func (m *Repository) Warlocks (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "warlocks.page.tmpl", &models.TemplateData{})
}

// Warriors renders the warriors room page
func (m *Repository) Warriors (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "warriors.page.tmpl", &models.TemplateData{})
}

// Availability renders the availability page
func (m *Repository) Availability (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}