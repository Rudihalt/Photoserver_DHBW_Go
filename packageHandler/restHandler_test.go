/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRESTHandler(t *testing.T) {
	// initialize templates
	InitTemplates()

	// create request which will later be passed to the handler
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RESTHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusSeeOther, status)
}
