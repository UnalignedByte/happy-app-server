package main

import "testing"
import "net/http/httptest"
import "net/http"
import "encoding/json"


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

func TestHapinessHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/happiness", nil)
    checkTestError(err, t)

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(HappinessHandler)
    handler.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusOK {
        t.Errorf("Received HTTP status %v instead of %v", recorder.Code, http.StatusOK)
    }

    decoder := json.NewDecoder(recorder.Body)
    var status HappinessStatus
    err  = decoder.Decode(&status)
    checkTestError(err, t)
    if status.Overall != 0 {
        t.Errorf("HappinessStatus.Overall is %v instead of %v", status.Overall, 0)
    }
}

func checkTestError(err error, t *testing.T) {
    if err != nil {
        t.Fatal(err)
    }
}
