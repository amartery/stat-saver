package statserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatServer_handleShow(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/stat/show?from=2000-01-01&to=2020-01-01", nil)

	s.handleShow().ServeHTTP(rec, req)
	assert.Equal(t, rec.Result().StatusCode, 200)
}
