package main

import "testing"
import "net/http/httptest"
import "net/http"

func TestDefaultHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/wrong_url", nil)
    checkTestError(err, t)

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(DefaultHandler)
    handler.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusNotFound {
        t.Errorf("Received HTTP status %v instead of %v", recorder.Code, http.StatusNotFound)
    }
}

func checkTestError(err error, t *testing.T) {
    if err != nil {
        t.Fatal(err)
    }
}
