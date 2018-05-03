package main

import "net/http"
import "encoding/json"

type HappinessStatus struct {
    OverallPercentage int `json:"overallPercentage"`
}

type HappinessSubmission struct {
    Percentage int `json:"percentage"`
}

var cumulative int
var submissionsCount int

func HappinessGetHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    overall := 0
    if submissionsCount > 0 {
        overall = cumulative / submissionsCount
    }
    json.NewEncoder(w).Encode(HappinessStatus{overall})
}

func HappinessPostHandler(w http.ResponseWriter, r *http.Request) {
    var submission HappinessSubmission
    err := json.NewDecoder(r.Body).Decode(&submission)
    checkError(err)

    cumulative += submission.Percentage
    submissionsCount++
}
