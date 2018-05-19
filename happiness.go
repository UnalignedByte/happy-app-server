package main

import "net/http"
import "encoding/json"
import "time"

type HappinessStatus struct {
    OverallPercentage int `json:"overallPercentage"`
}

type HappinessSubmission struct {
    Percentage int `json:"percentage"`
}

func HappinessGetHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    overall := GetOverallHappiness()
    json.NewEncoder(w).Encode(HappinessStatus{overall})
}

func HappinessPostHandler(w http.ResponseWriter, r *http.Request) {
    var submission HappinessSubmission
    err := json.NewDecoder(r.Body).Decode(&submission)
    checkError(err)

    AddHappiness(submission.Percentage, time.Now())
}
