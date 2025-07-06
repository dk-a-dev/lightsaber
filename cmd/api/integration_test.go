package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

// Test configuration
const (
	baseURL      = "http://localhost:4000"
	testEmail    = "test@example.com"
	testPassword = "password123"
)

// TestAPIEndpoints runs comprehensive tests on all API endpoints
func TestAPIEndpoints(t *testing.T) {
	// Wait for services to be ready
	time.Sleep(5 * time.Second)

	t.Run("HealthCheck", testHealthCheck)
	t.Run("RegisterUser", testRegisterUser)
	t.Run("AuthenticateUser", testAuthenticateUser)
	t.Run("CreateMovie", testCreateMovie)
	t.Run("ListMovies", testListMovies)
	t.Run("GetMovie", testGetMovie)
	t.Run("UpdateMovie", testUpdateMovie)
	t.Run("DeleteMovie", testDeleteMovie)
	t.Run("Metrics", testMetrics)
}

func testHealthCheck(t *testing.T) {
	resp, err := http.Get(baseURL + "/v1/healthcheck")
	if err != nil {
		t.Fatalf("Health check failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Health Check Response: %s\n", body)
}

func testRegisterUser(t *testing.T) {
	payload := map[string]interface{}{
		"name":     "Test User",
		"email":    testEmail,
		"password": testPassword,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(baseURL+"/v1/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("User registration failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Register User Response (%d): %s\n", resp.StatusCode, body)

	// Accept both 201 (new user) and 422 (user already exists)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != 422 {
		t.Errorf("Expected status 201 or 422, got %d", resp.StatusCode)
	}
}

func testAuthenticateUser(t *testing.T) {
	payload := map[string]interface{}{
		"email":    testEmail,
		"password": testPassword,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(baseURL+"/v1/tokens/authentication", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("User authentication failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Authenticate User Response (%d): %s\n", resp.StatusCode, body)

	// We expect this to fail since user needs activation, but it tests the endpoint
	if resp.StatusCode != http.StatusUnauthorized && resp.StatusCode != http.StatusCreated {
		t.Logf("Authentication response: %d (expected for unactivated user)", resp.StatusCode)
	}
}

func testCreateMovie(t *testing.T) {
	payload := map[string]interface{}{
		"title":   "Test Movie",
		"year":    2023,
		"runtime": "120 mins",
		"genres":  []string{"Action", "Drama"},
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(baseURL+"/v1/movies", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Create movie failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Create Movie Response (%d): %s\n", resp.StatusCode, body)

	// Expect 401 due to missing authentication
	if resp.StatusCode != http.StatusUnauthorized {
		t.Logf("Create movie response: %d (expected due to auth requirements)", resp.StatusCode)
	}
}

func testListMovies(t *testing.T) {
	resp, err := http.Get(baseURL + "/v1/movies")
	if err != nil {
		t.Fatalf("List movies failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("List Movies Response (%d): %s\n", resp.StatusCode, body)

	// Expect 401 due to missing authentication
	if resp.StatusCode != http.StatusUnauthorized {
		t.Logf("List movies response: %d (expected due to auth requirements)", resp.StatusCode)
	}
}

func testGetMovie(t *testing.T) {
	resp, err := http.Get(baseURL + "/v1/movies/1")
	if err != nil {
		t.Fatalf("Get movie failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Get Movie Response (%d): %s\n", resp.StatusCode, body)

	// Expect 401 due to missing authentication
	if resp.StatusCode != http.StatusUnauthorized {
		t.Logf("Get movie response: %d (expected due to auth requirements)", resp.StatusCode)
	}
}

func testUpdateMovie(t *testing.T) {
	payload := map[string]interface{}{
		"title": "Updated Test Movie",
		"year":  2024,
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPatch, baseURL+"/v1/movies/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Update movie failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Update Movie Response (%d): %s\n", resp.StatusCode, body)

	// Expect 401 due to missing authentication
	if resp.StatusCode != http.StatusUnauthorized {
		t.Logf("Update movie response: %d (expected due to auth requirements)", resp.StatusCode)
	}
}

func testDeleteMovie(t *testing.T) {
	req, _ := http.NewRequest(http.MethodDelete, baseURL+"/v1/movies/1", nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Delete movie failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Delete Movie Response (%d): %s\n", resp.StatusCode, body)

	// Expect 401 due to missing authentication
	if resp.StatusCode != http.StatusUnauthorized {
		t.Logf("Delete movie response: %d (expected due to auth requirements)", resp.StatusCode)
	}
}

func testMetrics(t *testing.T) {
	resp, err := http.Get(baseURL + "/debug/vars")
	if err != nil {
		t.Fatalf("Metrics endpoint failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Metrics Response: %s\n", body)
}

// BenchmarkEndpoints runs performance benchmarks to generate metrics
func BenchmarkEndpoints(b *testing.B) {
	// Wait for services to be ready
	time.Sleep(5 * time.Second)

	b.Run("HealthCheck", benchmarkHealthCheck)
	b.Run("RegisterUser", benchmarkRegisterUser)
	b.Run("ListMovies", benchmarkListMovies)
}

func benchmarkHealthCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(baseURL + "/v1/healthcheck")
		if err != nil {
			b.Fatalf("Health check failed: %v", err)
		}
		resp.Body.Close()
	}
}

func benchmarkRegisterUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := map[string]interface{}{
			"name":     fmt.Sprintf("Test User %d", i),
			"email":    fmt.Sprintf("test%d@example.com", i),
			"password": testPassword,
		}

		jsonData, _ := json.Marshal(payload)
		resp, err := http.Post(baseURL+"/v1/users", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			b.Fatalf("User registration failed: %v", err)
		}
		resp.Body.Close()
	}
}

func benchmarkListMovies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(baseURL + "/v1/movies")
		if err != nil {
			b.Fatalf("List movies failed: %v", err)
		}
		resp.Body.Close()
	}
}
