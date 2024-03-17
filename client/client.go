package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UpdateResponse struct {
	Water  float64 `json:"water"`
	Wind   float64 `json:"wind"`
	Status string  `json:"status"`
}

func main() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		sendUpdate(10.0, 6.0) // Contoh nilai water dan wind
	}
}

func sendUpdate(water, wind float64) {
	url := "http://localhost:8080/update"
	status := updateStatus(water, wind)
	fmt.Printf("Sending update to server: %s\n", status)

	// Prepare request body
	reqBody, err := json.Marshal(map[string]float64{"water": water, "wind": wind})
	if err != nil {
		fmt.Println("Error encoding request body:", err)
		return
	}

	// Send HTTP POST request to server
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error sending update:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Unexpected status code %d\n", resp.StatusCode)
		return
	}

	// Decode response
	var updateRes UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateRes); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("Update sent successfully")
	fmt.Printf("Received response: Water=%.2f, Wind=%.2f, Status=%s\n", updateRes.Water, updateRes.Wind, updateRes.Status)
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
