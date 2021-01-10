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

// https://blog.questionable.services/article/testing-http-handlers-go/
func TestIndexHandler(t *testing.T) {
	// initialize templates
	InitTemplates()

	// create request which will later be passed to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
}
