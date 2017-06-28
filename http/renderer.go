package http

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Renderer struct {
	templates   *template.Template
	defaultData map[string]interface{}
}

func NewRenderer(data map[string]interface{}) *Renderer {
	return &Renderer{
		templates:   template.Must(template.ParseGlob("web/*.html")),
		defaultData: data,
	}
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	mapData := data.(map[string]interface{})
	for k, v := range r.defaultData {
		mapData[k] = v
	}

	return r.templates.ExecuteTemplate(w, name, mapData)
}

func WebHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}
