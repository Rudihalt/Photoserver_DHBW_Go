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
	"time"
)

func TestDiashowHandlerNoToken(t *testing.T) {
	// initialize templates
	InitTemplates()

	// create request which will later be passed to the handler
	req, err := http.NewRequest("GET", "/diashow", nil)
	if err != nil {
		t.Fatal(err)
	}
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DiashowHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusSeeOther, status)
}

func TestDiashowHandler(t *testing.T) {
	// initialize templates
	InitTemplates()
	// test user token: d003b47ba0d9ade5f482e7e0

	// create request which will later be passed to the handler
	req, err := http.NewRequest("GET", "/diashow", nil)
	if err != nil {
		t.Fatal(err)
	}

	expiration := time.Now().Add(10 * time.Second)
	cookie := http.Cookie{Name: "csrftoken", Value: "d003b47ba0d9ade5f482e7e0", Expires: expiration}
	req.AddCookie(&cookie)
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DiashowHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
}
