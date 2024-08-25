package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/madalinpopa/loto/internal/generator"
)

// APIInfo type holds details about API.
type Info struct {
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
func RootHandler(routes []Route) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, routes)
	}
}

// AboutHandler handles requests to /api/v1/about/
func AboutHandler(c echo.Context) error {
	info := Info{
		Name:        "Number Generator API",
		Version:     "0.1.0",
		Description: "This API generates 6 random numbers and stores the history in Google Sheets",
	}
	return c.JSON(http.StatusOK, info)
}

func GenerateNumbersHandler(c echo.Context) error {
	ns := c.PathParam("len")

	// Attempt to convert string to number
	n, err := strconv.Atoi(ns)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a valid number provided")
	}

	numbers := generator.GenerateNumbers(n)
	return c.JSON(http.StatusOK, numbers)
}
