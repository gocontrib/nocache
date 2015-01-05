package nocache

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gohttp/app"
)

func TestNoCache(t *testing.T) {

	rr := httptest.NewRecorder()
	s := app.New()

	s.Use(New())
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	s.ServeHTTP(rr, r)

	for k, v := range noCacheHeaders {
		if rr.HeaderMap[k][0] != v {
			t.Errorf("%s header not set by middleware.", k)
		}
	}
}
