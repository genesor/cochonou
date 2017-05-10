package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	twingo_http "github.com/genesor/cochonou/http"
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	h := twingo_http.NewHelloHandler()

	// Assertions
	if assert.NoError(t, h.HandleHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
