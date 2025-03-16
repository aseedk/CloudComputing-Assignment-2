package xutil

import (
	"cloud-computing/users/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// APIResponse represents the expected response structure
type APIResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CheckOrganizationExists sends a GET request and returns true if the organization exists, otherwise false
func CheckOrganizationExists(ctx context.Context, id string) bool {
	// Create an HTTP client
	log.Println("Organization URI: ", config.OrganizationURI+id)

	// Create the request
	resp, err := http.Get(config.OrganizationURI + id)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false
	}
	log.Println(string(body))
	// Parse JSON response
	var apiResp APIResponse
	if err = json.Unmarshal(body, &apiResp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return false
	}

	// Return true if success is true, otherwise false
	return apiResp.Success
}
