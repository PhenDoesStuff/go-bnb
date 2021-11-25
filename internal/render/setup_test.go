package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/stephenmontague/go-bnb/internal/config"
	"github.com/stephenmontague/go-bnb/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// SCS is a package that helps us set up a session so people aren't logged out immediately
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {}

func (tw *myWriter) Write(b []byte) (int, error) {
	// Can't just return a random int
	length := len(b)
	return length, nil
}