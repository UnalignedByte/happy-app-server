package main

import "testing"
import "net/http/httptest"
import "net/http"
import "encoding/json"
import "bytes"

func TestHapinessGet(t *testing.T) {
    request, err := http.NewRequest("GET", "/happiness", nil)
    checkTestError(err, t)

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(HappinessGetHandler)
    handler.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusOK {
        t.Errorf("Received HTTP status %v instead of %v", recorder.Code, http.StatusOK)
    }

    decoder := json.NewDecoder(recorder.Body)
    var status HappinessStatus
    err  = decoder.Decode(&status)
    checkTestError(err, t)
    if status.OverallPercentage != 0 {
        t.Errorf("HappinessStatus.Overall is %v instead of %v", status.OverallPercentage, 0)
    }
}

func TestHappinessPost(t *testing.T) {
    // First submit new data
    data, err := json.Marshal(HappinessSubmission{40})
    checkTestError(err, t)

    request, err := http.NewRequest("POST", "/api/happiness", bytes.NewBuffer(data))
    checkTestError(err, t)

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(HappinessPostHandler)
    handler.ServeHTTP(recorder, request)

    // Submit another data
    data, err = json.Marshal(HappinessSubmission{20})
    checkTestError(err, t)

    request, err = http.NewRequest("POST", "/api/happiness", bytes.NewBuffer(data))
    checkTestError(err, t)

    recorder = httptest.NewRecorder()
    handler = http.HandlerFunc(HappinessPostHandler)
    handler.ServeHTTP(recorder, request)

    // Then check if we can get it back
    request, err = http.NewRequest("GET", "/api/happiness", nil)
    checkTestError(err, t)

    recorder = httptest.NewRecorder()
    handler = http.HandlerFunc(HappinessGetHandler)
    handler.ServeHTTP(recorder, request)

    var status HappinessStatus
    err = json.NewDecoder(recorder.Body).Decode(&status)
    checkTestError(err, t)

    if recorder.Code != http.StatusOK {
        t.Errorf("Received HTTP Code %v instead of %v", recorder.Code, http.StatusOK)
    }

    if status.OverallPercentage != 30 {
        t.Errorf("HappinessStatus.Overall is %v instead of %v", status.OverallPercentage, 30)
    }
}
