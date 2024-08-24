package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/madalinpopa/loto/internal/generator"
)

// APIInfo type holds details about API.
type APIInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// Route type holds details about different endpoints that can be used.
type Route struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

// RoodHandler handles requests to /api/v1/ route
func RootHandler(routes []Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(routes)
	}
}

// AboutHandler handles requests to /api/v1/about/
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	info := APIInfo{
		Name:        "Number Generator API",
		Version:     "0.1.0",
		Description: "This API generates 6 random numbers and stores the history in Google Sheets",
	}
	json.NewEncoder(w).Encode(info)
}

func GenerateNumbersHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue("length")

	// Attempt to convert string to number
	n, err := strconv.Atoi(ns)
	if err != nil {
		fmt.Fprint(w, "Not a valid number provide")
		log.Printf("Error: Not a valid number. Provided length: %s", ns)
		return
	}

	numbers := generator.GenerateNumbers(n)
	json.NewEncoder(w).Encode(numbers)
}
