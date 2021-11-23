package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/stephenmontague/go-bnb/internal/config"
	"github.com/stephenmontague/go-bnb/internal/handlers"
	"github.com/stephenmontague/go-bnb/internal/models"
	"github.com/stephenmontague/go-bnb/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {	
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting application on port %s \n", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	// SCS is a package that helps us set up a session so people aren't logged out immediately
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // In production we want this to be true. False for dev b/c localhost is not secure

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}