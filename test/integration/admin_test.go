package integration

import (
	"net/http"
	"testing"
)

func TestGetAdmin(t *testing.T) {
	// Use localhost since we're testing locally
	url := "http://localhost:8081/796e62b1-b093-4a38-912d-26c1c82d763c/0/admin/b1f48088-4dc1-416e-bde0-e654f04064ee"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6Ijc5NmU2MmIxLWIwOTMtNGEzOC05MTJkLTI2YzFjODJkNzYzYyIsInJvbGUiOjAsImV4cCI6MTc0NjAxMDg4MH0.Ep8z04QPJQSlKJu6Uvmd_G4CRXTkYTB4x2GdQOl1e8M")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	// Add more assertions here based on the expected response body
}
