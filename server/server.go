package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type UpdateRequest struct {
    Water float64 `json:"water"`
    Wind  float64 `json:"wind"`
}

type UpdateResponse struct {
    Water  float64 `json:"water"`
    Wind   float64 `json:"wind"`
    Status string  `json:"status"`
}

func main() {
    http.HandleFunc("/update", handleUpdate)
    http.HandleFunc("/", handleNotFound)
    http.ListenAndServe(":8080", nil)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Decode the request body
    var updateReq UpdateRequest
    err := json.NewDecoder(r.Body).Decode(&updateReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Calculate status
    status := updateStatus(updateReq.Water, updateReq.Wind)

    // Prepare response
    updateRes := UpdateResponse{
        Water:  updateReq.Water,
        Wind:   updateReq.Wind,
        Status: status,
    }

    // Encode response as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updateRes)
}

func updateStatus(water, wind float64) string {
    var waterStatus, windStatus string

    if water < 5 {
        waterStatus = "Aman"
    } else if water >= 6 && water <= 8 {
        waterStatus = "Siaga"
    } else {
        waterStatus = "Bahaya"
    }

    if wind < 6 {
        windStatus = "Aman"
    } else if wind >= 7 && wind <= 15 {
        windStatus = "Siaga"
    } else {
        windStatus = "Bahaya"
    }

    return fmt.Sprintf("Status: Water(%s), Wind(%s)", waterStatus, windStatus)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Coba lagi ya", http.StatusNotFound)
}
