//tests/main_test.go
package main

import (
 "fmt"
 "net/http"
 "net/http/httptest"
 "testing"
)

// TestHandler tests the Handler function.
func TestHandler(t *testing.T) {
 // Create a new HTTP request.
 req, err := http.NewRequest("GET", "/", nil)
 if err != nil {
  t.Fatal(err)
 }

 // Create a new recorder to capture the response.
 rr := httptest.NewRecorder()

 // Create a new handler that calls the Handler function.
 handler := http.HandlerFunc(Handler)

 // Serve the request and capture the response code.
 handler.ServeHTTP(rr, req)

 // Check that the response code is 200.
 if status := rr.Code; status != http.StatusOK {
  t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
 }

 // Check that the response body is "Hello, World!".
 expected := "Hello, World!"
 if rr.Body.String() != expected {
  t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
 }
}

// Handler is a simple HTTP handler that returns "Hello, World!".
func Handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Hello, World!")
}
