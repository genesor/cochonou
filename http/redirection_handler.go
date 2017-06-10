package http

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"

	"github.com/genesor/cochonou"
)

// RedirectionHandler is the struct containing all functions to handle HTTP actions
// with Redirections.
type RedirectionHandler struct {
	DomainHandler cochonou.DomainHandler
}

// CreatePayload represents the JSON payload for HandleCreate.
type CreatePayload struct {
	Target    string `json:"target,omitempty" valid:"url,required"`
	SubDomain string `json:"sub_domain,omitempty" valid:"alphanum,required"`
}

// HandleCreate creates a DomainRedirection from an HTTP request.
func (h *RedirectionHandler) HandleCreate(c echo.Context) error {
	payload := new(CreatePayload)

	err := c.Bind(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewJSONError("malformatted_body", "Malformatted body, unable to deserialize.", err))
	}

	isValid, err := valid.ValidateStruct(payload)
	if isValid == false {
		return c.JSON(http.StatusBadRequest, NewJSONError("invalid_payload", "Error in payload data.", err))
	}

	err = h.DomainHandler.CreateDomainRedirection(payload.SubDomain, payload.Target)
	if err != nil {
		if err == cochonou.ErrSubDomainAlreadyExists {
			return c.JSON(http.StatusBadRequest, NewJSONError("subdomain_unavailable", "Subdomain already used.", err))
		}

		// Add log
		return c.JSON(http.StatusServiceUnavailable, NewJSONError("error_internal", "Unexpected error.", err))
	}

	return c.NoContent(http.StatusCreated)
}
