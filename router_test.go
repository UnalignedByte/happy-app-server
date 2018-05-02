package main

import "testing"
import "net/http/httptest"
import "net/http"


func TestDefaultHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/wrong_url", nil)
    testCheckError(err, t)

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(defaultHandler)
    handler.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusNotFound {
        t.Errorf("Received HTTP status %v instead of %v", recorder.Code, http.StatusNotFound)
    }
}

func testCheckError(err error, t *testing.T) {
    if err != nil {
        t.Fatal(err)
    }
}
