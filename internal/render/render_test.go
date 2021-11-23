package render

import (
	"github.com/stephenmontague/go-bnb/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	result := AddDefaultData(&td, r)
	if result == nil {
		t.Error("failed")
	}
}