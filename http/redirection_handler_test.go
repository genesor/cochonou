package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"

	"github.com/genesor/cochonou"
	cochonou_http "github.com/genesor/cochonou/http"
	"github.com/genesor/cochonou/mock"
)

func setupRedirHandler() (*cochonou_http.RedirectionHandler, *mock.DomainHandler) {
	handler := &mock.DomainHandler{}

	httpHandler := cochonou_http.NewRedirectionHandler(handler)

	return httpHandler, handler
}

func TestCreateRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		e := echo.New()

		body := `{"target":"http://sadoma.so", "sub_domain": "oklm"}`
		req := httptest.NewRequest(echo.POST, "/redirections", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/redirections")

		httpHandler, domainHandler := setupRedirHandler()
		domainHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "oklm", sub)
			require.Equal(t, "http://sadoma.so", target)

			return nil
		}

		err := httpHandler.HandleCreate(c)

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rec.Code)
		require.Equal(t, "", rec.Body.String())
		require.Equal(t, 1, domainHandler.CreateDomainRedirectionCall)
	})

	t.Run("NOK - Crappy body", func(t *testing.T) {
		e := echo.New()

		body := `Bad Body`
		req := httptest.NewRequest(echo.POST, "/redirections", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/redirections")

		httpHandler, domainHandler := setupRedirHandler()

		err := httpHandler.HandleCreate(c)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rec.Code)
		require.JSONEq(t, `
			{
				"code": "malformatted_body",
				"message": "Malformatted body, unable to deserialize.",
				"details": "code=400, message=Syntax error: offset=1, error=invalid character 'B' looking for beginning of value"
			}
		`, rec.Body.String())
		require.Equal(t, 0, domainHandler.CreateDomainRedirectionCall)
	})

	t.Run("NOK - invalid payload", func(t *testing.T) {
		e := echo.New()

		body := `{"target":"i am not an URL", "sub_domain": "oklm"}`
		req := httptest.NewRequest(echo.POST, "/redirections", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/redirections")

		httpHandler, domainHandler := setupRedirHandler()

		err := httpHandler.HandleCreate(c)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rec.Code)
		require.JSONEq(t, `
			{
				"message": "Error in payload data.",
				"details": "target: i am not an URL does not validate as url;",
				"code": "invalid_payload"
			}
		`, rec.Body.String())
		require.Equal(t, 0, domainHandler.CreateDomainRedirectionCall)
	})

	t.Run("NOK - Error DomainHandler", func(t *testing.T) {
		e := echo.New()

		body := `{"target":"http://sadoma.so", "sub_domain": "oklm"}`
		req := httptest.NewRequest(echo.POST, "/redirections", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/redirections")

		httpHandler, domainHandler := setupRedirHandler()
		domainHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "oklm", sub)
			require.Equal(t, "http://sadoma.so", target)

			return errors.New("some error")
		}

		err := httpHandler.HandleCreate(c)

		require.NoError(t, err)
		require.Equal(t, http.StatusServiceUnavailable, rec.Code)
		require.JSONEq(t, `
			{
				"message": "Unexpected error.",
				"details": "some error",
				"code": "error_internal"
			}
		`, rec.Body.String())
		require.Equal(t, 1, domainHandler.CreateDomainRedirectionCall)
	})

	t.Run("NOK - Error DomainHandler alreadyExists", func(t *testing.T) {
		e := echo.New()

		body := `{"target":"http://sadoma.so", "sub_domain": "oklm"}`
		req := httptest.NewRequest(echo.POST, "/redirections", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/redirections")

		httpHandler, domainHandler := setupRedirHandler()
		domainHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "oklm", sub)
			require.Equal(t, "http://sadoma.so", target)

			return cochonou.ErrSubDomainAlreadyExists
		}

		err := httpHandler.HandleCreate(c)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rec.Code)
		require.JSONEq(t, `
			{
				"message": "Subdomain already used.",
				"details": "the subdomain cannot be created because it already exists",
				"code": "subdomain_unavailable"
			}
		`, rec.Body.String())
		require.Equal(t, 1, domainHandler.CreateDomainRedirectionCall)
	})
}
