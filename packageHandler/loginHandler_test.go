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

func TestLoginHandler(t *testing.T) {
	// initialize templates
	InitTemplates()

	// create request which will later be passed to the handler
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
}
