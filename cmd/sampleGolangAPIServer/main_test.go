package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", bytes.NewBufferString(""))
	recorder := httptest.NewRecorder()

	helloHandler(recorder, req)

	expected := `{"Message":"hello","Version":":VERSION:"}
`

	got := recorder.Body.String()
	if got != expected {
		t.Errorf(cmp.Diff(got, expected))
	}
}
