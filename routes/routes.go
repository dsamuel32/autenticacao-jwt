package routes

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Protected   bool
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
