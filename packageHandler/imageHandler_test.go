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

func TestImageHandlerNoToken(t *testing.T) {
	// initialize templates
	InitTemplates()

	// create request which will later be passed to the handler
	req, err := http.NewRequest("GET", "/image", nil)
	if err != nil {
		t.Fatal(err)
	}
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ImageHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusSeeOther, status)
}

func TestImageHandler(t *testing.T) {
	// initialize templates
	InitTemplates()
	// hneemann user token: a9767771509cb03991356332

	// create request which will later be passed to the handler
	// image=09a5fee1b9233ecc8c5ef25bf5030066bf43103690a15b3a0847867a59aad542

	req, err := http.NewRequest("GET", "/image?image=09a5fee1b9233ecc8c5ef25bf5030066bf43103690a15b3a0847867a59aad542", nil)
	if err != nil {
		t.Fatal(err)
	}

	expiration := time.Now().Add(10 * time.Second)
	cookie := http.Cookie{Name: "csrftoken", Value: "a9767771509cb03991356332", Expires: expiration}
	req.AddCookie(&cookie)
	// create ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ImageHandler)
	// pass the Request and the ResponseRecorder with calling ServeHTTP
	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
}
